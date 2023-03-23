package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"xyz.xeonds/nano-oj/worker"
)

func main() {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	})
	initRouter(router)

	go func() {
		for {
			worker.JudgeEnqueue()
			time.Sleep(5 * time.Second)
		}
	}()
	go func() {
		for {
			if !worker.IsEmpty() {
				go worker.JudgeWorker()
			}
			time.Sleep(1 * time.Second)
		}
	}()

	panic(router.Run(":8080"))
}
