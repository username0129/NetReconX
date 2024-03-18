package core

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"server/internal/config"
	"server/internal/util"
)

func GetZapLogger() *zap.Logger {
	if ok, _ := util.IsPathExist(config.GlobalConfig.ZapConfig.Director); !ok {
		_ = os.Mkdir(config.GlobalConfig.ZapConfig.Director, os.ModePerm)
	}
	cores := getZapCores()
	logger := zap.New(zapcore.NewTee(cores...))
	logger = logger.WithOptions(zap.AddCaller())
	return logger
}

func getZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := config.GlobalConfig.ZapConfig.GetZapLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, getEncoderCore(level, getLevelPriority(level)))
	}
	return cores
}

// getEncoder
//
//	@Description: 获取 zapcore.Encoder
//	@return logger
func getEncoder() zapcore.Encoder {
	if config.GlobalConfig.ZapConfig.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderConfig
//
//	@Description: 获取 zapcore.EncoderConfig
//	@receiver z
//	@return zapcore.EncoderConfig
func getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    config.GlobalConfig.ZapConfig.GetZapEncodeLevel(),
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func getEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer := getWriteSyncer(l.String())
	return zapcore.NewCore(getEncoder(), writer, level)
}

// getLevelPriority
//
//	@Description: 根据 zapcore.Level 获取 zap.LevelEnablerFunc
//	@receiver z
//	@param level
//	@return zap.LevelEnablerFunc
func getLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool {
			return level == zap.DebugLevel
		} // 调试级别
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日志级别
			return level == zap.InfoLevel
		} // 日志级别
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告级别
			return level == zap.WarnLevel
		} // 告警级别
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 错误级别
			return level == zap.ErrorLevel
		} // 错误级别
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic级别
			return level == zap.DPanicLevel
		} // dpanic 级别
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic级别
			return level == zap.PanicLevel
		} // panic 级别
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // 终止级别
			return level == zap.FatalLevel
		} // 终止级别
	default:
		return func(level zapcore.Level) bool {
			return level == zap.DebugLevel
		} // 调试级别
	}
}

func getWriteSyncer(level string) zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")

	//fileWriter := NewCutter(global.GVA_CONFIG.Zap.Director, level, WithCutterFormat("2006-01-02"))
	if config.GlobalConfig.ZapConfig.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(file))
	}
	return zapcore.AddSync(file)
}
