package db

import (
	"github.com/RhnAdi/Gomle/pkg/models"
	"gorm.io/gorm"
)

type userDetailDB struct {
	db *gorm.DB
}

func (u *userDetailDB) Find(user_detail models.UserDetail) (models.UserDetail, error) {
	err := u.db.First(&user_detail, "user_id=?", user_detail.UserID).Error
	return user_detail, err
}

func (u *userDetailDB) Create(user_detail models.UserDetail) (models.UserDetail, error) {
	err := u.db.Create(&user_detail).Error
	return user_detail, err
}

func (u *userDetailDB) Update(user_detail models.UserDetail) (models.UserDetail, error) {
	err := u.db.Save(&user_detail).Error
	return user_detail, err
}

func (u *userDetailDB) UpdatePhotoProfile(user_detail models.UserDetail) (string, error) {
	err := u.db.Model(&user_detail).Update("photo_profile", &user_detail.PhotoProfile).Error
	return user_detail.PhotoProfile, err
}

func (u *userDetailDB) UpdateBanner(user_detail models.UserDetail) (string, error) {
	err := u.db.Model(&user_detail).Update("banner", &user_detail.Banner).Error
	return user_detail.Banner, err
}

func (u *userDetailDB) Delete(user_detail models.UserDetail) (models.UserDetail, error) {
	err := u.db.Delete(&user_detail).Error
	return user_detail, err
}

func NewUserDetailDB(db *gorm.DB) *userDetailDB {
	return &userDetailDB{db: db}
}
