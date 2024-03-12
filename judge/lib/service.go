package lib

import (
	"context"
	"crypto/tls"
	"errors"
	"net/smtp"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jordan-wright/email"
)

// 验证码服务，使用redis存储
func MailSendCaptcha(mailFrom, mailServer, mailTo string, rdb *redis.Client) error {
	config, err := LoadConfig()
	if err != nil {
		return errors.New("[MailSend] config file read failed")
	}
	codeId, code := GenerateCaptcha(config.CaptchaLength)
	rdb.Set(context.Background(), codeId, code, time.Minute*time.Duration(config.CaptchaAlive))
	e := email.NewEmail()
	e.From = "Get <" + mailFrom + ">"
	e.To = []string{mailTo}
	e.Subject = "验证码"
	e.HTML = []byte("你的验证码为：<h1>" + code + "</h1>")
	err = e.SendWithTLS(mailServer+"465", smtp.PlainAuth("", config.MailUserName, config.MailPassword, mailServer),
		&tls.Config{InsecureSkipVerify: true, ServerName: mailServer})
	if err != nil {
		return err
	}
	return nil
}
