package entities

import "gorm.io/gorm"

type To_Do struct {
	gorm.Model
	//TODO_ID AUTO GENERATE
	ProjectID   uint
	UserID      uint
	Task        string
	Status      string
	Description string
}
