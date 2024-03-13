package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"xyz.xeonds/nano-oj/model"
)

func GenerateToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"role": user.AccountInfo.Permission,
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
