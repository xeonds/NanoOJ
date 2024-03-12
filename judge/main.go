package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"xyz.xeonds/nano-oj/config"
	"xyz.xeonds/nano-oj/controller"
	"xyz.xeonds/nano-oj/database/model"
	"xyz.xeonds/nano-oj/lib"
	"xyz.xeonds/nano-oj/worker"
)

func main() {
	config, err := lib.LoadConfig[config.Config]()
	if err != nil {
		log.Fatal("Failed to load config file")
	}
	migrator := func(db *gorm.DB) error {
		return db.AutoMigrate(&model.Submission{}, &model.Problem{}, &model.User{}, &model.Contest{}, &model.Notification{})
	}
	db := lib.NewDB(&config.DatabaseConfig, migrator)
	if config.ServerType == "web-judge" {
		worker.InitJudgerPool()
		go judgeEnqueuer(db)
		go judgeWorker(db)
	}
	redis := lib.NewRedis(&config.RedisConfig)
	router := gin.Default()
	apiRouter := router.Group("/api/v1")
	apiRouter.Use(lib.AuthMiddleware(0, 0))
	lib.AddCRUD[model.Problem](apiRouter, "/problems", db)
	lib.AddCRUD[model.Submission](apiRouter, "/submissions", db)
	lib.AddCRUD[model.Contest](apiRouter, "/contests", db)
	lib.AddCRUD[model.Notification](apiRouter, "/notifications", db)
	lib.AddCRUD[model.User](apiRouter, "/users", db)
	lib.AddCaptchaAPI(apiRouter, "/captcha", config.MailConfig, config.CaptchaConfig, redis)
	apiRouter.POST("/user/login", controller.Login)
	apiRouter.POST("/user/register", controller.Register)
	router.NoRoute(gin.WrapH(http.FileServer(gin.Dir("./dist", false))))

	if err := router.Run(config.ServerConfig.Port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func judgeEnqueuer(db *gorm.DB) {
	log.Println("Judge enqueuer process starting...")
	for {
		worker.JudgeEnqueue(db)
		time.Sleep(5 * time.Second)
	}
}

func judgeWorker(db *gorm.DB) {
	log.Println("Judge worker processes starting...")
	for {
		if !worker.IsEmpty() {
			go worker.JudgeWorker(db)
		}
		time.Sleep(1 * time.Second)
	}
}
