package dto

import (
	"mime/multipart"
	"time"

	"github.com/RhnAdi/Gomle/pkg/models"
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

type UserDetailInfo struct {
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

type UserDetailResponse struct {
	Status string         `json:"status"`
	Data   UserDetailInfo `json:"data"`
}

type UserResponse struct {
	ID        string         `json:"id"`
	Email     string         `json:"email"`
	Detail    UserDetailInfo `json:"detail"`
	Token     string         `json:"token"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type FollowData struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Followers struct {
	Count     int          `json:"count"`
	Followers []FollowData `json:"followers"`
}

type FollowersResponse struct {
	Status string    `json:"status"`
	Data   Followers `json:"data"`
}

type Following struct {
	Count      int          `json:"count"`
	Followings []FollowData `json:"followings"`
}

type FollowingsResponse struct {
	Status string    `json:"status"`
	Data   Following `json:"data"`
}

type FollowedResponse struct {
	Status   string `json:"status"`
	Followed string `json:"followed"`
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

type UserProfile struct {
	Status string         `json:"status"`
	Data   models.Profile `json:"data"`
}

type Account struct {
	Status string  `json:"status"`
	Data   Profile `json:"data"`
}
type UploadFileResponse struct {
	Status   string `json:"status"`
	Filename string `json:"filename"`
	Field    string `json:"field"`
}
