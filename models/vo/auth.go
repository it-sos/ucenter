package vo

type AuthVO struct {
	Account  string `json:"account" binding:"required"`  // 帐号
	Password string `json:"password" binding:"required"` // 密码
	Captcha  string `json:"captcha"`                     // 验证码（第2次认证必填）
}

type AuthRefreshTokenVO struct {
	RefreshToken string `json:"refresh_token" binding:"required"` // 刷新token
}

type AuthTokenVO struct {
	RefreshToken string `json:"refresh_token" binding:"required"` // 刷新token
	Token        string `json:"token" binding:"required"`         // token
}
