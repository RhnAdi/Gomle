package UserHandler

import (
	"net/http"

	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/gin-gonic/gin"
)

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

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": dto.FollowersResponse{
			Count:     len(data),
			Followers: data,
		},
	})
}
