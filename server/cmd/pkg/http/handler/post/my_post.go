package PostHandler

import (
	"net/http"

	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/gin-gonic/gin"
)

// @Tags post
// @Summary get all my post
// @Description get all my post need auth token in authorization
// @Accept json
// @Produce json
// @Success 200 {object} dto.AllMyPostResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /post/mypost [GET]
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
	c.JSON(http.StatusOK, dto.AllMyPostResponse{
		Status: "success",
		Data:   data,
	})
}
