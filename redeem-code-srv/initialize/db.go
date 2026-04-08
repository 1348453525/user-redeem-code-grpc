package initialize

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/1348453525/user-redeem-code-grpc/redeem-code-srv/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		global.Config.MySQL.User,
		global.Config.MySQL.Password,
		global.Config.MySQL.Host,
		global.Config.MySQL.Port,
		global.Config.MySQL.DB,
		global.Config.MySQL.Charset,
	)

	lg := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: lg,
	})
	if err != nil {
		zap.L().Error("gorm open failed", zap.Error(err))
	}

	// 连接池
	sqlDB, err := db.DB()
	if err != nil {
		zap.L().Error("get sql db failed", zap.Error(err))
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(global.Config.MySQL.MaxIdleConns)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(global.Config.MySQL.MaxOpenConns)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Minute * 5) // 防止空闲连接过长

	if err = sqlDB.Ping(); err != nil {
		zap.L().Error("database ping failed", zap.Error(err))
	} else {
		zap.L().Info("database connected successfully",
			zap.String("db", global.Config.MySQL.DB),
			zap.String("host", global.Config.MySQL.Host),
		)
	}

	global.DB = db
}
