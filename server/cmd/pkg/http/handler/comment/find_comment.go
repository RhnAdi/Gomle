package CommentHandler

import (
	"errors"
	"net/http"

	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *CommentHandler) FindComment(c *gin.Context) {
	id := c.Param("id")

	data, err := h.service.FindComment(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, helper.ErrorResponse{
				Status:  "failed",
				Field:   "id",
				Message: "comment not found",
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Status:  "failed",
			Message: "can't get comment",
			Error:   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, dto.CommentResponse{
		Status: "success",
		Data:   data,
	})

	return
}
