package config

// ServerConfig
// @Description: 后端配置
type ServerConfig struct {
	Ip         string    `mapstructure:"ip"`          // 后端 IP 地址
	Port       string    `mapstructure:"port"`        // 后端端口
	ConfigPath string    `mapstructure:"config_path"` // 配置文件路径
	Jwt        JwtConfig `mapstructure:"jwt"`         // jwt 配置
}
