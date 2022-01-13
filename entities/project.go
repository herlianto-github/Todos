package entities

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	//PROJECTID AUTO GENERATE
	ID          uint
	ProjectName string
	UserId      uint
	Todo        []ToDo
}
