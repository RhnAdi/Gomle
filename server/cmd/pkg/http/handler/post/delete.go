package PostHandler

import (
	"errors"
	"net/http"

	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/RhnAdi/Gomle/pkg/domain"
	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Tags post
// @Summary delete post
// @Description delete post need auth token in authorization
// @Param id path string true "id post"
// @Accept json
// @Produce json
// @Success 200 {object} dto.DeletePostResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 404 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /post/{id} [DELETE]
func (h *PostHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	claim := c.MustGet("claim").(auth.JWTClaim)
	data, err := h.service.Delete(claim, domain.Post{ID: id})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, helper.ErrorResponse{
				Status:  "failed",
				Message: "post not found",
				Field:   "post_id",
				Error:   err.Error(),
			})
			return
		}
		if errors.Is(err, helper.ErrYouAreNotOwner) {
			c.JSON(http.StatusForbidden, helper.ErrorResponse{
				Status:  "failed",
				Message: "can't delete post",
				Error:   err.Error(),
			})
			return
		}
		c.JSON(http.StatusForbidden, helper.ErrorResponse{
			Status:  "failed",
			Message: "can't delete post",
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.DeletePostResponse{
		Status: "success",
		Data:   "post deleted with id = " + data.ID,
	})
}
