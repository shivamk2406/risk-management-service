package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/shivamk2406/risk-management-service/interfaces"
)

type RiskController struct{
	RiskSvc interfaces.RiskService
}

func NewRiskController(riskSvc interfaces.RiskService) RiskController{
	return RiskController{
		RiskSvc: riskSvc,
	}
}

func(r *RiskController) GetRisks(c *gin.Context){
	appG := app.Gin{C: c}
	id := c.Param("id")

}