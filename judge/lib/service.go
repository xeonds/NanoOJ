package lib

import (
	"context"
	"crypto/tls"
	"embed"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jordan-wright/email"
	"gorm.io/gorm"
)

func AddCRUD[T any](router gin.IRouter, path string, db *gorm.DB) *gin.RouterGroup {
	return APIBuilder(router, func(group *gin.RouterGroup) *gin.RouterGroup {
		group.GET("", getAll[T](db))
		group.GET("/:id", get[T](db))
		group.POST("", create[T](db))
		group.PUT("/:id", update[T](db))
		group.DELETE("/:id", delete[T](db))
		return group
	})(router, path)
}
func AddStatic(router *gin.Engine, staticFileDir []string) {
	for _, dir := range staticFileDir {
		router.NoRoute(gin.WrapH(http.FileServer(http.Dir(dir))))
	}
}
func AddStaticFS(router *gin.Engine, fs embed.FS) {
	router.NoRoute(gin.WrapH(http.FileServer(http.FS(fs))))
}
func AddFindAPI[T any](router gin.IRouter, path string, mode string, db *gorm.DB) *gin.RouterGroup {
	return APIBuilder(router, func(group *gin.RouterGroup) *gin.RouterGroup {
		group.POST("", handleFind[T](mode, db))
		return group
	})(router, path)
}

// handlers for gorm
func create[T any](db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var d T
		if err := c.ShouldBindJSON(&d); err != nil {
			c.AbortWithStatus(404)
			log.Println("[gorm]parse creation data failed: ", err)
		} else {
			if err := db.Create(&d).Error; err != nil {
				c.AbortWithStatus(404)
				log.Println("[gorm]create data failed: ", err)
			} else {
				c.JSON(200, d)
			}
		}
	}
}
func get[T any](db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		params := c.Request.URL.Query()
		var d T
		if err := db.Where(params).Find(&d).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, d)
		}
	}
}
func getAll[T any](db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var d []T
		if err := db.Find(&d).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, d)
		}
	}
}
func update[T any](db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var d T
		if err := db.Save(&d).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, d)
		}
	}
}
func delete[T any](db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var d T
		params := c.Request.URL.Query()
		if err := db.Where(params).Delete(&d).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, d)
		}
	}
}
func handleFind[T any](mode string, db *gorm.DB) func(c *gin.Context) {
	if mode == "single" {
		return func(c *gin.Context) {
			var query T
			c.ShouldBindJSON(&query)
			err := db.Where(&query).Find(&query).Error
			if err != nil {
				c.JSON(200, gin.H{
					"message": "很抱歉，没有找到你要查询的内容哦",
				})
			} else {
				c.JSON(200, query)
			}
		}
	} else {
		return func(c *gin.Context) {
			var query T
			var results []T
			c.ShouldBindJSON(&query)
			err := db.Where(&query).Find(&results).Error
			if err != nil {
				c.JSON(200, gin.H{
					"message": "很抱歉，没有找到你要查询的内容哦",
				})
			} else {
				c.JSON(200, results)
			}
		}
	}
}

// 验证码服务，使用redis存储
func handleMailSendCaptcha(config Config, rdb *redis.Client) func(*gin.Context) {
	return func(c *gin.Context) {
		mailTo := c.Query("mail") // TODO: 增加对mail的合法性验证
		codeId, code := GenerateCaptcha(config.CaptchaLength)
		rdb.Set(context.Background(), codeId, code, time.Minute*time.Duration(config.CaptchaAlive))
		e := email.NewEmail()
		e.From = "Get <" + config.MailUserName + ">"
		e.To = []string{mailTo}
		e.Subject = "验证码"
		e.HTML = []byte("你的验证码为：<h1>" + code + "</h1>")
		err := e.SendWithTLS(config.MailServer+config.MailServerPort, smtp.PlainAuth("", config.MailUserName, config.MailPassword, config.MailServer),
			&tls.Config{InsecureSkipVerify: true, ServerName: config.MailServer})
		if err != nil {
			c.Error(err)
		} else {
			c.JSON(200, gin.H{"id": codeId, "code": code})
		}
	}
}
func handleCaptchaVerify(rdb *redis.Client) func(*gin.Context) {
	return func(c *gin.Context) {
		// TODO: 增加对id和code的合法性验证
		codeId := c.Query("id")
		code := c.Query("code")
		if VerifyCaptcha(codeId, code, rdb) {
			c.JSON(200, gin.H{
				"message": "验证码正确",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "验证码错误",
			})
		}
	}
}
func AddCaptchaAPI(router gin.IRouter, path string, config Config, rdb *redis.Client) *gin.RouterGroup {
	return APIBuilder(router, func(group *gin.RouterGroup) *gin.RouterGroup {
		group.POST("/gen_captcha", handleMailSendCaptcha(config, rdb))
		group.POST("/verify_captcha", handleCaptchaVerify(rdb))
		return group
	})(router, path)
}
