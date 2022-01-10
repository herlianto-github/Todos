package entities

import "gorm.io/gorm"

type Projects struct {
	gorm.Model
	//PROJECTID AUTO GENERATE
	UserID      uint   `json:"userid" form:"userid"`
	ProjectName string `json:"projectname" form:"projectname"`
}
