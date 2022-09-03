package UserHandler

import (
	"net/http"

	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (h *UserHandler) Login(c *gin.Context) {
	var user dto.UserLoginBody
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorAuthResponse{
			Status:  "failed",
			Message: err.Error(),
			Field:   "",
		})
		return
	}
	token, err := h.service.Login(user)

	if err != nil && err.Error() == gorm.ErrRecordNotFound.Error() {
		c.JSON(http.StatusForbidden, dto.ErrorAuthResponse{
			Status:  "failed",
			Message: "email not found",
			Field:   "email",
		})
		return
	}
	if err != nil && err.Error() == bcrypt.ErrMismatchedHashAndPassword.Error() {
		c.JSON(http.StatusForbidden, dto.ErrorAuthResponse{
			Status:  "failed",
			Message: "wrong password",
			Field:   "password",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Status:  "failed",
			Message: "can't login",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.AuthResponse{
		Status: "success",
		Token:  token,
	})
}
