package app

import (
	"os"

	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/RhnAdi/Gomle/pkg/domain"
	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/RhnAdi/Gomle/pkg/models"
)

type commentService struct {
	db domain.CommentDB
}

func (s *commentService) AddComment(claim auth.JWTClaim, postId string, commentReq dto.CommentRequest) (comment models.Comment, err error) {
	comment, err = s.db.AddComment(models.Comment{
		PostID: postId,
		UserID: claim.ID,
		Text:   commentReq.Text,
		File:   commentReq.File,
	})
	return
}

func (s *commentService) FindComment(id string) (comment models.Comment, err error) {
	comment, err = s.db.FindComment(models.Comment{ID: id})
	return
}

func (s *commentService) UpdateComment(claim auth.JWTClaim, id string, commentReq dto.CommentRequest) (comment models.Comment, err error) {
	comment, err = s.FindComment(id)
	if err != nil {
		return
	}

	if claim.ID != comment.UserID {
		err = helper.ErrYouAreNotOwner
		return
	}

	if _, e := os.Stat("public/images/" + comment.File); e == nil {
		if err = os.Remove("public/images/" + comment.File); err != nil {
			return
		}
	}

	comment.Text = commentReq.Text
	comment.File = commentReq.File

	comment, err = s.db.UpdateComment(comment.ID, comment)

	return
}

func (s *commentService) DeleteComment(claim auth.JWTClaim, comment_id string, post_id string) (string, error) {
	comment, err := s.FindComment(comment_id)
	if err != nil {
		return comment_id, err
	}

	if comment.UserID != claim.ID {
		return comment_id, helper.ErrYouAreNotOwner
	}

	if _, e := os.Stat("public/images/" + comment.File); e == nil {
		if err = os.Remove("public/images/" + comment.File); err != nil {
			return comment_id, err
		}
	}

	comment_id, err = s.db.DeleteComment(models.Post{ID: post_id}, models.Comment{ID: comment_id})

	return comment_id, err
}
