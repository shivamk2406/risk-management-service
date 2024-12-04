// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/shivamk2406/risk-management-service/cmd/transport/routers"
	"github.com/shivamk2406/risk-management-service/cmd/transport/routers/api/v1"
	"github.com/shivamk2406/risk-management-service/config"
	"github.com/shivamk2406/risk-management-service/repo"
	"github.com/shivamk2406/risk-management-service/service/risk"
)

// Injectors from wire.go:

func initializeConfig() config.AppConfig {
	appConfig := config.LoadAppConfig()
	return appConfig
}

func intializeService() *gin.Engine {
	api := repo.NewRiskRepo()
	riskService := service.NewRiskSvc(api)
	riskController := v1.NewRiskController(riskService)
	engine := routers.InitRouter(riskController)
	return engine
}