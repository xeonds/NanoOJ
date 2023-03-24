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
		database.NanoDB.Save(&submission)
		return
	}
	problem, err := database.GetProblemByID(uint32(id))
	if err != nil {
		submission.Status = "failed"
		database.NanoDB.Save(&submission)
		return
	}

	tempFolder := filepath.Join("tmp", fmt.Sprintf("temp_%d", time.Now().UnixNano()))
	_ = os.Mkdir(tempFolder, 0755)
	programFile := filepath.Join(tempFolder, "program.cc")
	err = os.WriteFile(programFile, []byte(sourceCode), 0644)
	if err != nil {
		submission.Status = "failed"
		database.NanoDB.Save(&submission)
		return
	}
	inputFiles := make([]string, len(problem.ProblemInputs))
	outputFiles := make([]string, len(problem.ExpectedOutputs))
	for i, inputFile := range problem.ProblemInputs {
		inputFiles[i] = filepath.Join(tempFolder, inputFile)
		err = os.WriteFile(inputFiles[i], []byte(inputFile), 0644)
		if err != nil {
			submission.Status = "failed"
			database.NanoDB.Save(&submission)
			return
		}
	}
	for i, outputFile := range problem.ExpectedOutputs {
		outputFiles[i] = filepath.Join(tempFolder, outputFile)
		err = os.WriteFile(outputFiles[i], []byte(outputFile), 0644)
		if err != nil {
			submission.Status = "failed"
			database.NanoDB.Save(&submission)
			return
		}
	}
	task := SourceCodeJudgement{
		Workdir:     tempFolder,
		Lang:        "c",
		SourceFile:  programFile,
		InputFiles:  inputFiles,
		ExpectFiles: outputFiles,
		TimeLimit:   1000,
	}
	result, err := task.Judge()
	if err != nil {
		submission.Status = "failed"
		database.NanoDB.Save(&submission)
		return
	}
	submission.Status = result
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
