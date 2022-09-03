package UserHandler

import (
	"net/http"

	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) FindProfile(c *gin.Context) {
	claim := c.MustGet("claim").(auth.JWTClaim)
	data, err := h.service.FindProfile(claim.ID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "failed",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}
