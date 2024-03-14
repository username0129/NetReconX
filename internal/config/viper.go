package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Viper
//
//	@Description: 使用 Viper 管理后端配置文件，默认配置文件为 ./config/config.yaml。
//	@param path
//	@return *viper.Viper
func Viper() *viper.Viper {
	fmt.Printf("当前使用的配置文件为：%v\n", GlobalConfig.ConfigPath)

	v := viper.New()
	v.SetConfigFile(GlobalConfig.ConfigPath) // 配置文件路径

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取配置文件错误: %v\n", err))
	} // 读取配置文件

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置文件发生更改：%v\n", in.Name)
		if err := v.Unmarshal(&GlobalConfig); err != nil {
			panic(fmt.Errorf("新配置文件解析错误: %v\n", err))
		}
	}) // 监视配置文件更改

	if err := v.Unmarshal(&GlobalConfig); err != nil {
		panic(fmt.Errorf("配置文件解析错误: %v\n", err))
	} // 解析配置文件到全局配置

	return v
}
