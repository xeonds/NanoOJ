package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"xyz.xeonds/nano-oj/controller"
)

func InitRouter(r *gin.Engine) {
	// Main router
	r.Use(corsMiddleware())

	// APIs
	apiRouter := r.Group("/api/v1")
	apiRouter.Use(authMiddleware())

	apiRouter.GET("/problems", controller.GetProblems)
	apiRouter.GET("/problems/:id", controller.GetProblemByID)
	apiRouter.POST("/problems", controller.CreateProblem)
	apiRouter.PUT("/problems/:id", controller.UpdateProblem)
	apiRouter.DELETE("/problems/:id", controller.DeleteProblem)

	apiRouter.GET("/submissions", controller.GetSubmissions)
	apiRouter.GET("/submissions/:id", controller.GetSubmissionByID)
	apiRouter.POST("/submissions", controller.CreateSubmission)

	apiRouter.GET("/contests", controller.GetContests)
	apiRouter.GET("/contests/:id", controller.GetContestByID)
	apiRouter.POST("/contests", controller.CreateContest)
	apiRouter.PUT("/contests/:id", controller.UpdateContest)
	apiRouter.DELETE("/contests/:id", controller.DeleteContest)

	usersRouter := apiRouter.Group("/users")
	usersRouter.GET("", controller.GetUser)
	usersRouter.GET("/:user_id", controller.GetUser)
	usersRouter.PUT("/:user_id", controller.UpdateUser)
	usersRouter.DELETE("/:user_id", controller.DeleteUser)

	apiRouter.POST("/user/login", controller.Login)
	apiRouter.POST("/user/register", controller.Register)

	// Static files
	r.NoRoute(gin.WrapH(http.FileServer(gin.Dir("./public", false))))
}
