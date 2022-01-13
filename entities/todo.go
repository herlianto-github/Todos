package entities

import "gorm.io/gorm"

type ToDo struct {
	gorm.Model
	//TODO_ID AUTO GENERATE
	ID          int
	ProjectID   *uint
	UserID      uint
	Task        string `gorm:"default: In Progress"`
	Status      string
	Description string
}
