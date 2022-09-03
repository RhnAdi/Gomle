package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         string     `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	Email      string     `json:"email" gorm:"unique;primaryKey"`
	Password   string     `json:"-"`
	UserDetail UserDetail `json:"user_detail"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	Followings []User     `json:"-" gorm:"many2many:user_relation;joinForeignKey:user_Id;JoinReferences:following_id"`
	Followers  []User     `json:"-" gorm:"many2many:user_relation;joinForeignKey:following_id;JoinReferences:user_Id"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 16)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
