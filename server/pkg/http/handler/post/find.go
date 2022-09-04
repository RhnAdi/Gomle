package PostHandler

import (
	"errors"
	"net/http"

	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *PostHandler) Find(c *gin.Context) {
	id := c.Param("id")
	data, err := h.service.Find(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, helper.ErrorResponse{
				Status:  "failed",
				Message: "post not found",
				Field:   "post_id",
				Error:   err.Error(),
			})
			return
		}
		c.JSON(http.StatusForbidden, helper.ErrorResponse{
			Status:  "failed",
			Message: "can't get post",
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": dto.PostResponse{
			ID:        data.ID,
			UserID:    data.UserID,
			Content:   data.Content,
			Files:     data.Files,
			Comments:  data.Comments,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		},
	})
}
