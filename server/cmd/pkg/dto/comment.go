package dto

import "github.com/RhnAdi/Gomle/server/pkg/models"

type CommentRequest struct {
	Text  string `json:"text" binding:"required"`
	File  string `json:"file"`
	Reply bool   `json:"reply"`
}

type CommentResponse struct {
	Status string         `json:"status"`
	Data   models.Comment `json:"data"`
}
