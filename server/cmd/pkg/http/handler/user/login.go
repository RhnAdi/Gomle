package UserHandler

import (
	"net/http"

	"github.com/RhnAdi/Gomle/server/pkg/dto"
	"github.com/RhnAdi/Gomle/server/pkg/http/helper"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// @Summary account login
// @Description login to account
// @Tags authentication
// @Accept json
// @Produce json
// @Param user body dto.UserLoginBody true "login user"
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} helper.ErrorResponse
// @Router /users/login [POST]
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
