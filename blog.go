package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/3Eeeecho/go-zero-blog/internal/config"
	"github.com/3Eeeecho/go-zero-blog/internal/handler"
	"github.com/3Eeeecho/go-zero-blog/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/blog-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	var logc logx.LogConf
	conf.MustLoad("etc/logConfig.yaml", &logc)
	logx.MustSetup(logc)

	//仅测试时使用
	logx.AddWriter(logx.NewWriter(os.Stdout))

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
