package UserHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) FindProfile(c *gin.Context) {
	id := c.Param("id")
	data, err := h.service.FindProfile(id)
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
