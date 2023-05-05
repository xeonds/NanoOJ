package worker

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"xyz.xeonds/nano-oj/database"
	"xyz.xeonds/nano-oj/database/model"
)

type task struct {
	Submission  model.Submission
	Workdir     string
	Lang        string
	SourceFile  string
	InputFiles  []string
	ExpectFiles []string
	TimeLimit   int
}

type result struct {
	Status model.Status
	Info   string
}

func JudgeWorker() {
	submission := <-judgeQueue // read a submission from judgeQueue
	CommitStatus(submission, model.InProgress, "Judging...")

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

	GetAvailableJudger().AddTask(task{ // build a new task and run
		Submission:  submission,
		Workdir:     tempFolder,
		Lang:        submission.Language,
		SourceFile:  programFile,
		InputFiles:  inputFiles,
		ExpectFiles: outputFiles,
		TimeLimit:   problem.TimeLimit,
	})
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
		CommitStatus(submission, model.CompilationError, "Internal error, please contact admin")
		log.Println(info, err)
		return true
	}
	return false
}

func CommitStatus(submission model.Submission, stat model.Status, info ...string) {
	submission.Status = stat
	submission.Information = append(submission.Information, info...)
	database.NanoDB.Save(&submission)
}
