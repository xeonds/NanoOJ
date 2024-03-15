package worker

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/gorm"
	"xyz.xeonds/nano-oj/config"
	"xyz.xeonds/nano-oj/database"
	"xyz.xeonds/nano-oj/model"
)

type Task struct {
	Submission  model.Submission
	Workdir     string
	Lang        string
	SourceFile  string
	InputFiles  []string
	ExpectFiles []string
	TimeLimit   int
}

type Result struct {
	Status model.Status
	Info   string
}

var JudgeQueue = make(chan model.Submission, 100)
var LocalJudge = map[string]func(*Task) (model.Status, string){
	"cpp":    CPP,
	"c":      CPP,
	"python": Python,
}

func JudgeEnqueuer(db *gorm.DB) {
	log.Println("Judge enqueuer process starting...")
	for {
		JudgeEnqueue(db)
		time.Sleep(5 * time.Second)
	}
}

// Scans the database, and enqueue all pending submissions timely
func JudgeEnqueue(db *gorm.DB) {
	repo := database.Repository{DB: db}
	submissions, err := repo.GetSubmissionsByStatus(model.Pending)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, submission := range submissions {
		JudgeQueue <- submission
	}
}

func JudgeWorker(db *gorm.DB, config *config.Config) {
	log.Println("Judge worker processes starting...")
	if config.ServerType == "web-judge" || config.ServerType == "main" {
		InitJudgerPool(config)
		RunJudgerPool()
		for {
			if !IsEmpty() {
				go func() {
					if t := FetchOneTaskFromList(db); t != nil {
						GetAvailableJudger().AddTask(t)
					}
				}()
			}
			time.Sleep(3 * time.Second)
		}
	} else if config.ServerType == "judge" || config.ServerType == "core" {
		for {
			if !IsEmpty() {
				go func() {
					if t := FetchOneTaskFromList(db); t != nil {
						LocalJudge[t.Lang](t)
					}
				}()
			}
			time.Sleep(3 * time.Second)
		}
	}
}

func IsEmpty() bool {
	return len(JudgeQueue) == 0
}

func FetchOneTaskFromList(db *gorm.DB) *Task { // Create task & enqueue it
	submission := <-JudgeQueue // read a submission from judgeQueue
	CommitStatus(db, submission, model.InProgress, "Judging...")
	repo := database.Repository{DB: db}
	sourceCode := submission.Code //fetch all required files for judge
	problem, err := repo.GetProblemByID(submission.ProblemID)
	if errorHandler(err, submission, "Failed to fetch problem", db) {
		return nil
	}
	tempFolder, programFile, inputFiles, outputFiles, shouldReturn := initWorkDir(sourceCode, submission, problem, db)
	if shouldReturn {
		log.Println("Failed to initialize workdir")
		return nil
	}

	return &Task{ // build a new task and run
		Submission:  submission,
		Workdir:     tempFolder,
		Lang:        submission.Language,
		SourceFile:  programFile,
		InputFiles:  inputFiles,
		ExpectFiles: outputFiles,
		TimeLimit:   problem.TimeLimit,
	}
}

func initWorkDir(sourceCode string, submission model.Submission, problem *model.Problem, db *gorm.DB) (string, string, []string, []string, bool) {
	tempFolder := filepath.Join("tmp", fmt.Sprintf("temp_%d", time.Now().UnixNano()))
	_ = os.MkdirAll(tempFolder, 0755)
	programFile := filepath.Join(tempFolder, "program.cc")
	err := os.WriteFile(programFile, []byte(sourceCode), 0644)
	if errorHandler(err, submission, "Failed to load source code", db) {
		return "", "", nil, nil, true
	}
	inputFiles := make([]string, len(problem.Inputs))
	outputFiles := make([]string, len(problem.Outputs))
	for i, inputFile := range problem.Inputs {
		inputFiles[i] = filepath.Join(tempFolder, fmt.Sprintf("%d.in", i))
		err = os.WriteFile(inputFiles[i], []byte(inputFile), 0644)
		if errorHandler(err, submission, "Failed to load files", db) {
			return "", "", nil, nil, true
		}
	}
	for i, outputFile := range problem.Outputs {
		outputFiles[i] = filepath.Join(tempFolder, fmt.Sprintf("%d.out", i))
		err = os.WriteFile(outputFiles[i], []byte(outputFile), 0644)
		if errorHandler(err, submission, "Failed to load files", db) {
			return "", "", nil, nil, true
		}
	}
	return tempFolder, programFile, inputFiles, outputFiles, false
}

func errorHandler(err error, submission model.Submission, info string, db *gorm.DB) bool {
	if err != nil {
		CommitStatus(db, submission, model.CompilationError, "Internal error, please contact admin")
		log.Println(info, err)
		return true
	}
	return false
}

func CommitStatus(db *gorm.DB, submission model.Submission, stat model.Status, info ...string) {
	submission.Status = stat
	submission.Information = append(submission.Information, info...)
	db.Save(&submission)
}

func ParseResult(out io.Reader) (model.Status, string, error) {
	var result model.Status
	var info string
	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Result: ") {
			result = model.Status(line[8:])
		} else if strings.HasPrefix(line, "Info: ") {
			info += line[6:] + "\n"
		}
	}
	if err := scanner.Err(); err != nil {
		return model.CompilationError, "Failed to read container logs", err
	}
	return result, info, nil
}
