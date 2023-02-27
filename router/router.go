package router

import (
	"github.com/afiifatuts/go-authentication/controllers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/signup", controllers.Signup)
		userRouter.POST("/login", controllers.Login)
		userRouter.PUT("/update/:userId", controllers.UserUpdate)
		// userRouter.DELETE("/:userId", controllers.UserDelete)

	}

	// photoRouter := r.Group("/photo")
	// {
	// 	photoRouter.Use(middleware.RequireAuth)

	// 	photoRouter.POST("/", controllers.CreatePhoto)
	// 	photoRouter.GET("/", controllers.ListPhoto)
	// 	photoRouter.PUT("/:photoID", middleware.PhotoAuthorization(), controllers.UpdatePhoto)
	// 	photoRouter.DELETE("/", middleware.PhotoAuthorization(), controllers.DeletePhoto)
	// }
	return r

}
