package domain

import "github.com/RhnAdi/Gomle/server/pkg/models"

type Profile struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	PhotoProfile string `json:"photo_profile"`
	Banner       string `json:"banner"`
}

type ProfileDB interface {
	Find(models.Profile) (models.Profile, error)
	Create(models.Profile) (models.Profile, error)
	Update(models.Profile) (models.Profile, error)
	UpdatePhotoProfile(models.Profile) (string, error)
	UpdateBanner(models.Profile) (string, error)
	Delete(models.Profile) (models.Profile, error)
}
