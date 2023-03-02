package domain

import (
	"github.com/RhnAdi/Gomle/server/internal/auth"
	"github.com/RhnAdi/Gomle/server/pkg/dto"
	"github.com/RhnAdi/Gomle/server/pkg/models"
)

type CommentService interface {
	FindComment(id string) (models.Comment, error)
	AddComment(claim auth.JWTClaim, post_id string, comment dto.CommentRequest) (models.Comment, error)
	UpdateComment(claim auth.JWTClaim, id string, comment dto.CommentRequest) (models.Comment, error)
	DeleteComment(claim auth.JWTClaim, comment_id string, post_id string) (string, error)
}

type CommentDB interface {
	FindComment(models.Comment) (models.Comment, error)
	AddComment(comment models.Comment) (models.Comment, error)
	UpdateComment(id string, comment models.Comment) (models.Comment, error)
	DeleteComment(models.Post, models.Comment) (string, error)
}
