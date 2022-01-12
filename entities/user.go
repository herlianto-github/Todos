package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//USERID AUTO GENERATE
	ID       uint
	Name     string
	Password string
	Todo     []ToDo
	Project  []Project
}
