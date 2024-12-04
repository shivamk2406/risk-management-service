package transport

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shivamk2406/risk-management-service/config"
)


func NewRestServer(config config.AppConfig, routersInit *gin.Engine) *http.Server{
	gin.SetMode(config.Web.ServerMode)
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