package initialize

import (
	"fmt"

	"github.com/1348453525/user-redeem-code-grpc/user-srv/global"
	"go.uber.org/zap"
)

func InitLogger() {
	var (
		logger *zap.Logger
		err    error
	)

	// TODO 待优化 zap 初始化
	if global.Config.Server.Mode == "dev" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		panic(fmt.Errorf("init logger failed, err: %w", err))
	}

	zap.ReplaceGlobals(logger)
}
