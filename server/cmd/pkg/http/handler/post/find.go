package PostHandler

import (
	"errors"
	"net/http"

	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary get post
// @Description get post by id
// @Tags post
// @Produce  json
// @Param id path string true "id post"
// @Success 200 {object} dto.PostResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /post/{id} [get]
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
		c.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Status:  "failed",
			Message: "can't get post",
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.PostResponse{
		Status: "success",
		Data: dto.Post{
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
