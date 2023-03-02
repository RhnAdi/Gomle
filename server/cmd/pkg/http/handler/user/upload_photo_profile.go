package UserHandler

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/RhnAdi/Gomle/server/internal/auth"
	"github.com/RhnAdi/Gomle/server/pkg/dto"
	"github.com/RhnAdi/Gomle/server/pkg/http/helper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// @Summary upload photo profile
// @Description need auth token in header for update upload photo profile
// @Tags accounts
// @Accept multipart/form-data
// @Produce json
// @Param photo_profile formData file true "upload photo profile"
// @in formData
// @type file
// @Success 200 {object} dto.UploadFileResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 404 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /users/upload/photo_profile [PUT]
func (h *UserHandler) UploadPhotoProfile(c *gin.Context) {
	claim := c.MustGet("claim").(auth.JWTClaim)

	file, _ := c.FormFile("photo_profile")
	extension := filepath.Ext(file.Filename)
	filetype := file.Header.Get("Content-Type")

	if !(filetype == "image/jpeg" || filetype == "image/png" || filetype == "image/webp") {
		c.JSON(http.StatusBadRequest, helper.ErrorResponse{
			Status:  "failed",
			Message: "file type not allowed",
			Field:   "photo_profile",
			Error:   "only image/jpeg or image/png or image/wwebp",
		})
		return
	}

	newFilename := uuid.New().String() + extension

	err := c.SaveUploadedFile(file, "public/images/"+newFilename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Status:  "failed",
			Message: "can't upload image",
			Field:   "photo_profile",
			Error:   err.Error(),
		})
		return
	}

	data, err := h.service.UpdatePhotoProfile(claim, newFilename)
	if err != nil {
		if e := os.Remove("public/images/" + newFilename); e != nil {
			c.JSON(http.StatusNotFound, helper.ErrorResponse{
				Status:  "failed",
				Message: "can't upload image: 124",
				Field:   "photo_profile",
				Error:   e.Error(),
			})
			return
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, helper.ErrorResponse{
				Status:  "failed",
				Message: "can't upload image",
				Field:   "photo_profile",
				Error:   err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Status:  "failed",
			Message: "can't update photo profile",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.UploadFileResponse{
		Status:   "success",
		Field:    "photo_profile",
		Filename: data,
	})
}
