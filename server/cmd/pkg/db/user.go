package db

import (
	"github.com/RhnAdi/Gomle/server/pkg/dto"
	"github.com/RhnAdi/Gomle/server/pkg/models"
	"gorm.io/gorm"
)

type userDB struct {
	db *gorm.DB
}

func (u *userDB) FindAll() (users []models.User, err error) {
	err = u.db.Find(&users).Error
	return
}

func (u *userDB) Find(user models.User) (models.User, error) {
	err := u.db.Preload("Profile").First(&user).Error
	return user, err
}

func (u *userDB) FindByEmail(email string) (user models.User, err error) {
	user.Email = email
	err = u.db.First(&user).Error
	return
}

func (u *userDB) Create(req models.User) (models.User, error) {
	err := u.db.Save(&req).Error
	return req, err
}

func (u *userDB) Update(req models.User) (user models.User, err error) {
	err = u.db.Save(&req).Error
	return
}

func (u *userDB) Delete(req models.User) (user models.User, err error) {
	err = u.db.Delete(&req).Error
	return
}

func (u *userDB) Follow(req models.User, to models.User) (models.User, error) {
	err := u.db.Model(&req).Association("Followings").Append([]models.User{to})
	return to, err
}

func (u *userDB) Followers(req models.User) (followers []dto.FollowData, err error) {
	tx := u.db.Model(&models.User{}).Select(
		"users.id id",
		"users.email email",
		"user_details.username username",
		"user_details.firstname firstname",
		"user_details.lastname lastname",
		"user_details.photo_profile photo_profile",
	).Joins(
		"INNER JOIN user_relation ON user_relation.user_id = users.id",
	).Joins(
		"INNER JOIN user_details ON user_details.user_id = users.id",
	).Where(
		"user_relation.following_id = ?", req.ID,
	).Scan(&followers)
	if tx.Error != nil {
		err = tx.Error
	}

	return
}

func (u *userDB) Followings(req models.User) (followings []dto.FollowData, err error) {
	tx := u.db.Model(&models.User{}).Select(
		"users.id",
		"users.email",
		"user_details.username",
	).Joins(
		"INNER JOIN user_relation ON user_relation.following_id = users.id",
	).Joins(
		"INNER JOIN user_details ON user_details.user_id = users.id",
	).Where(
		"user_relation.user_id = ?", req.ID,
	).Scan(&followings)

	if tx.Error != nil {
		err = tx.Error
	}

	return
}

func NewUserDB(db *gorm.DB) *userDB {
	return &userDB{db: db}
}
