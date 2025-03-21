package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	"github.com/zeromicro/go-zero/core/logx"
)

// EncryptPassword 使用 AES 加密密码（客户端模拟）
func EncryptPassword(key []byte, plainPassword string) (string, error) {
	if len(key) != 32 {
		return "", errors.New("AES key must be 32 bytes")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plainBytes := []byte(plainPassword)
	padded := pad(plainBytes, aes.BlockSize)

	ciphertext := make([]byte, aes.BlockSize+len(padded)) // IV + 密文
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], padded)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptPassword 解密密码
func DecryptPassword(key []byte, encryptedPassword string) (string, error) {
	if len(key) != 32 {
		return "", errors.New("AES key must be 32 bytes")
	}

	ciphertext, err := base64.StdEncoding.DecodeString(encryptedPassword)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	logx.Infof("Ciphertext length after IV: %d", len(ciphertext)) // 添加日志

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	plainBytes, err := unpad(ciphertext)
	if err != nil {
		return "", err
	}

	return string(plainBytes), nil
}

// pad 填充到块大小
func pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := make([]byte, len(src)+padding)
	copy(padtext, src)
	for i := len(src); i < len(padtext); i++ {
		padtext[i] = byte(padding)
	}
	return padtext
}

// unpad 移除填充
func unpad(src []byte) ([]byte, error) {
	length := len(src)
	if length == 0 {
		return nil, errors.New("invalid padding")
	}
	padding := int(src[length-1])
	if padding > length {
		return nil, errors.New("invalid padding")
	}
	return src[:length-padding], nil
}
