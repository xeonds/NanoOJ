package worker

import (
	"fmt"

	"gorm.io/gorm"
	"xyz.xeonds/nano-oj/database"
	"xyz.xeonds/nano-oj/database/model"
)

var judgeQueue = make(chan model.Submission, 100)

func IsEmpty() bool {
	return len(judgeQueue) == 0
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
		judgeQueue <- submission
	}
}
