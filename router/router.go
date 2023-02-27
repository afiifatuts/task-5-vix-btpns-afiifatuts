package router

import (
	"github.com/afiifatuts/go-authentication/controllers"
	"github.com/afiifatuts/go-authentication/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/signup", controllers.Signup)
		userRouter.POST("/login", controllers.Login)
		userRouter.PUT("/update/:userId", controllers.UserUpdate)
		userRouter.DELETE("/delete/:userId", controllers.UserDelete)

	}

	photoRouter := r.Group("/photo")
	{
		photoRouter.Use(middleware.Authentication())

		photoRouter.POST("/add", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.ListPhoto)
		photoRouter.PUT("/edit/:photoID", middleware.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/delete/:photoId", middleware.PhotoAuthorization(), controllers.DeletePhoto)

	}

	return r

}
