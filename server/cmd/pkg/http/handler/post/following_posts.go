package PostHandler

import (
	"net/http"

	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Tags post
// @Summary get all followings post in dashboard
// @Description get all followings post need auth token in authorization
// @Accept json
// @Produce json
// @Success 200 {object} dto.AllMyPostResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /post/dashboard [GET]
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
	c.JSON(http.StatusOK, dto.AllPostResponse{
		Status: "success",
		Data:   data,
	})
}
