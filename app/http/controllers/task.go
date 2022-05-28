package controllers

import (
	"database/sql"
	"go-todo/main/app/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskInput struct {
	Title      string `json:"title" binding:"required"`
	CategoryId uint   `json:"category_id" binding:"required"`
}

func ListTask(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	tasks, err := models.ListTask(user.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func GetTask(c *gin.Context) {
	categoryId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := models.FindTaskById(categoryId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func CreateTask(c *gin.Context) {
	var input TaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := c.MustGet("user").(models.User)

	task := models.Task{}
	task.Title = input.Title
	task.CategoryID = input.CategoryId
	task.UserID = user.ID

	_, err := task.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task has been created successfully!"})
}

func MarkTaskCompleted(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := models.FindTaskById(taskId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	task.CompletedAt = sql.NullTime{Time: time.Now(), Valid: true}
	_, err = task.Update()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task has been updated successfully!"})
}

func UpdateTask(c *gin.Context) {
	var input TaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := models.FindTaskById(taskId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	task.Title = input.Title
	task.CategoryID = input.CategoryId
	_, err = task.Update()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task has been updated successfully!"})
}

func DeleteTask(c *gin.Context) {
	categoryId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := models.FindTaskById(categoryId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = task.Delete()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task has been deleted successfully!"})
}
