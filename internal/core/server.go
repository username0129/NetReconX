package core

import (
	"fmt"
	"go.uber.org/zap"
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

	config.GlobalLogger.Info("服务端监听在 ", zap.String("地址", address))
	config.GlobalLogger.Error(server.ListenAndServe().Error())
}
