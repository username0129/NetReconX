package global

import (
	"github.com/allegro/bigcache/v3"
	"github.com/mojocn/base64Captcha"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"server/internal/config"
)

// 存储全局变量如后端配置、数据库链接等

var (
	Version      = "v0.1"                        // 后端版本
	CaptchaStore = base64Captcha.DefaultMemStore // 验证码存储
	Config       config.Server                   // 后端配置
	Viper        *viper.Viper                    // 全局 viper
	DB           *gorm.DB                        // 全局 viper
	Logger       *zap.Logger                     // zap 日志
	Cache        *bigcache.BigCache              // bigcache 缓存
)
