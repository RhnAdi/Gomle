package PostHandler

import (
	"net/http"
	"path/filepath"

	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/RhnAdi/Gomle/pkg/domain"
	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/RhnAdi/Gomle/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *PostHandler) Update(c *gin.Context) {
	var body dto.PostRequestBody
	formdata, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	// Bind Form Data to body
	if val, ok := formdata.Value["content"]; ok && val[0] != "" {
		body.Content = val[0]
	} else {
		c.JSON(http.StatusBadRequest, helper.ErrorResponse{
			Status:  "failed",
			Message: "content not empty",
			Field:   "content",
		})
	}

	// Handle Multiple Upload
	for _, file := range formdata.File["files"] {
		extension := filepath.Ext(file.Filename)
		filetype := file.Header.Get("Content-Type")

		if !(filetype == "image/jpeg" || filetype == "image/png" || filetype == "image/webp") {
			c.JSON(http.StatusBadRequest, helper.ErrorResponse{
				Status:  "failed",
				Message: "file type not allowed",
				Field:   "files",
				Error:   "only image/jpeg or image/png or image/webp",
			})
			c.Abort()
			return
		}

		// Generate Name
		newFilename := uuid.New().String() + extension

		// Upload File
		err := c.SaveUploadedFile(file, "../../public/images/"+newFilename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "failed",
				"message": "can't upload file",
			})
			c.Abort()
			return
		}

		// Save to Body
		body.Files = append(body.Files, models.Image{
			Filename: newFilename,
		})
	}

	id := c.Param("id")
	claim := c.MustGet("claim").(auth.JWTClaim)
	data, err := h.service.Update(claim, domain.Post{
		ID:      id,
		Content: body.Content,
		Files:   body.Files,
	})
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			c.JSON(http.StatusInternalServerError, helper.ErrorResponse{
				Status:  "failed",
				Message: "post not found",
				Field:   "post_id",
				Error:   err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Status:  "failed",
			Message: "can't delete post",
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
		"data": dto.PostResponse{
			ID:        data.ID,
			UserID:    data.UserID,
			Content:   data.Content,
			Files:     data.Files,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		},
	})
}
