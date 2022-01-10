package entities

import "gorm.io/gorm"

type To_Do struct {
	gorm.Model
	//TODO_ID AUTO GENERATE
	ProjectID   uint   `json:"projectid" form:"projectid"`
	Task        string `json:"task" form:"task"`
	description string `json:"description" form:"description"`
}
