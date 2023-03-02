package UserHandler

import (
	"net/http"

	"github.com/RhnAdi/Gomle/server/internal/auth"
	"github.com/RhnAdi/Gomle/server/pkg/dto"
	"github.com/RhnAdi/Gomle/server/pkg/http/helper"
	"github.com/gin-gonic/gin"
)

// @Tags accounts
// @Summary get list your followings user
// @Description your list followings user
// @Accept json
// @Produce json
// @securitydefinitions.bearer
// @Success 200 {object} dto.FollowingsResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /users/followings/ [GET]
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

	c.JSON(http.StatusOK, dto.FollowingsResponse{
		Status: "success",
		Data: dto.Following{
			Count:      len(data),
			Followings: data,
		},
	})
}
