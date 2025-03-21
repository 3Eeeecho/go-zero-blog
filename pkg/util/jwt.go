package util

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(jwtSecret []byte, username string, expires int) (string, error) {
	expireTime := time.Now().Add(time.Duration(expires) * time.Second)
	claims := Claims{
		username,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expireTime.Unix(),
			Issuer:    "3Eeeecho",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return token, err
}

func ParseToken(jwtSecret []byte, tokenStr string) (*Claims, error) {
	if tokenStr == "" {
		return nil, errors.New("token cannot be empty")
	}

	// 从配置文件读取密钥
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
