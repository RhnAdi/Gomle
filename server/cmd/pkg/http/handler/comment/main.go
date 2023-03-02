package CommentHandler

import (
	"github.com/RhnAdi/Gomle/server/pkg/domain"
)

type CommentHandler struct {
	service domain.CommentService
}

func NewCommentHandler(service domain.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}
