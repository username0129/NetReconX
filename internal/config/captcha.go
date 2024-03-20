package config

import "time"

type Captcha struct {
	Long               int           `mapstructure:"long" json:"long,omitempty" yaml:"long"`                                                 // 验证码字符长度
	ImgWidth           int           `mapstructure:"img_width" json:"img_width,omitempty" yaml:"img_width"`                                  // 图片宽度
	ImgHeight          int           `mapstructure:"img_height" json:"img_height,omitempty" yaml:"img_height"`                               // 图片高度
	OpenCaptcha        int           `mapstructure:"open_captcha" json:"open_captcha,omitempty" yaml:"open_captcha"`                         // 0 -> 开启验证码，其他数字 -> 密码错误超过指定次数开启验证码
	OpenCaptchaTimeout time.Duration `mapstructure:"open_captcha_timeout" json:"open_captcha_timeout,omitempty" yaml:"open_captcha_timeout"` // 验证码超时时间
}
