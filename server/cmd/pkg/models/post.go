package models

import "time"

type Post struct {
	ID        string    `json:"id" gorm:"default:uuid_generate_v4();primaryKey"`
	UserID    string    `json:"user_id" gorm:"foreignkey"`
	User      User      `json:"-"`
	Content   string    `json:"content"`
	Files     []Image   `json:"files" gorm:"foreignKey:PostID"`
	Comments  []Comment `json:"comments"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
