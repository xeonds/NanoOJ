package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"xyz.xeonds/nano-oj/config"
	"xyz.xeonds/nano-oj/router"
	"xyz.xeonds/nano-oj/worker"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	conf, err := config.LoadConfig()
	if err != nil {
		log.Println("Failed to load config file, using default config")
		conf = config.DefaultConfig()
	}
	if conf.Server.Side == "web-judge" {
		worker.InitJudgerPool()
	}

	go judgeEnqueuer()
	go judgeWorker()

	if err := r.Run(fmt.Sprintf(":%d", conf.Server.Port)); err != nil {
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
