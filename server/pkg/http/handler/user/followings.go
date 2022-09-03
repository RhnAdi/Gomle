package UserHandler

import (
	"net/http"

	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) Followings(c *gin.Context) {
	claim := c.MustGet("claim").(auth.JWTClaim)

	data, err := h.service.Followings(claim)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Status:  "failed",
			Message: "can't get followings",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": dto.FollowingResponse{
			Count:      len(data),
			Followings: data,
		},
	})
}
