package config

// JwtConfig
// @Description: jwt 认证相关配置
type JwtConfig struct {
	SigningKey     string `mapstructure:"signing_key"`     // 用于签名和验证 JWT 的秘钥
	ExpirationTime string `mapstructure:"expiration_time"` // 令牌过期时间
	BufferTime     string `mapstructure:"buffer_time"`     // 令牌缓冲时间
	Issuer         string `mapstructure:"issuer"`          // 令牌发行人标识
}
