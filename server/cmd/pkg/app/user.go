package app

import (
	"os"
	"time"

	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/RhnAdi/Gomle/pkg/domain"
	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/models"
)

type userService struct {
	UserDB    domain.UserDB
	ProfileDB domain.ProfileDB
}

func (s *userService) Login(req dto.UserLoginBody) (token string, err error) {
	var user models.User
	user, err = s.FindByEmail(req.Email)
	if err == nil {
		err = user.ComparePassword(req.Password)
		if err == nil {
			token, err = auth.GenerateJWT(user.ID, user.Email)
			return
		}
		return
	}
	return
}

func (s *userService) Register(req dto.UserRegisterBody) (token string, err error) {
	var data models.User
	data, err = s.Create(domain.User{
		Email:    req.Email,
		Password: req.Password,
		Profile: domain.Profile{
			Username:  req.Username,
			Firstname: req.Firstname,
			Lastname:  req.Lastname,
		},
	})
	if err != nil {
		return
	}
	token, err = auth.GenerateJWT(data.ID, data.Email)
	return
}

func (s *userService) FindAll() (users []models.User, err error) {
	users, err = s.UserDB.FindAll()
	return
}

func (s *userService) Find(id string) (user models.User, err error) {
	user, err = s.UserDB.Find(models.User{ID: id})
	return
}

func (s *userService) FindProfile(id string) (data models.Profile, err error) {
	data, err = s.ProfileDB.Find(models.Profile{UserID: id})
	return
}

func (s *userService) FindByEmail(email string) (user models.User, err error) {
	user, err = s.UserDB.FindByEmail(email)
	return
}

func (s *userService) Create(req domain.User) (user models.User, err error) {
	// User
	user = models.User{
		Email:    req.Email,
		Password: req.Password,
		Profile: models.Profile{
			Username:  req.Profile.Username,
			Firstname: req.Profile.Firstname,
			Lastname:  req.Profile.Lastname,
		},
	}
	err = user.HashPassword()
	if err != nil {
		return
	}

	user, err = s.UserDB.Create(user)
	if err != nil {
		return
	}

	return
}

func (s *userService) UpdateProfile(claim auth.JWTClaim, userReq domain.Profile) (profile models.Profile, err error) {
	profile, err = s.ProfileDB.Find(models.Profile{UserID: claim.ID})
	if err != nil {
		return models.Profile{}, err
	}
	if claim.ID != profile.UserID {
		return models.Profile{}, err
	}
	if userReq.Username != "" {
		profile.Username = userReq.Username
	}
	if userReq.Firstname != "" {
		profile.Firstname = userReq.Firstname
	}
	if userReq.Lastname != "" {
		profile.Lastname = userReq.Lastname
	}
	if userReq.PhotoProfile != "" {
		if _, e := os.Stat("../../public/images/" + profile.PhotoProfile); e == nil {
			if err = os.Remove("../../public/images/" + profile.PhotoProfile); err != nil {
				return profile, err
			}
		}
		profile.PhotoProfile = userReq.PhotoProfile
	}
	if userReq.Banner != "" {
		if _, e := os.Stat("../../public/images/" + profile.Banner); e == nil {
			if err = os.Remove("../../public/images/" + profile.Banner); err != nil {
				return profile, err
			}
		}
		profile.Banner = userReq.Banner
	}
	profile, err = s.ProfileDB.Update(profile)

	return
}

func (s *userService) UpdatePhotoProfile(claim auth.JWTClaim, filename string) (updatedFilename string, err error) {
	user_detail, err := s.ProfileDB.Find(models.Profile{UserID: claim.ID})
	if err != nil {
		return "", err
	}

	if _, e := os.Stat("../../public/images/" + user_detail.PhotoProfile); e == nil {
		if err = os.Remove("../../public/images/" + user_detail.PhotoProfile); err != nil {
			return "", err
		}
	}
	user_detail.PhotoProfile = filename
	updatedFilename, err = s.ProfileDB.UpdatePhotoProfile(user_detail)
	return
}

func (s *userService) UpdateBanner(claim auth.JWTClaim, filename string) (updatedFilename string, err error) {
	user_detail, err := s.ProfileDB.Find(models.Profile{UserID: claim.ID})
	if err != nil {
		return "", err
	}
	if _, e := os.Stat("../../public/images/" + user_detail.Banner); e == nil {
		if err = os.Remove("../../public/images/" + user_detail.Banner); err != nil {
			return "", err
		}
	}
	user_detail.Banner = filename
	updatedFilename, err = s.ProfileDB.UpdateBanner(user_detail)
	return
}

func (s *userService) Delete(id string) (user models.User, err error) {
	user, err = s.UserDB.Find(models.User{ID: id})
	if err != nil {
		return
	}
	user, err = s.UserDB.Delete(user)
	return
}

func (s *userService) UpdateEmail(claim auth.JWTClaim, id string, email string) (user models.User, err error) {
	user, err = s.UserDB.Find(models.User{ID: id})
	if err != nil {
		return models.User{}, err
	}
	if claim.ID != user.ID {
		return models.User{}, err
	}
	user.Email = email
	user.UpdatedAt = time.Now()
	user, err = s.UserDB.Update(user)
	return
}

func (s *userService) Follow(claim auth.JWTClaim, following_id string) (data models.User, err error) {
	_, err = s.UserDB.Find(models.User{ID: following_id})
	if err != nil {
		return data, err
	}
	data, err = s.UserDB.Follow(models.User{ID: claim.ID}, models.User{ID: following_id})
	return data, err
}

func (s *userService) Followers(claim auth.JWTClaim) ([]dto.FollowData, error) {
	data, err := s.UserDB.Followers(models.User{
		ID:    claim.ID,
		Email: claim.Email,
	})
	if err != nil {
		return []dto.FollowData{}, err
	}

	return data, nil
}

func (s *userService) Followings(claim auth.JWTClaim) ([]dto.FollowData, error) {
	data, err := s.UserDB.Followings(models.User{
		ID:    claim.ID,
		Email: claim.Email,
	})
	if err != nil {
		return []dto.FollowData{}, err
	}

	return data, nil
}

func NewUserService(UserDB domain.UserDB, ProfileDB domain.ProfileDB) *userService {
	return &userService{UserDB: UserDB, ProfileDB: ProfileDB}
}
