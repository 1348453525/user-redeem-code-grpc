package initialize

import (
	"context"

	"github.com/1348453525/user-redeem-code-grpc/user-api/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitRedis() {
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Addr,
		Password: global.Config.Redis.Password,
		DB:       0,
	})

	if _, err := global.Redis.Ping(context.Background()).Result(); err != nil {
		zap.L().Error("redis connect failed", zap.Error(err))
	} else {
		zap.L().Info("redis connected successfully",
			zap.String("Addr", global.Config.Redis.Addr),
		)
	}
}
