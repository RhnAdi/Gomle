package UserHandler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/http/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgconn"
)

func (h *UserHandler) Register(c *gin.Context) {
	var user dto.UserRegisterBody
	err := c.ShouldBind(&user)
	var fieldError validator.ValidationErrors
	if err != nil && errors.As(err, &fieldError) {
		var fe []string
		for _, e := range fieldError {
			fe = append(fe, e.Field())
		}
		c.JSON(http.StatusBadRequest, helper.ErrorResponse{
			Status:  "failed",
			Message: "reuqired fields",
			Field:   strings.Join(fe, ","),
			Error:   fieldError.Error(),
		})
		return
	}
	token, err := h.service.Register(user)
	var perr *pgconn.PgError
	if err != nil && errors.As(err, &perr) && perr.Code == "23505" {
		spl_err_col := strings.Split(perr.ConstraintName, "_")

		c.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Status:  "failed",
			Message: spl_err_col[len(spl_err_col)-2] + " already exist",
			Error:   perr.Message,
			Field:   spl_err_col[len(spl_err_col)-2],
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Status:  "failed",
			Message: "can't register",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.AuthResponse{
		Status: "success",
		Token:  token,
	})
}
