package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shivamk2406/risk-management-service/service"
)

func InitRouter(riskController service.Registry) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())


	apiv1 := r.Group("/api/v1")
	{
		
		apiv1.GET("/risks",riskController.GetRisks )
		
		apiv1.POST("/risks",riskController.AddRisk)

		apiv1.GET("/risks/:id", riskController.GetRisksById)
	}

	return r
}