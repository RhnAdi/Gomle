package UserHandler

import (
	"net/http"

	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) Follow(c *gin.Context) {
	id := c.Param("id")
	claim := c.MustGet("claim").(auth.JWTClaim)

	data, err := h.service.Follow(claim, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Status:  "failed",
			Message: "can't follow user",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": dto.FollowedResponse{
			User: data.ID,
		},
	})
}
