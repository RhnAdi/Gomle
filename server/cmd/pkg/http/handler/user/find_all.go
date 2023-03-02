package UserHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) FindAll(c *gin.Context) {
	users, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": users})
}
