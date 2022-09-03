package models

import "time"

type Post struct {
	ID        string    `json:"id" gorm:"default:uuid_generate_v4()"`
	UserID    string    `json:"user_id" gorm:"foreignkey"`
	User      User      `json:"-"`
	Files     []Image   `json:"files"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
