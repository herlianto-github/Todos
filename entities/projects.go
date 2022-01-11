package entities

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	//PROJECTID AUTO GENERATE
	ProjectName string `json:"projectname" form:"projectname"`
	Todo        []To_Do
}
