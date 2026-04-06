package global

import (
	"github.com/1348453525/user-redeem-code-grpc/user-srv/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Client
)
