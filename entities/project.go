package entities

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	//PROJECTID AUTO GENERATE
	ProjectName string
	UserId      uint
	Todo        []ToDo
}