package PostHandler

import "github.com/RhnAdi/Gomle/server/pkg/domain"

type PostHandler struct {
	service domain.PostService
}

func NewPostHandler(service domain.PostService) *PostHandler {
	return &PostHandler{service: service}
}
