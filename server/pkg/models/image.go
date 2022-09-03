package models

import "time"

type Image struct {
	ID        string    `json:"id" gorm:"default:uuid_generate_v4()"`
	PostId    string    `json:"post_id"`
	Filename  string    `json:"filename"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
