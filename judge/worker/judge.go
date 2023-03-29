package worker

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"xyz.xeonds/nano-oj/database"
	"xyz.xeonds/nano-oj/database/model"
)

var JudgeQueue = make(chan model.Submission, 100)

type task struct {
	Workdir     string
	Lang        string
	SourceFile  string
	InputFiles  []string
	ExpectFiles []string
	TimeLimit   int
}

func (s *task) judge() (string, error) {
	switch s.Lang {
	case "c", "cpp":
		return s.cJudger(), nil
	// case "python":
	// return judgement.pyJudger(), nil
	default:
		return "", fmt.Errorf("unsupported language: %s", s.Lang)
	}
}

func IsEmpty() bool {
	return len(JudgeQueue) == 0
}

func JudgeWorker() {
	// read a submission from judgeQueue
	submission := <-JudgeQueue
	submission.Status = "processing"
	database.NanoDB.Save(&submission)
	problemID := strconv.Itoa(int(submission.ProblemID))
	sourceCode := submission.Code
	id, err := strconv.ParseUint(problemID, 10, 32)
	if err != nil {
		submission.Status = "failed"
		fmt.Println("failed to parse", err)
		database.NanoDB.Save(&submission)
		return
	}
	// get the problem
	problem, err := database.GetProblemByID(uint32(id))
	if err != nil {
		submission.Status = "failed"
		fmt.Println("failed to fetch problem", err)
		database.NanoDB.Save(&submission)
		return
	}
	// create workdir
	tempFolder := filepath.Join("tmp", fmt.Sprintf("temp_%d", time.Now().UnixNano()))
	_ = os.MkdirAll(tempFolder, 0755)
	programFile := filepath.Join(tempFolder, "program.cc")
	err = os.WriteFile(programFile, []byte(sourceCode), 0644)
	if err != nil {
		submission.Status = "failed"
		fmt.Println("failed to create workdir", err)
		database.NanoDB.Save(&submission)
		return
	}
	// load input/output files to disk
	inputFiles := make([]string, len(problem.Inputs))
	outputFiles := make([]string, len(problem.Outputs))
	for i, inputFile := range problem.Inputs {
		inputFiles[i] = filepath.Join(tempFolder, inputFile)
		err = os.WriteFile(inputFiles[i], []byte(inputFile), 0644)
		if err != nil {
			submission.Status = "failed"
			fmt.Println("failed to load file", err)
			database.NanoDB.Save(&submission)
			return
		}
	}
	for i, outputFile := range problem.Outputs {
		outputFiles[i] = filepath.Join(tempFolder, outputFile)
		err = os.WriteFile(outputFiles[i], []byte(outputFile), 0644)
		if err != nil {
			fmt.Println("failed to load file", err)
			database.NanoDB.Save(&submission)
			return
		}
	}
	// build a new task and run
	t := task{
		Workdir:     tempFolder,
		Lang:        submission.Language,
		SourceFile:  programFile,
		InputFiles:  inputFiles,
		ExpectFiles: outputFiles,
		TimeLimit:   1000,
	}
	result, err := t.judge()
	if err != nil {
		submission.Status = "failed"
		fmt.Println("language not supported", err)
		database.NanoDB.Save(&submission)
		return
	}
	submission.Status = result
	fmt.Println(result)
	database.NanoDB.Save(&submission)
}

func JudgeEnqueue() {
	submissions, err := database.GetSubmissionsByStatus("waiting")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, submission := range submissions {
		JudgeQueue <- submission
	}
}
