package PostHandler

import (
	"net/http"

	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/gin-gonic/gin"
)

func (h *PostHandler) FindMyPost(c *gin.Context) {
	claim := c.MustGet("claim").(auth.JWTClaim)
	data, err := h.service.FindMyPost(claim)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
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
