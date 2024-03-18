package config

type SystemConfig struct {
	Ip           string `mapstructure:"ip"`            // 后端 IP 地址
	Port         string `mapstructure:"port"`          // 后端端口
	DBType       string `mapstructure:"db_type"`       // 数据库类型
	ConfigPath   string `mapstructure:"config_path"`   // 配置文件路径
	RouterPrefix string `mapstructure:"router_prefix"` // api 路由前缀
}
