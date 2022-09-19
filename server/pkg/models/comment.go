package models

type Comment struct {
	ID     string `json:"id" gorm:"default:uuid_generate_v4();primaryKey"`
	PostID string `json:"post_id"`
	UserID string `json:"user_id"`
	Text   string `json:"text"`
	File   string `json:"file"`
}
