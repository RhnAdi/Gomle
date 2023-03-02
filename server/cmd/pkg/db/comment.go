package db

import (
	"github.com/RhnAdi/Gomle/server/pkg/models"
	"gorm.io/gorm"
)

type commentDB struct {
	db *gorm.DB
}

func (r *commentDB) FindComment(comment models.Comment) (models.Comment, error) {
	err := r.db.Find(&comment).Error
	return comment, err
}

func (r *commentDB) AddComment(comment models.Comment) (models.Comment, error) {
	err := r.db.Model(&models.Post{ID: comment.PostID}).Association("Comments").Append(&comment)
	return comment, err
}

func (r *commentDB) UpdateComment(comment models.Comment) (models.Comment, error) {
	err := r.db.Save(&comment).Error
	return comment, err
}

func (r *commentDB) DeleteComment(post models.Post, comment models.Comment) (string, error) {
	err := r.db.Model(&post).Association("Comment").Delete(&comment)
	if err != nil {
		return comment.ID, err
	}
	err = r.db.Delete(&comment).Error
	return comment.ID, err
}
