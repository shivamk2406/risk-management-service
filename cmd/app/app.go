package app

import (
	"fmt"
	"net/http"

	"github.com/shivamk2406/risk-management-service/cmd/transport"
)

func Start(){
    config:=initializeConfig()
    engine:=intializeService()
    server:=transport.NewRestServer(config,engine)

    fmt.Printf("Starting server on %s...\n", config.Web.APIHost)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Error starting server: %v\n", err)
	}
}