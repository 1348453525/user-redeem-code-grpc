package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/1348453525/user-redeem-code-grpc/user-srv/global"
	"github.com/1348453525/user-redeem-code-grpc/user-srv/handler"
	"github.com/1348453525/user-redeem-code-grpc/user-srv/initialize"
	"github.com/1348453525/user-redeem-code-grpc/user-srv/pkg/util"
	proto "github.com/1348453525/user-redeem-code-grpc/user-srv/proto/user"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	// 运行时指定地址和端口
	// IP := *flag.String("ip", "192.168.50.102", "ip地址")
	// Port := *flag.Int("port", 6060, "端口号")
	// flag.Parse()

	// 使用配置文件：使用 docker 部署的 consul，如果服务地址为 127.0.0.1，健康检查会失败
	IP := global.Config.Server.Addr
	Port := global.Config.Server.Port

	// 动态获取端口
	var err error
	Port, err = util.GetFreePort(IP, Port)
	if err != nil {
		panic(err)
	}
	zap.S().Infof("启动服务，监听地址：%s:%d", IP, Port)

	// 添加服务
	serviceID := initialize.ServiceRegister(IP, Port)

	// 监听端口
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", IP, Port))
	if err != nil {
		panic("failed to listen: " + err.Error())
	}

	// 创建 grpc 服务器
	server := grpc.NewServer()

	// 注册用户服务
	proto.RegisterUserServer(server, &handler.User{})
	// 注册健康检查服务
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 启动服务
	go func() {
		if err = server.Serve(listen); err != nil {
			panic("failed to start grpc: " + err.Error())
		}
	}()

	// 监听系统信号
	quit := make(chan os.Signal, 1)
	// 监听 SIGINT（Ctrl+C）、SIGTERM（kill 命令）
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // 阻塞等待信号

	// 删除服务
	initialize.ServiceDeregister(serviceID, IP, Port)
	zap.S().Info("shutting down server...")

	// 优雅关闭 gRPC 服务器
	server.GracefulStop()
	zap.S().Info("server stopped")
}

func init() {
	cfgFile := "./config.yaml"
	// 初始化配置文件
	initialize.InitConfig(cfgFile)
	// 初始化日志
	initialize.InitLogger()
	// 初始化数据库
	initialize.InitDB()
	// 初始化 Redis
	initialize.InitRedis()
}
