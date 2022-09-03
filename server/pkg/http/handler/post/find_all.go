package PostHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *PostHandler) FindAll(c *gin.Context) {
	data, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}
