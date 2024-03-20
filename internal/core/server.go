package core

import (
	"fmt"
	"net/http"
	"server/internal/global"
	"time"
)

func StartServer() {
	router := InitializeRout()

	address := fmt.Sprintf("%v:%v", global.Config.System.Ip, global.Config.System.Port)
	server := &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	global.Logger.Info("服务端监听在 " + address)
	global.Logger.Error(server.ListenAndServe().Error())
}
