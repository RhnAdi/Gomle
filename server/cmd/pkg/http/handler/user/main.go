package UserHandler

import "github.com/RhnAdi/Gomle/server/pkg/domain"

type UserHandler struct {
	service domain.UserService
}

func NewUserHandler(service domain.UserService) *UserHandler {
	return &UserHandler{service: service}
}
