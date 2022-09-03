package PostHandler

import (
	"net/http"

	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/RhnAdi/Gomle/pkg/domain"
	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *PostHandler) Update(c *gin.Context) {
	var (
		id   string
		body dto.PostRequestBody
	)
	id = c.Param("id")
	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
	}

	claim := c.MustGet("claim").(auth.JWTClaim)
	data, err := h.service.Update(claim, domain.Post{
		ID:      id,
		Content: body.Content,
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
