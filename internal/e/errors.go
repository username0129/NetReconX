package e

import "errors"

var (
	// ErrInvalidDBConfig 数据库配置文件无效
	ErrInvalidDBConfig = errors.New("数据库配置文件无效")

	// ErrInvalidInitializers 无可用初始化过程
	ErrInvalidInitializers = errors.New("无可用初始化过程")
)
