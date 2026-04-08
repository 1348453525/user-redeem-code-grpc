package global

import (
	"github.com/1348453525/user-redeem-code-grpc/user-api/config"
	rcproto "github.com/1348453525/user-redeem-code-grpc/user-api/proto/redeem_code"
	proto "github.com/1348453525/user-redeem-code-grpc/user-api/proto/user"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config           config.Config
	DB               *gorm.DB
	Redis            *redis.Client
	UserClient       proto.UserClient
	RedeemCodeClient rcproto.RedeemCodeClient
)
