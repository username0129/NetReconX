package config

import (
	"go.uber.org/zap/zapcore"
)

// ZapConfig
// @Description: zap 日志配置
type ZapConfig struct {
	Level         string `mapstructure:"level"`          // 最低记录日志等级
	Format        string `mapstructure:"format"`         // 输出格式
	Prefix        string `mapstructure:"prefix"`         // 记录前缀
	Director      string `mapstructure:"director"`       // 日志保存路径
	EncodeLevel   string `mapstructure:"encode_level"`   // 日志级别的编码器
	StacktraceKey string `mapstructure:"stacktrace_key"` // 记录堆栈信息
	LogInConsole  bool   `mapstructure:"log_in_console"` // 日志输出到控制台
}

// GetZapLevel
//
//	@Description: 根据 Level 获取 zapcore.Level 日志级别
//	@receiver zc
//	@return zapcore.LevelEncoder
func (zc *ZapConfig) GetZapLevel() zapcore.Level {

	switch zc.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

// GetZapEncodeLevel
//
//	@Description: 根据 EncodeLevel 获取 zapcore.LevelEncoder
//	@receiver zc
//	@return zapcore.LevelEncoder
func (zc *ZapConfig) GetZapEncodeLevel() zapcore.LevelEncoder {
	switch zc.EncodeLevel {
	case "LowercaseLevelEncoder":
		return zapcore.LowercaseLevelEncoder // 小写编码器
	case "LowercaseColorLevelEncoder":
		return zapcore.LowercaseColorLevelEncoder // 小写编码器带颜色
	case "CapitalLevelEncoder":
		return zapcore.CapitalLevelEncoder // 大写编码器
	case "CapitalColorLevelEncoder":
		return zapcore.CapitalColorLevelEncoder // 大写编码器带颜色
	default:
		return zapcore.LowercaseLevelEncoder // 默认小写编码器
	}
}
