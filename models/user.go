package models

import "github.com/jinzhu/gorm"

type User struct{
	gorm.Model
	Name string `json:"name" gorm:"unique"`
	Email string  `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}