package initialize

import (
	"fmt"

	"github.com/1348453525/user-redeem-code-grpc/redeem-code-srv/global"
	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
)

func ServiceRegister(IP string, Port int) string {
	cfgConsul := api.DefaultConfig()
	cfgConsul.Address = fmt.Sprintf("%s:%d", global.Config.Consul.Host, global.Config.Consul.Port)
	client, err := api.NewClient(cfgConsul)
	if err != nil {
		panic(err)
	}

	// 健康检查
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", IP, Port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	// 注册
	ID := uuid.NewV4().String()
	registration := &api.AgentServiceRegistration{
		ID:      ID,
		Name:    global.Config.Server.Name,
		Tags:    global.Config.Server.Tags, // []string{"user-srv"}
		Port:    Port,
		Address: IP,
		Check:   check,
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	return ID
}

func ServiceDeregister(ID, IP string, Port int) {
	cfgConsul := api.DefaultConfig()
	cfgConsul.Address = fmt.Sprintf("%s:%d", global.Config.Consul.Host, global.Config.Consul.Port)
	client, err := api.NewClient(cfgConsul)
	if err != nil {
		panic(err)
	}

	err = client.Agent().ServiceDeregister(ID)
	if err != nil {
		panic(err)
	}
}
