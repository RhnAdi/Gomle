package db

import (
	"github.com/RhnAdi/Gomle/pkg/models"
	"gorm.io/gorm"
)

type ProfileDB struct {
	db *gorm.DB
}

func (u *ProfileDB) Find(profile models.Profile) (models.Profile, error) {
	err := u.db.First(&profile, "user_id=?", profile.UserID).Error
	return profile, err
}

func (u *ProfileDB) Create(profile models.Profile) (models.Profile, error) {
	err := u.db.Create(&profile).Error
	return profile, err
}

func (u *ProfileDB) Update(profile models.Profile) (models.Profile, error) {
	err := u.db.Save(&profile).Error
	return profile, err
}

func (u *ProfileDB) UpdatePhotoProfile(profile models.Profile) (string, error) {
	err := u.db.Model(&profile).Update("photo_profile", &profile.PhotoProfile).Error
	return profile.PhotoProfile, err
}

func (u *ProfileDB) UpdateBanner(profile models.Profile) (string, error) {
	err := u.db.Model(&profile).Update("banner", &profile.Banner).Error
	return profile.Banner, err
}

func (u *ProfileDB) Delete(profile models.Profile) (models.Profile, error) {
	err := u.db.Delete(&profile).Error
	return profile, err
}

func NewProfileDB(db *gorm.DB) *ProfileDB {
	return &ProfileDB{db: db}
}
