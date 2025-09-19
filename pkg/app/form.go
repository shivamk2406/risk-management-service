package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errs "github.com/shivamk2406/risk-management-service/pkg/err"
)

func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, errs.INVALID_PARAMS
	}
	return http.StatusOK, errs.SUCCESS
}