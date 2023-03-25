package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"xyz.xeonds/nano-oj/controller"
)

func InitRouter(r *gin.Engine) {
	r.Use(corsMiddleware())

	// APIs
	apiRouter := r.Group("/api/v1")
	apiRouter.GET("/problems", controller.GetProblems)
	apiRouter.GET("/problems/:id", controller.GetProblemByID)
	apiRouter.POST("/problems", authMiddleware(1), controller.CreateProblem)
	apiRouter.PUT("/problems/:id", authMiddleware(1), controller.UpdateProblem)
	apiRouter.DELETE("/problems/:id", authMiddleware(1), controller.DeleteProblem)

	apiRouter.GET("/submissions", controller.GetSubmissions)
	apiRouter.GET("/submissions/:id", controller.GetSubmissionByID)
	apiRouter.POST("/submissions", authMiddleware(0), controller.CreateSubmission)

	apiRouter.GET("/contests", controller.GetContests)
	apiRouter.GET("/contests/:id", controller.GetContestByID)
	apiRouter.POST("/contests", authMiddleware(1), controller.CreateContest)
	apiRouter.PUT("/contests/:id", authMiddleware(1), controller.UpdateContest)
	apiRouter.DELETE("/contests/:id", authMiddleware(1), controller.DeleteContest)

	usersRouter := apiRouter.Group("/users")
	usersRouter.GET("", controller.GetUser)
	usersRouter.GET("/:user_id", controller.GetUser)
	usersRouter.PUT("/:user_id", authMiddleware(0), controller.UpdateUser)
	usersRouter.DELETE("/:user_id", authMiddleware(1), controller.DeleteUser)

	apiRouter.POST("/user/login", controller.Login)
	apiRouter.POST("/user/register", controller.Register)

	// Static files
	r.NoRoute(gin.WrapH(http.FileServer(gin.Dir("./public", false))))
}
