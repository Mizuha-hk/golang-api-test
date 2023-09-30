package models

import (
	"errors"

	"golang-api/db"
)

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}


func GetUserByNameAndPassword(name, password string) (*User, error) {
	var user User
	result := db.DB.Where("name = ? AND password = ?", name, password).First(&user)
	if result.Error != nil {
		// no user found or other error occurred
		return nil, errors.New("invalid credentials")
	}
	return &user, nil
}
