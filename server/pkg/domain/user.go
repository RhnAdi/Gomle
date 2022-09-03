package domain

import (
	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/models"
)

type User struct {
	Email      string `json:"email"`
	Password   string `json:"-"`
	UserDetail UserDetail
}

type UserRelation struct {
	Count int `json:"count"`
	Data  []dto.FollowData
}

type UserDB interface {
	FindAll() ([]models.User, error)
	Find(models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
	Create(models.User) (models.User, error)
	Update(models.User) (models.User, error)
	Delete(models.User) (models.User, error)
	Follow(req models.User, to models.User) (models.User, error)
	Followers(models.User) ([]dto.FollowData, error)
	Followings(models.User) ([]dto.FollowData, error)
}

type UserService interface {
	Register(req dto.UserRegisterBody) (token string, err error)
	Login(req dto.UserLoginBody) (token string, err error)
	FindAll() ([]models.User, error)
	Find(id string) (models.User, error)
	FindProfile(id string) (models.UserDetail, error)
	FindByEmail(email string) (models.User, error)
	Create(User) (models.User, error)
	UpdateEmail(claim auth.JWTClaim, id string, email string) (models.User, error)
	UpdateUserDetail(claim auth.JWTClaim, userReq UserDetail) (models.UserDetail, error)
	Delete(string) (models.User, error)
	Follow(claim auth.JWTClaim, following_id string) (models.User, error)
	Followers(claim auth.JWTClaim) ([]dto.FollowData, error)
	Followings(claim auth.JWTClaim) ([]dto.FollowData, error)
	UpdatePhotoProfile(claim auth.JWTClaim, filename string) (string, error)
	UpdateBanner(claim auth.JWTClaim, filename string) (string, error)
}
