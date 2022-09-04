package db

import (
	"github.com/RhnAdi/Gomle/pkg/models"
	"gorm.io/gorm"
)

type postDB struct {
	db *gorm.DB
}

func (r *postDB) FindAll() ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Model(&models.Post{}).Preload("Files").Find(&posts).Error
	return posts, err
}

func (r *postDB) Find(post models.Post) (models.Post, error) {
	err := r.db.Model(&models.Post{}).Preload("Files").Preload("Comments").First(&post).Error
	return post, err
}

func (r *postDB) FindMyPost(userID string) ([]models.Post, error) {
	var myPosts []models.Post
	err := r.db.Model(&models.Post{}).Preload("Files").Find(&myPosts, "user_id=?", userID).Error
	return myPosts, err
}

func (r *postDB) Create(post models.Post) (models.Post, error) {
	err := r.db.Create(&post).Error
	return post, err
}

func (r *postDB) Update(post models.Post) (models.Post, error) {
	err := r.db.Save(&post).Error
	return post, err
}

func (r *postDB) Delete(post models.Post) (models.Post, error) {
	err := r.db.Delete(&post).Error
	return post, err
}

func (r *postDB) FollowingPosts(userId string) ([]models.Post, error) {
	var data []models.Post

	// Slow Database Request
	// err := r.db.Model(&models.Post{}).Select(
	// 	"posts.user_id",
	// 	"posts.id",
	// 	"posts.content",
	// 	"posts.created_at",
	// 	"posts.updated_at",
	// 	"user_details.username",
	// ).Joins(
	// 	"INNER JOIN user_relation ON posts.user_id = user_relation.following_id",
	// ).Joins(
	// 	"INNER JOIN user_details ON user_details.user_id = posts.user_id",
	// ).Where("user_relation.user_id = ?", userId).Find(&data).Error

	err := r.db.Model(&models.Post{}).Preload("Files").Joins(
		"INNER JOIN user_relation ON posts.user_id = user_relation.following_id",
	).Where("user_relation.user_id = ?", userId).Find(&data).Error

	return data, err
}

func (r *postDB) AddComment(comment models.Comment) (models.Comment, error) {
	err := r.db.Model(&models.Post{ID: comment.PostId}).Association("Comments").Append(&comment)
	return comment, err
}

func NewPostDB(db *gorm.DB) *postDB {
	return &postDB{db: db}
}
