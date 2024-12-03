package transport

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shivamk2406/risk-management-service/cmd/transport/routers"
	v1 "github.com/shivamk2406/risk-management-service/cmd/transport/routers/api/v1"
	"github.com/shivamk2406/risk-management-service/config"
)


func NewRestServer(config config.Appconfig, riskController v1.RiskController) *http.Server{
	gin.SetMode(gin.DebugMode)
	routersInit := routers.InitRouter(riskController)
	readTimeout := config.Web.ReadTimeout 
	writeTimeout := config.Web.WriteTimeout 
	endPoint := fmt.Sprintf(":%s", config.Web.APIHost)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	return server

}