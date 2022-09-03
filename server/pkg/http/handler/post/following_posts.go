package PostHandler

import (
	"net/http"

	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *PostHandler) FollowingPosts(c *gin.Context) {
	claim := c.MustGet("claim").(auth.JWTClaim)
	data, err := h.service.FollowingPosts(claim)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			c.JSON(http.StatusInternalServerError, helper.ErrorResponse{
				Status:  "failed",
				Message: "post not found",
				Field:   "post_id",
				Error:   err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Status:  "failed",
			Message: "can't delete post",
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}
