package PostHandler

import (
	"net/http"

	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/gin-gonic/gin"
)

// @Summary list post
// @Description list all post
// @Tags post
// @Produce  json
// @Success 200 {object} dto.AllPostResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /post/ [get]
func (h *PostHandler) FindAll(c *gin.Context) {
	data, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Status:  "failed",
			Message: "get all post failed",
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.AllPostResponse{
		Status: "success",
		Data:   data,
	})
}
