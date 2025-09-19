package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/shivamk2406/risk-management-service/cmd/transport/routers/api/v1"
	docs "github.com/shivamk2406/risk-management-service/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Risk Management APIs
// @version 1.0
// @description Testing Swagger APIs.
// @termsOfService http://swagger.io/terms/
func InitRouter(riskController v1.RiskController) *gin.Engine {
	r := gin.New()
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))


	apiv1 := r.Group("/api/v1")
	{
		
		apiv1.GET("/risks",riskController.GetRisks )
		
		apiv1.POST("/risks",riskController.AddRisk)

		apiv1.GET("/risks/:id", riskController.GetRisksById)
	}
	return r
}