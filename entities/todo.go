package entities

import "gorm.io/gorm"

type ToDo struct {
	gorm.Model
	//TODO_ID AUTO GENERATE
	ID          uint
	ProjectID   *uint  `json:"projectid" form:"projectid"`
	UserID      uint   `json:"userid" form:"userid"`
	Task        string `json:"task" form:"task"`
	Status      string `json:"status" form:"status" gorm:"default: In Progress"`
	Description string `json:"description" form:"description"`
}
