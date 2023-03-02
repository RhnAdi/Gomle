package db

import (
	"github.com/RhnAdi/Gomle/server/pkg/models"
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
	if err := r.db.Model(&models.Post{ID: post.ID}).Association("Files").Clear(); err != nil {
		return post, err
	}

	if err := r.db.Delete(&models.Image{PostID: post.ID}).Error; err != nil {
		return post, err
	}

	err := r.db.Save(&post).Error

	return post, err
}

func (r *postDB) Delete(post models.Post) (models.Post, error) {
	if err := r.db.Delete(&post).Error; err != nil {
		return post, err
	}
	err := r.db.Delete(&models.Image{PostID: post.ID}).Error
	return post, err
}

func (r *postDB) FollowingPosts(userId string) ([]models.Post, error) {
	var data []models.Post

	err := r.db.Model(&models.Post{}).Preload("Files").Raw(
		`SELECT * FROM ( SELECT "posts"."id","posts"."user_id","posts"."content","posts"."created_at","posts"."updated_at" FROM "posts" WHERE posts.user_id = ? UNION ALL SELECT "posts"."id","posts"."user_id","posts"."content","posts"."created_at","posts"."updated_at" FROM "posts" INNER JOIN user_relation ON posts.user_id = user_relation.following_id WHERE user_relation.user_id = ?) dum ORDER BY dum.updated_at DESC;`,
		userId,
		userId,
	).Find(&data).Error
	// err := r.db.Model(&models.Post{}).Preload("Files").Joins(
	// 	"INNER JOIN user_relation ON posts.user_id = user_relation.following_id AND post.user_id = ?",
	// 	userId,
	// ).Joins("RIGHT JOIN posts myPost ON myPost.user_id = ?", userId).Where("user_relation.user_id = ?", userId).Find(&data).Error

	return data, err
}

func NewPostDB(db *gorm.DB) *postDB {
	return &postDB{db: db}
}
