package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/shivamk2406/risk-management-service/enums"
	"github.com/shivamk2406/risk-management-service/interfaces"
	"github.com/shivamk2406/risk-management-service/internal/models"
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

// @Summary Get risk by id
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /risks/{id} [get]
func (r *RiskController) GetRisksById(c *gin.Context) {
	logger := ctxzap.Extract(c).Sugar()

	appG := app.Gin{C: c}
	id := c.Param("id")
	err := validation.Validate(id,
		validation.Required,
		validation.Length(5, 100),
		is.UUID,
	)

	if err != nil {
		logger.Errorf("error in the request %s",err.Error())
		appG.Response(http.StatusBadRequest, errs.INVALID_PARAMS, nil)
		return
	}

	risks, err := r.RiskSvc.GetRiskById(id)
	if err != nil {
		logger.Errorf("error while fetching risk by id for id: %s error: %s",id, err.Error())
		appG.Response(http.StatusInternalServerError, errs.ERROR, nil)
	}

	appG.Response(http.StatusOK, errs.SUCCESS, risks)

}

// @Summary Get list of available risks
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /risks [get]
func (r *RiskController) GetRisks(c *gin.Context) {
	logger := ctxzap.Extract(c).Sugar()

	appG := app.Gin{C: c}

	risks, err := r.RiskSvc.GetRisks()
	if err != nil {
		logger.Errorf("error while fetching risk %s",err.Error())
		appG.Response(http.StatusInternalServerError, errs.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, errs.SUCCESS, risks)

}

// @Summary Add risk
// @Produce  json
// @Accept json
// @Param risk body models.RiskRequestDto true "risk request model"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /risks [post]
func (r *RiskController) AddRisk(c *gin.Context) {
	logger := ctxzap.Extract(c).Sugar()
	appG := app.Gin{C: c}

	var risk models.RiskRequestDto

	if err := c.Bind(&risk); err != nil {
		logger.Errorf("error while binding risk to struct %s",err.Error())
		appG.Response(http.StatusBadRequest, errs.INVALID_PARAMS, nil)
		return
	}

	err := validation.ValidateStruct(&risk,
		validation.Field(&risk.State, validation.Required, validation.In(enums.RiskStateStrings())),
	)
	if err != nil {
		logger.Errorf("error in the risk provided as input %s",err.Error())
		appG.Response(http.StatusBadRequest, errs.ERROR, nil)
		return
	}

	resp, err := r.RiskSvc.CreateRisk(&risk)
	if err != nil {
		logger.Errorf("error while creating risk %s",err.Error())
		appG.Response(http.StatusInternalServerError, errs.ERROR, nil)
		return
	}

	appG.Response(http.StatusCreated, errs.SUCCESS, resp)

}
