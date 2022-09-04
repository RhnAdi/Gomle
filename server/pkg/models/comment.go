package models

type Comment struct {
	ID     string `json:"id" gorm:"default:uuid_generate_v4();primaryKey"`
	PostId string `json:"post_id"`
	UserId string `json:"user_id"`
	Text   string `json:"text"`
	File   string `json:"file"`
}
