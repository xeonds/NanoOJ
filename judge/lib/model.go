package lib

import "time"

type UserClaim struct {
	ID         int       `json:"id"`
	Permission int       `json:"permission"`
	Name       string    `json:"name"`
	Expire     time.Time `json:"expire"` // token过期时间
}

// 通用配置
type ServerConfig struct {
	Port string `json:"port"`
}

type DatabaseConfig struct {
	Type     string `json:"type"` // 数据库类型
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"db"` // 数据库名
	Migrate  bool   `json:"migrate"`
}

type MailConfig struct {
	MailUserName   string `json:"mail_username"`
	MailServer     string `json:"mail_server"`
	MailServerPort string `json:"mail_server_port"`
	MailPassword   string `json:"mail_password"`
}

type CaptchaConfig struct {
	CaptchaLength int `json:"captcha_length"`
	CaptchaAlive  int `json:"captcha_alive"`
}

type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
}

// 分页数据
type Pagination struct {
	PageSize int
	PageNum  int
}

var jwtSecret = "secret"
