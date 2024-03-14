package cmd

import (
	"github.com/spf13/cobra"
	"server/internal/config"
	"server/internal/core"
	"server/internal/database"
)

var (
	configPath string // 配置文件路径

	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start the Gin web server",
		Run: func(cmd *cobra.Command, args []string) {
			start()
		},
	}
)

func init() {
	startCmd.Flags().StringVarP(&config.GlobalConfig.SystemConfig.ConfigPath, "config", "c", "./config/config.yaml", "配置文件路径")
	startCmd.Flags().StringVarP(&config.GlobalConfig.SystemConfig.Ip, "ip", "i", "0.0.0.0", "后端 IP 地址")
	startCmd.Flags().StringVarP(&config.GlobalConfig.SystemConfig.Port, "port", "p", "8080", "后端监听地址")
}

func start() {
	config.GlobalViper = config.Viper()                // 初始化 Viper 用于管理配置文件
	config.GlobalDB = database.GetDatabaseConnection() // 获取数据库连接
	if config.GlobalDB != nil {
		db, _ := config.GlobalDB.DB()
		defer db.Close()
	}
	core.StartServer()
}
