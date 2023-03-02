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

// My Account
// @Tags accounts
// @Summary Get user account information
// @Description Need auth token in header to call this endpoint
// @Accept json
// @Produce json
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @Success 200 {object} dto.Account
// @Failure 404 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /users/account/ [GET]
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
	c.JSON(http.StatusOK, dto.Account{
		Status: "success",
		Data: dto.Profile{
			ID:             data.ID,
			Username:       data.Profile.Username,
			Email:          data.Email,
			Firstname:      data.Profile.Username,
			Lastname:       data.Profile.Lastname,
			PhotoProfile:   data.Profile.PhotoProfile,
			Banner:         data.Profile.Banner,
			FollowersCount: len(data.Followers),
			FollowingCount: len(data.Followings),
			CreatedAt:      data.Profile.CreatedAt,
			UpdatedAt:      data.Profile.UpdatedAt,
		},
	})
}
