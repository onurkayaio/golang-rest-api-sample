package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" valid:"required~username is required.,length(6|12)~username length should be between 6-12."`
	Password string `json:"password" valid:"required~password is required."`
	Email    string `json:"email" valid:"email~email is not correct.,required~email is required."`
	Role     string `json:"role" valid:"required~role is required.,in(admin|test|user)~role should be admin|test|user."`
	Status   string `json:"status" valid:"required~status is required.,in(active|inactive)~status should be active|inactive."`
}

type Login struct {
	Username string `json:"username" valid:"required~username is required."`
	Password string `json:"password" valid:"required~username is required."`
}
