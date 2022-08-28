package models

import (
	"time"

	"github.com/jinzhu/gorm"
)


type Project struct{
	gorm.Model
	Task string `json:"task"`
	State int `json:"state"`
	UserID int `json:"appointee"`
	Priority  int `json:"priority"`
	Due time.Time `json:"due"`
	Created time.Time `json:"creation"`

}