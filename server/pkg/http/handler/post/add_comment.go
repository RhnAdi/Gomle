package PostHandler

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *PostHandler) AddComment(c *gin.Context) {
	formdata, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	claim := c.MustGet("claim").(auth.JWTClaim)
	id := c.Param("id")

	// Handle Upload
	file := formdata.File["file"][0]
	filetype := file.Header.Get("Content-Type")
	if !(filetype == "image/jpeg" || filetype == "image/png" || filetype == "image/webp") {
		c.JSON(http.StatusBadRequest, helper.ErrorResponse{
			Status:  "failed",
			Message: "file type not allowed",
			Field:   "file",
			Error:   "only image/jpeg or image/png or image/webp",
		})
		return
	}

	extension := filepath.Ext(file.Filename)
	newFilename := uuid.New().String() + extension

	// Saving Image
	err = c.SaveUploadedFile(file, "../../public/images/"+newFilename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Status:  "failed",
			Message: "can't upload file",
			Field:   "file",
			Error:   err.Error(),
		})
		return
	}

	// Handle Service
	data, err := h.service.AddComment(claim, id, dto.CommentRequest{
		Text: formdata.Value["text"][0],
		File: newFilename,
	})

	if err != nil {
		// Delete saving image
		if e := os.Remove("../../public/images/" + newFilename); e != nil {
			c.JSON(http.StatusNotFound, helper.ErrorResponse{
				Status:  "failed",
				Message: "can't upload image: 124",
				Field:   "photo_profile",
				Error:   e.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Status:  "failed",
			Message: "can't add comment",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})

}
