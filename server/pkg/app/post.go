package app

import (
	"time"

	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/RhnAdi/Gomle/pkg/domain"
	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/RhnAdi/Gomle/pkg/models"
)

type postService struct {
	PostDB domain.PostDB
}

func (s *postService) FindAll() (posts []models.Post, err error) {
	posts, err = s.PostDB.FindAll()
	return
}

func (s *postService) Find(id string) (post models.Post, err error) {
	post, err = s.PostDB.Find(models.Post{ID: id})
	return
}

func (s *postService) FindMyPost(claim auth.JWTClaim) ([]dto.MyPostResponse, error) {
	myPosts, err := s.PostDB.FindMyPost(claim.ID)
	var data []dto.MyPostResponse
	for _, post := range myPosts {
		data = append(data, dto.MyPostResponse{
			ID:        post.ID,
			Content:   post.Content,
			Files:     post.Files,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}
	return data, err
}

func (s *postService) Create(claim auth.JWTClaim, postReq domain.Post) (post models.Post, err error) {
	post, err = s.PostDB.Create(models.Post{
		Content: postReq.Content,
		UserID:  claim.ID,
		Files:   postReq.Files,
	})
	return
}

func (s *postService) Update(claim auth.JWTClaim, postReq domain.Post) (post models.Post, err error) {
	post, err = s.PostDB.Find(models.Post{ID: postReq.ID})
	if err != nil {
		return
	}
	if claim.ID != post.UserID {
		return models.Post{}, helper.ErrYouAreNotOwner
	}
	post, err = s.PostDB.Update(models.Post{
		ID:        post.ID,
		UserID:    post.UserID,
		Content:   postReq.Content,
		UpdatedAt: time.Now(),
		CreatedAt: post.CreatedAt,
	})
	return
}

func (s *postService) Delete(claim auth.JWTClaim, postReq domain.Post) (post models.Post, err error) {
	post, err = s.PostDB.Find(models.Post{ID: postReq.ID})
	if err != nil {
		return
	}
	if claim.ID != post.UserID {
		return models.Post{}, helper.ErrYouAreNotOwner
	}
	post, err = s.PostDB.Delete(post)
	return
}

func (s *postService) FollowingPosts(claim auth.JWTClaim) (posts []models.Post, err error) {
	posts, err = s.PostDB.FollowingPosts(claim.ID)
	return
}

func NewPostService(postDB domain.PostDB) *postService {
	return &postService{PostDB: postDB}
}
