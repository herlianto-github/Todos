package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//USERID AUTO GENERATE
	Name     string
	Password string
	Todo     []ToDo
}
