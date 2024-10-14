package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//Id    int `gorm:"primarykey; autoincrement"`
	Name  string
	Email string
}
