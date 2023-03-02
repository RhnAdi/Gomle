package models

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	UserID       string
	ID           string    `json:"id" gorm:"default:uuid_generate_v4();primaryKey"`
	Username     string    `json:"username" gorm:"unique"`
	Firstname    string    `json:"firstname"`
	Lastname     string    `json:"lastname"`
	PhotoProfile string    `json:"photo_profile"`
	Banner       string    `json:"banner"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    gorm.DeletedAt
}
