package entities

import "gorm.io/gorm"

type To_Do struct {
	gorm.Model
	//TODO_ID AUTO GENERATE
	ProjectID   uint   `json:"projectid" form:"projectid"`
	UserID      uint   `json:"UserID" form:"UserID"`
	Task        string `json:"task" form:"task"`
	Status      string `json:"status" form:"status"`
	Description string `json:"description" form:"description"`
}
