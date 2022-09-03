package domain

import "github.com/RhnAdi/Gomle/pkg/models"

type UserDetail struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	PhotoProfile string `json:"photo_profile"`
	Banner       string `json:"banner"`
}

type UserDetailDB interface {
	Find(models.UserDetail) (models.UserDetail, error)
	Create(models.UserDetail) (models.UserDetail, error)
	Update(models.UserDetail) (models.UserDetail, error)
	UpdatePhotoProfile(models.UserDetail) (string, error)
	UpdateBanner(models.UserDetail) (string, error)
	Delete(models.UserDetail) (models.UserDetail, error)
}
