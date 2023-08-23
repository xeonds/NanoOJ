package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"xyz.xeonds/nano-oj/controller"
	"xyz.xeonds/nano-oj/middleware"
)

func InitRouter(r *gin.Engine) {
	r.Use(middleware.CorsMiddleware())

	// APIs
	apiRouter := r.Group("/api/v1")
	apiRouter.GET("/problems", controller.GetProblems)
	apiRouter.GET("/problems/:id", controller.GetProblemByID)
	apiRouter.POST("/problems", middleware.AuthMiddleware(1), controller.CreateProblem)
	apiRouter.PUT("/problems/:id", middleware.AuthMiddleware(1), controller.UpdateProblem)
	apiRouter.DELETE("/problems/:id", middleware.AuthMiddleware(1), controller.DeleteProblem)

	apiRouter.GET("/submissions", controller.GetSubmissions)
	apiRouter.GET("/submissions/:id", controller.GetSubmissionByID)
	apiRouter.POST("/submissions", middleware.AuthMiddleware(0), controller.CreateSubmission)

	apiRouter.GET("/contests", controller.GetContests)
	apiRouter.GET("/contests/:id", controller.GetContestByID)
	apiRouter.POST("/contests", middleware.AuthMiddleware(1), controller.CreateContest)
	apiRouter.PUT("/contests/:id", middleware.AuthMiddleware(1), controller.UpdateContest)
	apiRouter.DELETE("/contests/:id", middleware.AuthMiddleware(1), controller.DeleteContest)

	apiRouter.GET("/notifications", controller.GetNotifications)
	apiRouter.GET("/notifications/:id", controller.GetNotificationByID)
	apiRouter.POST("/notifications", middleware.AuthMiddleware(0), controller.CreateNotification)
	apiRouter.PUT("/notifications/:id", middleware.AuthMiddleware(0), controller.UpdateNotification)
	apiRouter.DELETE("/notifications/:id", middleware.AuthMiddleware(0), controller.DeleteNotification)

	usersRouter := apiRouter.Group("/users")
	usersRouter.GET("", controller.GetUser)
	usersRouter.GET("/:user_id", controller.GetUser)
	usersRouter.PUT("/:user_id", middleware.AuthMiddleware(0), controller.UpdateUser)
	usersRouter.DELETE("/:user_id", middleware.AuthMiddleware(1), controller.DeleteUser)

	apiRouter.POST("/user/login", controller.Login)
	apiRouter.POST("/user/register", controller.Register)

	// Static files
	r.NoRoute(gin.WrapH(http.FileServer(gin.Dir("./public", false))))
}
