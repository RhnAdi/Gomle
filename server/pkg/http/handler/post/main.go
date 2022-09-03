package PostHandler

import "github.com/RhnAdi/Gomle/pkg/domain"

type PostHandler struct {
	service domain.PostService
}

func NewPostHandler(service domain.PostService) *PostHandler {
	return &PostHandler{service: service}
}
