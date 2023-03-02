package UserHandler

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/RhnAdi/Gomle/pkg/domain"
	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// @Summary update account profile
// @Description need auth token in header for update account profile
// @Tags accounts
// @Accept json
// @Produce json
// @Param user body dto.UserDetailBody true "update profile"
// @Success 200 {object} dto.UserDetailResponse
// @Failure 403 {object} helper.ErrorResponse
// @Router /users/profile/ [PUT]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	var body dto.UserDetailBody
	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
			"error":   err.Error(),
		})
	}
	claim := c.MustGet("claim").(auth.JWTClaim)

	photoProfileExt := filepath.Ext(body.PhotoProfile.Filename)
	bannerExt := filepath.Ext(body.Banner.Filename)

	newPhotoProfileFilename := uuid.New().String() + photoProfileExt
	newBannerFilename := uuid.New().String() + bannerExt

	err = c.SaveUploadedFile(body.PhotoProfile, "public/images/"+newPhotoProfileFilename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Status:  "failed",
			Message: "upload image failed",
			Field:   "photo_profile",
			Error:   err.Error(),
		})
		return
	}

	err = c.SaveUploadedFile(body.Banner, "public/images/"+newBannerFilename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Status:  "failed",
			Message: "upload image failed",
			Field:   "banner",
			Error:   err.Error(),
		})
		return
	}

	data, err := h.service.UpdateProfile(claim, domain.Profile{
		Username:     body.Username,
		Firstname:    body.Lastname,
		Lastname:     body.Lastname,
		PhotoProfile: newPhotoProfileFilename,
		Banner:       newBannerFilename,
	})

	if err != nil {
		if e := os.Remove("public/images/" + newPhotoProfileFilename); e != nil {
			c.JSON(http.StatusNotFound, helper.ErrorResponse{
				Status:  "failed",
				Message: "can't upload image: 124",
				Field:   "photo_profile",
				Error:   e.Error(),
			})
			return
		}

		if e := os.Remove("public/images/" + newBannerFilename); e != nil {
			c.JSON(http.StatusNotFound, helper.ErrorResponse{
				Status:  "failed",
				Message: "can't upload image: 124",
				Field:   "banner",
				Error:   e.Error(),
			})
			return
		}

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			c.JSON(http.StatusNotFound, helper.ErrorResponse{
				Status:  "failed",
				Message: "profile not found",
				Field:   "id",
				Error:   err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, dto.UserDetailResponse{
		Status: "update success",
		Data: dto.UserDetailInfo{
			ID:        data.ID,
			UserID:    data.UserID,
			Username:  data.Username,
			Firstname: data.Firstname,
			Lastname:  data.Lastname,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		},
	})
}
