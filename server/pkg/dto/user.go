package dto

import (
	"mime/multipart"
	"time"
)

type UserRegisterBody struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
}

type UserLoginBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}

type ErrorAuthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Field   string `json:"field"`
}

type UserDetailBody struct {
	Username     string                `json:"username"`
	Firstname    string                `json:"firstname"`
	Lastname     string                `json:"lastname"`
	PhotoProfile *multipart.FileHeader `json:"photo_profile"`
	Banner       *multipart.FileHeader `json:"banner"`
}

type UserDetailResponse struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	Username     string    `json:"username"`
	Firstname    string    `json:"firstname"`
	Lastname     string    `json:"lastname"`
	PhotoProfile string    `json:"photo_profile"`
	Banner       string    `json:"banner"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `josn:"updated_at"`
}

type UserResponse struct {
	ID        string             `json:"id"`
	Email     string             `json:"email"`
	Detail    UserDetailResponse `json:"detail"`
	Token     string             `json:"token"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type FollowData struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type FollowersResponse struct {
	Count     int          `json:"count"`
	Followers []FollowData `json:"followers"`
}

type FollowingResponse struct {
	Count      int          `json:"count"`
	Followings []FollowData `json:"followings"`
}

type FollowedResponse struct {
	User string `json:"user"`
}

type Profile struct {
	ID             string    `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Firstname      string    `json:"firstname"`
	Lastname       string    `json:"lastname"`
	PhotoProfile   string    `json:"photo_profile"`
	Banner         string    `json:"banner"`
	FollowersCount int       `json:"followers_count"`
	FollowingCount int       `json:"following_count"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
