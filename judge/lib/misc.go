package lib

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 为Gin router 添加CRUD
func APIBuilder(router gin.IRouter, handlers ...func(*gin.RouterGroup) *gin.RouterGroup) func(gin.IRouter, string) *gin.RouterGroup {
	return func(router gin.IRouter, path string) *gin.RouterGroup {
		group := router.Group(path)
		for _, handler := range handlers {
			group = handler(group)
		}
		return group
	}
}

// 生成token
func GenerateToken(userClaim *UserClaim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token
func ParseToken(token string) (*UserClaim, error) {
	userClaim := new(UserClaim)
	claim, err := jwt.ParseWithClaims(token, userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if !claim.Valid {
		return nil, errors.New("token validation failed")
	}
	return userClaim, nil
}

// 验证token
func (uc *UserClaim) Valid() error {
	// 过期时间为一周
	if time.Now().Unix() > uc.Expire.Unix() {
		return errors.New("token expired")
	}
	return nil
}

// 生成uuid
func GenerateUUID() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		return ""
	}
	return uuid.String()
}

// 生成验证码
// 返回验证码id，验证码
func GenerateCaptcha(length int) (string, string) {
	var chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	return GenerateUUID(), func() string {
		rand.NewSource(time.Now().UnixNano())
		result := make([]byte, length)
		for i := range result {
			result[i] = chars[rand.Intn(len(chars))]
		}
		return string(result)
	}()
}

// 验证验证码
func VerifyCaptcha(id string, captcha string, db *redis.Client) bool {
	ctx := context.Background()
	value := db.Get(ctx, id).String()
	return captcha == value
}

func NewRedis(config *RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
	})
}

func NewDB(config *DatabaseConfig, migrator func(*gorm.DB) error) *gorm.DB {
	var db *gorm.DB
	var err error
	switch config.Type {
	case "mysql":
		dsn := config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.DB + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case "sqlite":
		db, err = gorm.Open(mysql.Open(config.DB), &gorm.Config{})
	}
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if config.Migrate {
		if migrator == nil {
			log.Fatalf("Migrator is nil")
		}
		if err = migrator(db); err != nil {
			log.Fatalf("Failed to migrate tables: %v", err)
		}
	}
	return db
}
