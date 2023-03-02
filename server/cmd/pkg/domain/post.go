package domain

import (
	"github.com/RhnAdi/Gomle/server/internal/auth"
	"github.com/RhnAdi/Gomle/server/pkg/dto"
	"github.com/RhnAdi/Gomle/server/pkg/models"
)

type Post struct {
	ID      string         `json:"id"`
	Content string         `json:"content"`
	Files   []models.Image `json:"files"`
}

type PostService interface {
	FindAll() ([]models.Post, error)
	Find(id string) (models.Post, error)
	FindMyPost(claim auth.JWTClaim) ([]dto.MyPostResponse, error)
	Create(claim auth.JWTClaim, post Post) (models.Post, error)
	Update(claim auth.JWTClaim, post Post) (models.Post, error)
	Delete(claim auth.JWTClaim, post Post) (models.Post, error)
	FollowingPosts(claim auth.JWTClaim) ([]models.Post, error)
}

type PostDB interface {
	FindAll() ([]models.Post, error)
	Find(models.Post) (models.Post, error)
	FindMyPost(id string) ([]models.Post, error)
	Create(models.Post) (models.Post, error)
	Update(models.Post) (models.Post, error)
	Delete(models.Post) (models.Post, error)
	FollowingPosts(userId string) ([]models.Post, error)
}
