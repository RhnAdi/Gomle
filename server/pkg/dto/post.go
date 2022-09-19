package dto

import (
	"mime/multipart"
	"time"

	"github.com/RhnAdi/Gomle/pkg/models"
)

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

type PostResponse struct {
	ID        string           `json:"id"`
	UserID    string           `json:"user_id"`
	Files     []models.Image   `json:"files"`
	Content   string           `json:"content"`
	Comments  []models.Comment `json:"comments"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

type CommentBody struct {
	Text string                `json:"text"`
	File *multipart.FileHeader `json:"file"`
}

type CommentRequest struct {
	Text string `json:"text" binding:"required"`
	File string `json:"file"`
}

type FriendPostResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Content   string    `json:"content"`
	Files     []string  `json:"files"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
