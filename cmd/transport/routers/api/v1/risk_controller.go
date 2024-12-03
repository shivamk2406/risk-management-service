package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/shivamk2406/risk-management-service/interfaces"
	"github.com/shivamk2406/risk-management-service/pkg/app"
	errs "github.com/shivamk2406/risk-management-service/pkg/err"
)

type RiskController struct {
	RiskSvc interfaces.RiskService
}

func NewRiskController(riskSvc interfaces.RiskService) RiskController {
	return RiskController{
		RiskSvc: riskSvc,
	}
}

func (r *RiskController) GetRisksById(c *gin.Context) {
	appG := app.Gin{C: c}
	id := c.Param("id")
	err := validation.Validate(id,
		validation.Required,
		validation.Length(5, 100),
		is.UUID,
	)

	if err != nil {
		appG.Response(http.StatusBadRequest, errs.INVALID_PARAMS, nil)
		return
	}

	risks, err := r.RiskSvc.GetRiskById(id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, errs.ERROR, nil)
	}

	appG.Response(http.StatusOK, errs.SUCCESS, risks)

}
