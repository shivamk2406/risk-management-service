package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/shivamk2406/risk-management-service/cmd/transport/routers/api/v1"
)

func InitRouter(riskController v1.RiskController) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())


	apiv1 := r.Group("/api/v1")
	{
		
		apiv1.GET("/risks",riskController.GetRisks )
		
		apiv1.POST("/risks",)

		apiv1.GET("/risks/:id", riskController.GetRisksById)
	}

	return r
}