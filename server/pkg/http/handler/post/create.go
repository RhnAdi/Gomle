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
)

func (h *PostHandler) Create(c *gin.Context) {
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
	body.Content = formdata.Value["content"][0]

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

	// Saving Data on service
	claim := c.MustGet("claim").(auth.JWTClaim)
	data, err := h.service.Create(claim, domain.Post{
		Content: body.Content,
		Files:   body.Files,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed",
			"error":   err.Error(),
		})
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data": dto.PostResponse{
			ID:        data.ID,
			UserID:    data.UserID,
			Files:     data.Files,
			Content:   data.Content,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		},
	})
}
