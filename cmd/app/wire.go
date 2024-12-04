//go:build wireinject
// +build wireinject

//go:generate go run github.com/google/wire/cmd/wire

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/shivamk2406/risk-management-service/cmd/transport/routers"
	v1 "github.com/shivamk2406/risk-management-service/cmd/transport/routers/api/v1"
	"github.com/shivamk2406/risk-management-service/config"
	"github.com/shivamk2406/risk-management-service/internal/repo"
	service "github.com/shivamk2406/risk-management-service/service/risk"
)


func initializeConfig() config.AppConfig {
	wire.Build(config.LoadAppConfig)
	return config.AppConfig{}
}

func intializeService() *gin.Engine {
	wire.Build(repo.NewRiskRepo, service.NewRiskSvc, v1.NewRiskController, routers.InitRouter)
	return &gin.Engine{}
}