package UserHandler

import (
	"errors"
	"net/http"

	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *UserHandler) MyAccount(c *gin.Context) {
	claim := c.MustGet("claim").(auth.JWTClaim)
	data, err := h.service.Find(claim.ID)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, helper.ErrorResponse{
			Status:  "failed",
			Message: "profile not fount",
			Field:   "id",
			Error:   err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Status:  "failed",
			Message: "can't get profile",
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": dto.Profile{
			ID:             data.ID,
			Username:       data.UserDetail.Username,
			Email:          data.Email,
			Firstname:      data.UserDetail.Username,
			Lastname:       data.UserDetail.Lastname,
			PhotoProfile:   data.UserDetail.PhotoProfile,
			Banner:         data.UserDetail.Banner,
			FollowersCount: len(data.Followers),
			FollowingCount: len(data.Followings),
			CreatedAt:      data.UserDetail.CreatedAt,
			UpdatedAt:      data.UserDetail.UpdatedAt,
		},
	})
}
