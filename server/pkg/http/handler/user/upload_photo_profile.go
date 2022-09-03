package UserHandler

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

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

	err := c.SaveUploadedFile(file, "../../public/images/"+newFilename)
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
		if e := os.Remove("../../public.images/" + newFilename); e != nil {
			c.JSON(http.StatusNotFound, helper.ErrorResponse{
				Status:  "failed",
				Message: "can't upload image: 124",
				Field:   "photo_profile",
				Error:   e.Error(),
			})
			return
		}

		if err.Error() == gorm.ErrRecordNotFound.Error() {
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

	c.JSON(http.StatusOK, gin.H{
		"status":        "success",
		"photo_profile": data,
	})
}
