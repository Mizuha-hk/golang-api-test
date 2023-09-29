package models

type User struct {
	Id int `gorm:"primaryKey" json "id"`
	Name string `json "name"`
}