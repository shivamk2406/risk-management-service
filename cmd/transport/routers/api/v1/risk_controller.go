package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/shivamk2406/risk-management-service/interfaces"
	"github.com/shivamk2406/risk-management-service/models"
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

// @Summary Get a single Risk
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /risks/{id} [get]
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


// @Summary Get multiple risks
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /risks [get]
func (r *RiskController) GetRisks(c *gin.Context) {
	appG := app.Gin{C: c}

	risks, err := r.RiskSvc.GetRisks()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errs.ERROR, nil)
	}
	appG.Response(http.StatusOK, errs.SUCCESS, risks)

}

// @Summary Add risks
// @Produce  json
// @Accept json
// @Param risk body models.RiskRequestDto true "risk"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /risks [post]
func (r *RiskController) AddRisk(c *gin.Context) {
	appG := app.Gin{C: c}

	var risk models.RiskRequestDto

	if err := c.Bind(&risk); err != nil {
		fmt.Println(err)
		appG.Response(http.StatusBadRequest, errs.INVALID_PARAMS, nil)
		return
	}

	resp, err := r.RiskSvc.CreateRisk(&risk)
	if err != nil {
		appG.Response(http.StatusInternalServerError, errs.ERROR, nil)
		return
	}

	appG.Response(http.StatusCreated, errs.SUCCESS, resp)

}
