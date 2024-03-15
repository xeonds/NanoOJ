package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"xyz.xeonds/nano-oj/config"
	"xyz.xeonds/nano-oj/lib"
	"xyz.xeonds/nano-oj/model"
	"xyz.xeonds/nano-oj/worker"
)

func main() {
	config := lib.LoadConfig[config.Config]()
	db := lib.NewDB(&config.DatabaseConfig, func(db *gorm.DB) error {
		return db.AutoMigrate(&model.Submission{}, &model.Problem{}, &model.User{}, &model.Contest{}, &model.Notification{}, &model.Rank{})
	})
	if config.ServerType == "main" || config.ServerType == "core" || config.ServerType == "web-judge" || config.ServerType == "judge" {
		go worker.JudgeEnqueuer(db)
		go worker.JudgeWorker(db, config)
	}
	if config.ServerType == "main" || config.ServerType == "core" || config.ServerType == "web" {
		redis := lib.NewRedis(&config.RedisConfig)
		router := gin.Default()
		apiRouter := router.Group("/api/v1")
		lib.AddCRUD[model.Problem](apiRouter, "/problems", db)
		lib.AddCRUD[model.Submission](apiRouter, "/submissions", db)
		lib.AddCRUD[model.Contest](apiRouter, "/contests", db)
		lib.AddCRUD[model.Notification](apiRouter, "/notifications", db)
		lib.AddCRUD[model.Rank](apiRouter, "/ranks", db)
		lib.AddCRUD[model.User](apiRouter, "/users", db)
		lib.AddCaptchaAPI(apiRouter, "/captcha", config.MailConfig, config.CaptchaConfig, redis)
		lib.AddLoginAPI(apiRouter, "/user", db)
		router.NoRoute(gin.WrapH(http.FileServer(gin.Dir("./dist", false))))

		if err := router.Run(config.ServerConfig.Port); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}
}
