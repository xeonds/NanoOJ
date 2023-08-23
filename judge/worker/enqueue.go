package worker

import (
	"fmt"

	"xyz.xeonds/nano-oj/database"
	"xyz.xeonds/nano-oj/database/model"
)

var judgeQueue = make(chan model.Submission, 100)

func IsEmpty() bool {
	return len(judgeQueue) == 0
}

// Scans the database, and enqueue all pending submissions timely
func JudgeEnqueue() {
	submissions, err := database.GetSubmissionsByStatus(model.Pending)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, submission := range submissions {
		judgeQueue <- submission
	}
}
