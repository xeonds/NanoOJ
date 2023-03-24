package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"xyz.xeonds/nano-oj/router"
	"xyz.xeonds/nano-oj/worker"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)

	go judgeEnqueuer()
	go judgeWorker()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func judgeEnqueuer() {
	log.Println("Judge enqueuer process starting...")
	for {
		worker.JudgeEnqueue()
		time.Sleep(5 * time.Second)
	}
}

func judgeWorker() {
	log.Println("Judge worker processes starting...")
	for {
		if !worker.IsEmpty() {
			go worker.JudgeWorker()
		}
		time.Sleep(1 * time.Second)
	}
}
