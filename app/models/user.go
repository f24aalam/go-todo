package models

import (
	"errors"
	"go-todo/main/app/utils/token"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name       string `gorm:"size:255;not null;" json:"name"`
	Email      string `gorm:"size:255;not null;unique;" json:"email"`
	Password   string `gorm:"size:255;not null;" json:"password"`
	Categories []Category
}

func (user *User) Save() (*User, error) {

	err := DB.Create(&user).Error
	if err != nil {
		return &User{}, err
	}

	return user, nil
}

func (user *User) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return nil
}

func GetUserById(userId uint) (User, error) {
	var user User

	if err := DB.First(&user, userId).Error; err != nil {
		return user, errors.New("User Not Found")
	}

	user.PrepareGive()

	return user, nil
}

func (user *User) PrepareGive() {
	user.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email string, password string) (string, error) {
	var err error

	user := User{}

	err = DB.Model(User{}).Where("email = ?", email).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.Generate(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
