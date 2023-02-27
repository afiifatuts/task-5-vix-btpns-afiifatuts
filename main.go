package main

import (
	"github.com/afiifatuts/go-authentication/initializer"
	"github.com/afiifatuts/go-authentication/router"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDb()
	initializer.SyncDatabase()
}

func main() {
	// r := gin.Default()

	// r.POST("/signup", controllers.Signup)
	// r.POST("/login", controllers.Login)
	// r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r := router.Router()
	r.Run()
}
