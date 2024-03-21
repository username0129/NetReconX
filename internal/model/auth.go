package model

type LoginRequest struct {
	Username  string `json:"username,omitempty"`   // 用户名
	Password  string `json:"password,omitempty"`   // 登陆密码
	Answer    string `json:"answer,omitempty"`     // 验证码
	CaptchaId string `json:"captcha_id,omitempty"` // 验证码 ID
}

type LoginUser struct {
	Username string `json:"username,omitempty"` // 用户名
	Password string `json:"password,omitempty"` // 登陆密码
}
