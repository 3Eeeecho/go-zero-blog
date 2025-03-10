package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		Dsn         string // 数据库连接字符串
		TablePrefix string // 表前缀
	}
	Redis redis.RedisConf // Redis 配置，使用 go-zero 内置结构体
	App   struct {        // 应用相关配置
		PageSize        int
		PrefixUrl       string
		RuntimeRootPath string
		ImageSavePath   string
		ImageMaxSize    int
		ImageAllowExts  []string
		LogSavePath     string
		LogSaveName     string
		LogFileExt      string
		TimeFormat      string
		ExportSavePath  string
		QrCodeSavePath  string
	}
	User struct {
		AccessExpire int64
	}
}
