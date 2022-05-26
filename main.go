package main

import (
	"go-todo/main/app/http/controllers"
	"go-todo/main/app/http/middlewares"
	"go-todo/main/app/models"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	protected.GET("/user", controllers.User)

	return r
}

func main() {
	models.ConnectDatabase()

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
