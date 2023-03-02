package UserHandler

import (
	"errors"
	"net/http"

	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Tags accounts
// @Summary follow user to get her post
// @Description  follow people for get update their post and save them in your contact. need auth token in header to call this endpoint
// @Accept json
// @Produce json
// @securitydefinitions.bearer
// @Success 200 {object} dto.FollowedResponse
// @Failure 404 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /users/follow/{id} [GET]
func (h *UserHandler) Follow(c *gin.Context) {
	id := c.Param("id")
	claim := c.MustGet("claim").(auth.JWTClaim)

	data, err := h.service.Follow(claim, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, helper.ErrorResponse{
				Status:  "failed",
				Message: "user not found",
				Error:   err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Status:  "failed",
			Message: "can't follow user",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.FollowedResponse{
		Status:   "success",
		Followed: data.ID,
	})
}
