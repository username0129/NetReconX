package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 存储全局变量如后端配置、数据库链接等

var (
	GlobalVersion = "v0.1"     // 后端版本
	GlobalConfig  ServerConfig // 后端配置
	GlobalViper   *viper.Viper // 全局 viper
	GlobalDB      *gorm.DB     // 全局 viper
	GlobalLogger  *zap.Logger  // zap 日志
)
