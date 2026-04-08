package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/1348453525/user-redeem-code-grpc/user-api/global"
	"github.com/1348453525/user-redeem-code-grpc/user-api/initialize"
)

func main() {
	// 初始化 Gin 引擎和路由
	r := initialize.InitRouter()

	// 配置 HTTP 服务
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", global.Config.Server.Addr, global.Config.Server.Port),
		Handler: r,
	}

	// 启动服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 监听系统信号
	quit := make(chan os.Signal, 1)
	// 监听 SIGINT（Ctrl+C）、SIGTERM（kill 命令）
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // 阻塞等待信号
	log.Println("shutdown server ...")

	// 优雅关闭：设置 5 秒超时，处理完现有请求后退出
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}
	log.Println("server exited")
}

func init() {
	cfgFile := "./config.yaml"
	// 初始化配置
	initialize.InitConfig(cfgFile)
	// 初始化日志
	initialize.InitLogger()
	// 初始化数据库
	initialize.InitDB()
	// 初始化 Redis
	initialize.InitRedis()
	// 初始化 grpc 客户端
	// initialize.InitGrpcUserClient()
	initialize.InitGrpcUserClientUseLB()
	initialize.InitGrpcRedeemCodeClientUseLB()
}
