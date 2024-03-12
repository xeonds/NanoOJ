package config

import "xyz.xeonds/nano-oj/lib"

type Config struct {
	lib.ServerConfig
	lib.DatabaseConfig `json:"-"`
	lib.RedisConfig
	lib.CaptchaConfig
	lib.MailConfig
	Judger struct {
		Daemons []string
	}
	ServerType string `json:"server_type"`
}
