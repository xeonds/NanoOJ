package lib

import "time"

type UserClaim struct {
	Permission int       `json:"permission"`
	Name       string    `json:"name"`
	Expire     time.Time `json:"expire"` // token过期时间
}

// 常量定义
type serverConfig struct {
	MailUserName   string `json:"mail_username"`
	MailServer     string `json:"mail_server"`
	MailServerPort string `json:"mail_server_port"`
	MailPassword   string `json:"mail_password"`
	CaptchaLength  int    `json:"captcha_length"`
	CaptchaAlive   int    `json:"captcha_alive"`
}

// 应用配置数据
type Config struct {
	serverConfig
}

// 分页数据
type Pagination struct {
	PageSize int
	PageNum  int
}

var jwtSecret = "secret"
