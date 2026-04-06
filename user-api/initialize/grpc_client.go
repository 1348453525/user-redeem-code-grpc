package initialize

import (
	"context"
	"fmt"
	"time"

	"github.com/1348453525/user-redeem-code-grpc/user-api/global"
	proto "github.com/1348453525/user-redeem-code-grpc/user-api/proto/user"
	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// 初始化 grpc 客户端
func InitGrpcUserClientUseLB() {
	// 连接 grpc 服务端
	target := fmt.Sprintf(
		"consul://%s:%d/%s?wait=10s",
		global.Config.Consul.Host,
		global.Config.Consul.Port,
		global.Config.UserSrv.Name,
	)
	conn, err := grpc.Dial(
		target,
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Errorw(
			"[InitClient] 连接 【用户服务失败】",
			"err", err.Error(),
		)
	}

	// 健康检查
	if !HealthCheck(conn) {
		zap.S().Errorw(
			"[InitClient] 【健康检查失败】",
			"service", global.Config.UserSrv.Name,
		)
	} else {
		zap.S().Info("[InitClient] 链接【用户服务成功】")
	}

	// 创建 grpc 客户端
	global.UserClient = proto.NewUserClient(conn)
}

// 初始化 grpc 客户端
func InitGrpcUserClient() {
	// 从配置文件获取
	// address := global.Config.UserSrv.Address
	// port := global.Config.UserSrv.Port

	// 从服务发现获取
	address, port := ServiceDiscovery()

	// 连接 grpc 服务端
	target := fmt.Sprintf("%s:%d", address, port)
	conn, err := grpc.Dial(
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		zap.S().Errorw(
			"[InitClient] 连接 【用户服务失败】",
			"err", err.Error(),
		)
	}

	// 健康检查
	if !HealthCheck(conn) {
		zap.S().Errorw(
			"[InitClient] 【健康检查失败】",
			"service", global.Config.UserSrv.Name,
		)
	} else {
		zap.S().Info("[InitClient] 链接【用户服务成功】")
	}

	// 创建 grpc 客户端
	global.UserClient = proto.NewUserClient(conn)
}

// 健康检查
func HealthCheck(conn *grpc.ClientConn) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	client := grpc_health_v1.NewHealthClient(conn)
	resp, err := client.Check(ctx, &grpc_health_v1.HealthCheckRequest{
		Service: "", // proto.User
	})

	if err != nil || resp.Status != grpc_health_v1.HealthCheckResponse_SERVING {
		return false
	}

	zap.S().Infof("[HealthCheck] 服务【%s】状态正常", global.Config.UserSrv.Name)
	return true
}

// 服务发现
func ServiceDiscovery() (string, int) {
	cfgConsul := api.DefaultConfig()
	cfgConsul.Address = fmt.Sprintf("%s:%d", global.Config.Consul.Host, global.Config.Consul.Port)
	client, err := api.NewClient(cfgConsul)
	if err != nil {
		zap.S().Errorw(
			"[ServiceDiscovery] 创建 Consul 客户端失败",
			"err", err.Error(),
		)
		return "", 0
	}

	filter := fmt.Sprintf("Service==\"%s\"", global.Config.UserSrv.Name)
	data, err := client.Agent().ServicesWithFilter(filter)
	if err != nil {
		zap.S().Errorw(
			"[ServiceDiscovery] 查询服务失败",
			"err", err.Error(),
		)
		return "", 0
	}

	address := ""
	port := 0
	for k, v := range data {
		zap.S().Infof("[ServiceDiscovery] serviceID:%s serviceName:%s %s:%d", k, v.Service, v.Address, v.Port)
		address = v.Address
		port = v.Port
	}

	return address, port
}
