package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name   string `gorm:"size:255;not null;" json:"name"`
	UserID uint
}

func (category *Category) Save() (*Category, error) {
	err := DB.Create(&category).Error

	if err != nil {
		return &Category{}, err
	}

	return category, nil
}

func (category *Category) Update() (*Category, error) {
	err := DB.Save(&category).Error

	if err != nil {
		return &Category{}, err
	}

	return category, nil
}

func (category *Category) Delete() error {
	err := DB.Delete(&category).Error

	if err != nil {
		return err
	}

	return nil
}

func ListCategory(userId uint) ([]Category, error) {
	var categories []Category

	if err := DB.Where("user_id=?", userId).Find(&categories).Error; err != nil {
		return []Category{}, err
	}

	return categories, nil
}

func FindCategoryById(categoryId int) (Category, error) {
	var category Category

	if err := DB.First(&category, categoryId).Error; err != nil {
		return Category{}, err
	}

	return category, nil
}
