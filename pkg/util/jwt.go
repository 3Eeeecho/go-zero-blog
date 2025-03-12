package util

import (
	"errors"
	"time"

	"github.com/3Eeeecho/go-zero-blog/internal/config"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// TODO JwtSecert测试使用233,应该替换更安全的
func GenerateToken(c config.Config, username string, expires int) (string, error) {
	expireTime := time.Now().Add(time.Duration(expires) * time.Second)
	claims := Claims{
		username,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expireTime.Unix(),
			Issuer:    "3Eeeecho",
		},
	}

	jwtSecret := []byte(c.User.JwtSecret)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return token, err
}

func ParseToken(c config.Config, tokenStr string) (*Claims, error) {
	if tokenStr == "" {
		return nil, errors.New("token cannot be empty")
	}

	// 从配置文件读取密钥
	jwtSecret := []byte(c.User.JwtSecret)
	if len(jwtSecret) == 0 {
		return nil, errors.New("JWT secret is not configured")
	}

	tokenClaims, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {
		// 验证签名方法
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
