package core

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"server/internal/config"
	"server/internal/util"
)

// customLevelEnabler 仅当等级严格匹配时启用日志记录
type customLevelEnabler zapcore.Level

// Enabled 实现了 zapcore.LevelEnabler 接口
func (cle customLevelEnabler) Enabled(level zapcore.Level) bool {
	return level == zapcore.Level(cle)
}

// GetZapLogger
//
//	@Description: 获取 zap.Logger
//	@return *zap.Logger
//	@Router:
func GetZapLogger() *zap.Logger {
	if ok, _ := util.IsPathExist(config.GlobalConfig.ZapConfig.Director); !ok {
		_ = os.Mkdir(config.GlobalConfig.ZapConfig.Director, 0755)
	} // 创建日志目录

	cores := getZapCores()
	logger := zap.New(zapcore.NewTee(cores...))
	logger = logger.WithOptions(zap.AddCaller())
	return logger
}

// getZapCores
//
//	@Description: 根据日志级别获取 core 进行日志处理
//	@return []zapcore.Core
func getZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := zapcore.DebugLevel; level <= zapcore.FatalLevel; level++ {
		cores = append(cores, getEncoderCore(level))
	}
	return cores
}

// getEncoderCore
//
//	@Description: 获取 zapcore 核心
//	@param l
//	@param level
//	@return zapcore.Core
func getEncoderCore(l zapcore.Level) zapcore.Core {

	writer := getWriteSyncer(l.String())
	return zapcore.NewCore(getEncoder(), writer, customLevelEnabler(l))
}

// getEncoder
//
//	@Description: 设置日志条目格式
//	@return logger
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()              // 标准配置
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder          // 时间显示格式
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder         // 显示完整路径
	encoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder // 小写带颜色

	switch config.GlobalConfig.ZapConfig.Format {
	case "json":
		return zapcore.NewJSONEncoder(encoderConfig)
	case "console":
		fallthrough
	default:
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
}

// getWriteSyncer
//
//	@Description: 自定义 zap 日志消息的输出目的地
//	@param level
//	@return zapcore.WriteSyncer
func getWriteSyncer(level string) zapcore.WriteSyncer {
	fileWriter := NewRotate(level, "2006-01-02", config.GlobalConfig.ZapConfig.Director)
	if config.GlobalConfig.ZapConfig.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	}
	return zapcore.AddSync(fileWriter)
}
