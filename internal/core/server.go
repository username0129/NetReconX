package core

import (
	"fmt"
	"net/http"
	"server/internal/config"
	"time"
)

func StartServer() {
	router := InitializeRout()

	address := fmt.Sprintf("%v:%v", config.GlobalConfig.SystemConfig.Ip, config.GlobalConfig.SystemConfig.Port)
	server := &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()

	if err != nil {
		return
	}
}
