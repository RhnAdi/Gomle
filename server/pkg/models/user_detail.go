package models

import (
	"gorm.io/gorm"
)

type UserDetail struct {
	gorm.Model
	ID           string `json:"id" gorm:"default:uuid_generate_v4();primaryKey"`
	UserID       string `gorm:"unique"`
	Username     string `json:"username" gorm:"unique"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	PhotoProfile string `json:"photo_profile"`
	Banner       string `json:"banner"`
}
