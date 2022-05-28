package models

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Title       string       `gorm:"not null"`
	UserID      uint         `gorm:"not null"`
	CategoryID  uint         `gorm:"not null"`
	CompletedAt sql.NullTime `gorm:"type:timestamp null"`
	Category    Category
}

func (task *Task) Save() (*Task, error) {
	err := DB.Create(&task).Error

	if err != nil {
		return &Task{}, err
	}

	return task, nil
}

func (task *Task) Update() (*Task, error) {
	err := DB.Save(&task).Error

	if err != nil {
		return &Task{}, err
	}

	return task, nil
}

func (task *Task) Delete() error {
	err := DB.Delete(&task).Error

	if err != nil {
		return err
	}

	return nil
}

func ListTask(userId uint) ([]Task, error) {
	var tasks []Task

	if err := DB.Where("user_id=?", userId).Find(&tasks).Error; err != nil {
		return []Task{}, err
	}

	return tasks, nil
}

func FindTaskById(taskId int) (Task, error) {
	var task Task

	if err := DB.First(&task, taskId).Error; err != nil {
		return Task{}, err
	}

	return task, nil
}
