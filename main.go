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

	protected.GET("/categories", controllers.ListCategory)
	protected.GET("/categories/:id", controllers.GetCategory)
	protected.POST("/categories", controllers.CreateCategory)
	protected.PUT("/categories/:id", controllers.UpdateCategory)
	protected.DELETE("/categories/:id", controllers.DeleteCategory)

	protected.GET("/tasks", controllers.ListTask)
	protected.GET("/tasks/:id", controllers.GetTask)
	protected.POST("/tasks", controllers.CreateTask)
	protected.PUT("/tasks/:id", controllers.UpdateTask)
	protected.PUT("/tasks/:id/mark-completed", controllers.MarkTaskCompleted)
	protected.DELETE("/tasks/:id", controllers.DeleteTask)

	return r
}

func main() {
	models.ConnectDatabase()

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
