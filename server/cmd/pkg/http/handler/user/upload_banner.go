package UserHandler

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// @Summary upload banner
// @Description need auth token in header for update banner
// @Tags accounts
// @Accept multipart/form-data
// @Produce json
// @Param banner formData file true "upload banner"
// @in formData
// @type file
// @Success 200 {object} dto.UploadFileResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 404 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /users/upload/banner [PUT]
func (h *UserHandler) UploadBanner(c *gin.Context) {
	claim := c.MustGet("claim").(auth.JWTClaim)

	file, _ := c.FormFile("banner")
	extension := filepath.Ext(file.Filename)
	filetype := file.Header.Get("Content-Type")

	if !(filetype == "image/jpeg" || filetype == "image/png" || filetype == "image/webp") {
		c.JSON(http.StatusBadRequest, helper.ErrorResponse{
			Status:  "failed",
			Message: "file type not allowed",
			Field:   "photo_profile",
			Error:   "only image/jpeg or image/png",
		})
		return
	}

	newFilename := uuid.New().String() + extension

	err := c.SaveUploadedFile(file, "public/images/"+newFilename)
	if err != nil {
		if e := os.Remove("public.images/" + newFilename); e != nil {
			c.JSON(http.StatusNotFound, helper.ErrorResponse{
				Status:  "failed",
				Message: "can't upload image: 124",
				Field:   "banner",
				Error:   e.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Status:  "failed",
			Message: "can't upload image",
			Field:   "banner",
			Error:   err.Error(),
		})
		return
	}

	data, err := h.service.UpdateBanner(claim, newFilename)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			c.JSON(http.StatusNotFound, helper.ErrorResponse{
				Status:  "failed",
				Message: "can't upload image",
				Field:   "banner",
				Error:   err.Error(),
			})
			return
		}
		c.JSON(http.StatusNotFound, helper.ErrorResponse{
			Status:  "failed",
			Message: "can't update banner",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.UploadFileResponse{
		Status:   "success",
		Field:    "banner",
		Filename: data,
	})
}
