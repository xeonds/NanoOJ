package main

import (
	"github.com/gin-gonic/gin"
	"xyz.xeonds/nano-oj/controller"
)

func initRouter(r *gin.Engine) {
	// APIs
	apiRouter := r.Group("/api/v1")
	apiRouter.POST("/judge", controller.Judge)
	apiRouter.POST("/login", controller.Login)
	apiRouter.POST("/register", controller.Register)

	// Static files
	r.Static("/", "./public")
}
