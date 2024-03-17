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
	"xyz.xeonds/nano-oj/model"
)

func AddCRUD[T any](router gin.IRouter, path string, db *gorm.DB) *gin.RouterGroup {
	return APIBuilder(router, func(group *gin.RouterGroup) *gin.RouterGroup {
		group.GET("", GetAll[T](db))
		group.GET("/:id", Get[T](db))
		group.POST("", Create[T](db))
		group.PUT("/:id", Update[T](db))
		group.DELETE("/:id", Delete[T](db))
		return group
	})(router, path)
}
func AddCRUDWithAuth[T any](router gin.IRouter, path string, db *gorm.DB, permLo, permHi int) *gin.RouterGroup {
	return APIBuilder(router, func(group *gin.RouterGroup) *gin.RouterGroup {
		group.GET("", GetAll[T](db))
		group.GET("/:id", Get[T](db))
		return group
	}, func(group *gin.RouterGroup) *gin.RouterGroup {
		group.Use(JWTMiddleware(AuthPermission(permLo, permHi)))
		group.POST("", Create[T](db))
		group.PUT("/:id", Update[T](db))
		group.DELETE("/:id", Delete[T](db))
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
func AddLoginAPI(router gin.IRouter, path string, db *gorm.DB) *gin.RouterGroup {
	return APIBuilder(router, func(group *gin.RouterGroup) *gin.RouterGroup {
		group.POST("/login", HandleLogin(db))
		group.POST("/register", HandleRegister(db))
		return group
	})(router, path)
}
func AddCaptchaAPI(router gin.IRouter, path string, conf1 MailConfig, conf2 CaptchaConfig, rdb *redis.Client) *gin.RouterGroup {
	return APIBuilder(router, func(group *gin.RouterGroup) *gin.RouterGroup {
		group.POST("/gen_captcha", HandleMailSendCaptcha(conf1, conf2, rdb))
		group.POST("/verify_captcha", HandleCaptchaVerify(rdb))
		return group
	})(router, path)
}

// handlers for gorm
func Create[T any](db *gorm.DB) func(c *gin.Context) {
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
func Get[T any](db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, d := c.Param("id"), new(T)
		if err := db.Where("id = ?", id).First(d).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, d)
		}
	}
}
func GetAll[T any](db *gorm.DB) func(c *gin.Context) {
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
func Update[T any](db *gorm.DB) func(c *gin.Context) {
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
func Delete[T any](db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var d T
		if err := db.Where("id = ?", id).Delete(&d).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, d)
		}
	}
}
func HandleFind[T any](queryProcess func(c *gin.Context) *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		query := new(T)
		err := queryProcess(c).First(query).Error
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Query Content Not Found",
			})
		} else {
			c.JSON(200, query)
		}
	}
}
func HandleFindAll[T any](queryProcess func(c *gin.Context) *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var query []T
		err := queryProcess(c).Find(&query).Error
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Query Content Not Found",
			})
		} else {
			c.JSON(200, query)
		}
	}
}

// 验证码服务，使用redis存储
func HandleMailSendCaptcha(mailConfig MailConfig, captchaConfig CaptchaConfig, rdb *redis.Client) func(*gin.Context) {
	return func(c *gin.Context) {
		mailTo := c.Query("mail") // TODO: 增加对mail的合法性验证
		codeId, code := GenerateCaptcha(captchaConfig.CaptchaLength)
		rdb.Set(context.Background(), codeId, code, time.Minute*time.Duration(captchaConfig.CaptchaAlive))
		e := email.NewEmail()
		e.From = "Get <" + mailConfig.MailServer + ">"
		e.To = []string{mailTo}
		e.Subject = "验证码"
		e.HTML = []byte("你的验证码为：<h1>" + code + "</h1>")
		err := e.SendWithTLS(mailConfig.MailServer+mailConfig.MailServerPort, smtp.PlainAuth("", mailConfig.MailUserName, mailConfig.MailPassword, mailConfig.MailServer),
			&tls.Config{InsecureSkipVerify: true, ServerName: mailConfig.MailServer})
		if err != nil {
			c.Error(err)
		} else {
			c.JSON(200, gin.H{"id": codeId, "code": code})
		}
	}
}
func HandleCaptchaVerify(rdb *redis.Client) func(*gin.Context) {
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

// login service
func HandleLogin(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		input, user := new(model.User), new(model.User)
		// user is already a pointer, so no need to use &user
		if err := c.ShouldBindJSON(input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Where("email = ?", input.Email).Find(user).Error; err != nil {
			log.Println("Find user by email failed: ", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User does not exist"})
			return
		}
		if err := CheckPasswordHash(input.Password, user.Password); err != nil {
			log.Println("Incorrect password: ", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
			return
		}
		token, err := GenerateToken(&UserClaim{
			Name:       user.Username,
			Expire:     time.Now().Add(time.Hour * 24),
			Permission: int(user.AccountInfo.Permission),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
func HandleRegister(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var user model.User
		var count int64
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Where("username = ?", user.Username).Find(new(model.User)).Count(&count).Error; count != 0 {
			log.Println("User already exists: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "username already exists"})
			return
		}
		user.Password = HashedPassword(user.Password)
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
			return
		}
		if user.ID == 1 { // if it is the first user, set it as admin
			user.AccountInfo.Permission = 0
			if err := db.Save(&user).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
	}
}
