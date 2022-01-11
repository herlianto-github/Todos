package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//USERID AUTO GENERATE
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
	Todo     []To_Do
}
