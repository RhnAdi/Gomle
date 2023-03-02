package dto

import (
	"time"

	"github.com/RhnAdi/Gomle/pkg/models"
)

type AllPostResponse struct {
	Status string        `json:"status"`
	Data   []models.Post `json:"data"`
}

type AllMyPostResponse struct {
	Status string           `json:"status"`
	Data   []MyPostResponse `json:"data"`
}

type PostRequestBody struct {
	Content string         `json:"content"`
	Files   []models.Image `json:"files"`
}

type MyPostResponse struct {
	ID        string         `json:"id"`
	Content   string         `json:"content"`
	Files     []models.Image `json:"files"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type Post struct {
	ID        string           `json:"id"`
	UserID    string           `json:"user_id"`
	Files     []models.Image   `json:"files"`
	Content   string           `json:"content"`
	Comments  []models.Comment `json:"comments"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

type PostResponse struct {
	Status string `json:"status"`
	Data   Post   `json:"data"`
}

type FriendPostResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Content   string    `json:"content"`
	Files     []string  `json:"files"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeletePostResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}
