package model

import (
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username  string `gorm:"unique_index;not null"`
	Email     string `gorm:"unique_index;not null"`
	Password  string `gorm:"not null"`
	Favorites []Book `gorm:"many2many:favorites;"`
}

func (u *User) HashPassword(inputPassword string) (string, error) {
	if len(inputPassword) == 0 {
		return "", errors.New("password should not be empty")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(inputPassword), bcrypt.DefaultCost)
	return string(hashPassword), err
}

func (u *User) CheckPassword(inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(inputPassword))
	return err == nil
}
