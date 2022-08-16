package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type status int

const (
	coldCall  status = iota
	other
)

type Lead struct{
	gorm.Model
	Name string `json:"name"`
	LastName string  `json:"lastname"`
	Company string `json:"company"`
	Location string `json:"location"`
	IVA string `json:"iva"`
	PrivatePhone string `json:"privatephone"`
	PublicPhone string `json:"publicphone"`
	PrivateEmail string `json:"privateemail"`
	PublicEmail string `json:"publicemail"`
	Status status `json:"status"`
}
