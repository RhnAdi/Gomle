package UserHandler

import (
	"net/http"

	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/gin-gonic/gin"
)

// @Tags accounts
// @Summary get list your followers user
// @Description your list followers user
// @Accept json
// @Produce json
// @securitydefinitions.bearer
// @Success 200 {object} dto.FollowersResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /users/followers [GET]
func (h *UserHandler) Followers(c *gin.Context) {
	claim := c.MustGet("claim").(auth.JWTClaim)

	data, err := h.service.Followers(claim)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Status:  "failed",
			Message: "can't get followers",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.FollowersResponse{
		Status: "success",
		Data: dto.Followers{
			Count:     len(data),
			Followers: data,
		},
	})
}
