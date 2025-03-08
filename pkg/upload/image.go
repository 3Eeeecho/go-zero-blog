package upload

import (
	"mime/multipart"
)

// GetImageFullUrl 生成图片的完整访问 URL
// name: 图片文件名
// 返回: 完整的图片 URL，格式为 "ImagePrefixUrl/ImageSavePath/name"
func GetImageFullUrl(name string) string {
	return ""
}

// GetImageName 生成唯一的图片文件名
// name: 原始文件名
// 返回: 唯一的图片文件名，格式为 "MD5(文件名) + 扩展名"
func GetImageName(name string) string {
	return ""
}

// GetImagePath 获取图片的存储路径
// 返回: 配置文件中设置的图片存储路径
func GetImagePath() string {
	return ""
}

// GetImageFullPath 获取图片的完整存储路径
// 返回: 运行时根路径 + 图片存储路径
func GetImageFullPath() string {
	return ""
}

// CheckImageExt 检查图片扩展名是否允许
// fileName: 图片文件名
// 返回: 如果扩展名允许返回 true，否则返回 false
func CheckImageExt(fileName string) bool {
	return false
}

// CheckImageSize 检查图片大小是否超过限制
// f: 上传的文件
// 返回: 如果文件大小未超过限制返回 true，否则返回 false
func CheckImageSize(f multipart.File) bool {
	return false
}

// CheckImage 检查图片存储路径是否存在并验证权限
// src: 图片存储的相对路径
// 返回: 如果检查通过返回 nil，否则返回错误
func CheckImage(src string) error {
	return nil
}
