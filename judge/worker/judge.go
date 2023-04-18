package worker

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"xyz.xeonds/nano-oj/database"
	"xyz.xeonds/nano-oj/database/model"
	"xyz.xeonds/nano-oj/utils"
)

type task struct {
	Workdir     string
	Lang        string
	SourceFile  string
	InputFiles  []string
	ExpectFiles []string
	TimeLimit   int
}

func (s task) run() (model.Status, string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	ctx := context.Background()
	if err != nil {
		return model.CompilationError, "Docker daemon connection failed", err
	}
	defer cli.Close()

	// create a container from the judge image, give it the workdir and source file as volume and run it
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:           "nano-oj/judge",
		Env:             []string{fmt.Sprintf("LANG=%s", s.Lang), fmt.Sprintf("TIME_LIMIT=%d", s.TimeLimit)},
		NetworkDisabled: true,
	}, &container.HostConfig{
		Binds: []string{fmt.Sprintf("%s:/code", s.Workdir)},
	}, nil, nil, "")
	if err != nil {
		return model.CompilationError, "Failed to create container", err
	}
	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return model.CompilationError, "Failed to start container", err
	}
	// wait for the container to finish
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return model.CompilationError, "Failed to wait container", err
		}
	case <-statusCh:
	}
	// read the result from the container
	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return model.CompilationError, "Failed to read container logs", err
	}
	// delete the container
	err = cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{})
	if err != nil {
		return model.CompilationError, "Failed to remove container", err
	}
	// parse the result
	return utils.ParseResult(out)
}

func JudgeWorker() {
	submission := <-judgeQueue // read a submission from judgeQueue
	commitStatus(submission, model.InProgress, "Judging...")

	sourceCode := submission.Code //fetch all required files for judge
	problem, err := database.GetProblemByID(submission.ProblemID)
	if errorHandler(err, submission, "Failed to fetch problem") {
		return
	}
	tempFolder, programFile, inputFiles, outputFiles, shouldReturn := initWorkDir(sourceCode, submission, problem)
	if shouldReturn {
		log.Println("Failed to initialize workdir")
		return
	}

	result, info, err := task{ // build a new task and run
		Workdir:     tempFolder,
		Lang:        submission.Language,
		SourceFile:  programFile,
		InputFiles:  inputFiles,
		ExpectFiles: outputFiles,
		TimeLimit:   problem.TimeLimit,
	}.run()
	commitStatus(submission, result, strings.Split(info, "\n")...)
	if err != nil {
		log.Println("Failed to run task", err)
	}

	_ = os.RemoveAll(tempFolder)
}

func initWorkDir(sourceCode string, submission model.Submission, problem *model.Problem) (string, string, []string, []string, bool) {
	tempFolder := filepath.Join("tmp", fmt.Sprintf("temp_%d", time.Now().UnixNano()))
	_ = os.MkdirAll(tempFolder, 0755)
	programFile := filepath.Join(tempFolder, "program.cc")
	err := os.WriteFile(programFile, []byte(sourceCode), 0644)
	if errorHandler(err, submission, "Failed to load source code") {
		return "", "", nil, nil, true
	}
	inputFiles := make([]string, len(problem.Inputs))
	outputFiles := make([]string, len(problem.Outputs))
	for i, inputFile := range problem.Inputs {
		inputFiles[i] = filepath.Join(tempFolder, fmt.Sprintf("%d.in", i))
		err = os.WriteFile(inputFiles[i], []byte(inputFile), 0644)
		if errorHandler(err, submission, "Failed to load files") {
			return "", "", nil, nil, true
		}
	}
	for i, outputFile := range problem.Outputs {
		outputFiles[i] = filepath.Join(tempFolder, fmt.Sprintf("%d.out", i))
		err = os.WriteFile(outputFiles[i], []byte(outputFile), 0644)
		if errorHandler(err, submission, "Failed to load files") {
			return "", "", nil, nil, true
		}
	}
	return tempFolder, programFile, inputFiles, outputFiles, false
}

func errorHandler(err error, submission model.Submission, info string) bool {
	if err != nil {
		commitStatus(submission, model.CompilationError, "Internal error, please contact admin")
		log.Println(info, err)
		return true
	}
	return false
}

func commitStatus(submission model.Submission, stat model.Status, info ...string) {
	submission.Status = stat
	submission.Information = append(submission.Information, info...)
	database.NanoDB.Save(&submission)
}
