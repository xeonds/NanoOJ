package worker

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

func (s *task) judge() (model.Status, string, error) {
	switch s.Lang {
	case "c", "cpp":
		res, info := s.cJudger()
		return res, info, nil
	// case "python":
	// return s.pyJudger(), nil
	default:
		return model.CompilationError, "", fmt.Errorf("unsupported language: %s", s.Lang)
	}
}

func IsEmpty() bool {
	return len(JudgeQueue) == 0
}

func JudgeWorker() {
	// read a submission from judgeQueue
	submission := <-JudgeQueue
	submission.Status = model.Pending
	database.NanoDB.Save(&submission)
	problemID := strconv.Itoa(int(submission.ProblemID))
	sourceCode := submission.Code
	id, err := strconv.ParseUint(problemID, 10, 32)
	if err != nil {
		submission.Status = 0
		fmt.Println("failed to parse", err)
		// submission.Information = append(submission.Information, "Failed to parse problem id")
		database.NanoDB.Save(&submission)
		return
	}
	// get the problem
	problem, err := database.GetProblemByID(uint32(id))
	if err != nil {
		submission.Status = 0
		fmt.Println("failed to fetch problem", err)
		// submission.Information = append(submission.Information, "Failed to fetch problem")
		database.NanoDB.Save(&submission)
		return
	}
	// create workdir
	tempFolder := filepath.Join("tmp", fmt.Sprintf("temp_%d", time.Now().UnixNano()))
	_ = os.MkdirAll(tempFolder, 0755)
	programFile := filepath.Join(tempFolder, "program.cc")
	err = os.WriteFile(programFile, []byte(sourceCode), 0644)
	if err != nil {
		submission.Status = 0
		fmt.Println("failed to create workdir", err)
		database.NanoDB.Save(&submission)
		return
	}
	// load input/output files to disk
	inputFiles := make([]string, len(problem.Inputs))
	outputFiles := make([]string, len(problem.Outputs))
	for i, inputFile := range problem.Inputs {
		inputFiles[i] = filepath.Join(tempFolder, fmt.Sprintf("%d.in", i))
		err = os.WriteFile(inputFiles[i], []byte(inputFile), 0644)
		if err != nil {
			submission.Status = 0
			fmt.Println("failed to load file", err)
			database.NanoDB.Save(&submission)
			return
		}
	}
	for i, outputFile := range problem.Outputs {
		outputFiles[i] = filepath.Join(tempFolder, fmt.Sprintf("%d.out", i))
		err = os.WriteFile(outputFiles[i], []byte(outputFile), 0644)
		if err != nil {
			submission.Status = 0
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
		TimeLimit:   5,
	}
	result, info, err := t.judge()
	if err != nil {
		submission.Status = model.CompilationError
		fmt.Println("language not supported", err)
		submission.Information = append(submission.Information, "Language not supported")
		database.NanoDB.Save(&submission)
		return
	}
	submission.Status = result
	submission.Information = append(submission.Information, strings.Split(info, "\n")...)
	fmt.Println(result)
	database.NanoDB.Save(&submission)
}

func JudgeEnqueue() {
	submissions, err := database.GetSubmissionsByStatus(model.Pending)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, submission := range submissions {
		JudgeQueue <- submission
	}
}
