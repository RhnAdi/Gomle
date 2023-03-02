package CommentHandler

import (
	"github.com/RhnAdi/Gomle/pkg/domain"
)

type CommentHandler struct {
	service domain.CommentService
}

func NewCommentHandler(service domain.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}
