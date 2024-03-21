package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"xyz.xeonds/nano-oj/config"
	"xyz.xeonds/nano-oj/database"
	"xyz.xeonds/nano-oj/lib"
	"xyz.xeonds/nano-oj/model"
	"xyz.xeonds/nano-oj/worker"
)

func main() {
	config := lib.LoadConfig[config.Config]()
	db := lib.NewDB(&config.DatabaseConfig, func(db *gorm.DB) error {
		return db.AutoMigrate(&model.Submission{}, &model.Problem{}, &model.User{}, &model.Contest{}, &model.Notification{}, &model.Rank{}, &model.PersonalInfo{}, &model.AccountInfo{})
	})
	if config.ServerType == "main" || config.ServerType == "core" || config.ServerType == "web-judge" || config.ServerType == "judge" {
		go worker.JudgeEnqueuer(db)
		go worker.JudgeWorker(db, config)
	}
	if config.ServerType == "main" || config.ServerType == "core" || config.ServerType == "web" {
		redis := lib.NewRedis(&config.RedisConfig)
		router := gin.Default()
		apiRouter := router.Group("/api/v1")
		lib.AddCRUDWithAuth[model.Problem](apiRouter, "/problems", db, 0, 1)
		lib.AddCRUDWithAuth[model.Submission](apiRouter, "/submissions", db, 0, 1)
		lib.APIBuilder(router, func(group *gin.RouterGroup) *gin.RouterGroup {
			group.GET("", lib.GetAll[model.Contest](db, database.GetAllContests))
			group.GET("/:id", lib.Get[model.Contest](db, database.GetContestById))
			group.GET("/:id/rank", lib.GetAll[model.Rank](db, database.GetRankByContestID))
			return group
		}, func(group *gin.RouterGroup) *gin.RouterGroup {
			group.Use(lib.JWTMiddleware(lib.AuthPermission(0, 1)))
			group.POST("", lib.Create[model.Contest](db, nil))
			group.PUT("/:id", lib.Update[model.Contest](db))
			group.DELETE("/:id", lib.Delete[model.Contest](db))
			return group
		})(apiRouter, "/contests")
		lib.APIBuilder(router, func(group *gin.RouterGroup) *gin.RouterGroup {
			return group
		}, func(group *gin.RouterGroup) *gin.RouterGroup {
			group.Use(lib.JWTMiddleware(lib.AuthPermission(0, 1)))
			group.GET("/problems/export", handleExport(db))
			group.POST("/problems/import", handleImport(db))
			return group
		})(apiRouter, "/admin")
		lib.AddCRUDWithAuth[model.Notification](apiRouter, "/notifications", db, 0, 1)
		lib.AddCRUDWithAuth[model.Rank](apiRouter, "/ranks", db, 0, 1)
		lib.APIBuilder(router, func(group *gin.RouterGroup) *gin.RouterGroup {
			group.GET("", lib.GetAll[model.Rank](db, nil))
			group.GET("/:id", lib.Get[model.Rank](db, nil))
			return group
		}, func(group *gin.RouterGroup) *gin.RouterGroup {
			group.Use(lib.JWTMiddleware(lib.AuthPermission(0, 2)))
			return group
		})(router, "/ranks")
		lib.AddCRUD[model.User](apiRouter, "/users", db).Use(lib.JWTMiddleware(lib.AuthPermission(0, 1)))
		lib.AddCaptchaAPI(apiRouter, "/captcha", config.MailConfig, config.CaptchaConfig, redis)
		lib.AddLoginAPI(apiRouter, "/user", db)
		actionRouter := apiRouter.Group("/actions").Use(lib.JWTMiddleware(lib.AuthPermission(0, 2)))
		actionRouter.POST("/submit", lib.Create(db, database.CreateSubmission))
		actionRouter.POST("/info/update", lib.Update[model.User](db)) // TODO: Add UserID check
		actionRouter.GET("/rerun/:id", lib.HandleReRunJudge(db))
		router.NoRoute(gin.WrapH(http.FileServer(gin.Dir("./dist", false))))

		if err := router.Run(config.ServerConfig.Port); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}
}

func handleExport(db *gorm.DB) func(*gin.Context) {
	problems := new([]model.Problem)
	if err := db.Find(problems).Error; err != nil {
		log.Println("[gorm] export problems failed")
	}
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, problems)
	}
}

func handleImport(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		problems := new([]model.Problem)
		if err := c.ShouldBindJSON(problems); err != nil {
			c.AbortWithError(http.StatusBadRequest, errors.New("[gin] parse problems failed"))
			return
		}
		if err := db.Create(problems).Error; err != nil {
			c.AbortWithError(http.StatusBadRequest, errors.New("[gorm] insert problems to database failed"))
			return
		}
		c.JSON(http.StatusOK, problems)
	}
}
