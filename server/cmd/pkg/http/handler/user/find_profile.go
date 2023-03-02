package UserHandler

import (
	"net/http"

	"github.com/RhnAdi/Gomle/server/pkg/dto"
	"github.com/RhnAdi/Gomle/server/pkg/http/helper"
	"github.com/gin-gonic/gin"
)

// @Summary account profile
// @Description get account profile
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path string true "id account"
// @Success 200 {object} dto.UserProfile
// @Failure 403 {object} helper.ErrorResponse
// @Router /users/{id}/profile [GET]
func (h *UserHandler) FindProfile(c *gin.Context) {
	id := c.Param("id")
	data, err := h.service.FindProfile(id)
	if err != nil {
		c.JSON(http.StatusNotFound, helper.ErrorResponse{
			Field:   "id",
			Status:  "failed",
			Message: "profile not found",
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.UserProfile{
		Status: "success",
		Data:   data,
	})
}
