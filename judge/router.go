package main

import (
	"github.com/gin-gonic/gin"
	"xyz.xeonds/nano-oj/controller"
)

func initRouter(r *gin.Engine) {
	// APIs
	apiRouter := r.Group("/api/v1")
	
	judgeRouter := apiRouter.Group("/judge")
	judgeRouter.POST("/problems/:id", controller.JudgeProblemByID)
		
	apiRouter.GET("/problems", controller.GetProblems)
	apiRouter.GET("/problems/:id", controller.GetProblemByID)
	apiRouter.POST("/problems", controller.CreateProblem)
	apiRouter.PUT("/problems/:id", controller.UpdateProblem)
	apiRouter.DELETE("/problems/:id", controller.DeleteProblem)
	
	apiRouter.GET("/submissions", controller.GetSubmissions)
	apiRouter.GET("/submissions/:id", controller.GetSubmissionByID)
	
	apiRouter.GET("/users", controller.GetUsers)
	apiRouter.POST("/user/register", controller.Register)
	apiRouter.POST("/user/login", controller.Login)
	apiRouter.GET("/user/:id", controller.GetUserByID)
	apiRouter.PUT("/user/:id", controller.UpdateUser)
	apiRouter.DELETE("/user/:id", controller.DeleteUser)
	
	// Static files
	r.Static("/", "./public")
}
