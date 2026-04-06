package initialize

import (
	"net/http"

	"github.com/1348453525/user-redeem-code-grpc/user-api/global"
	"github.com/1348453525/user-redeem-code-grpc/user-api/router"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化 Gin 引擎
	r := gin.Default()

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "ok",
			"data": map[string]string{
				"name":    "user-redeem-code-grpc/user-api",
				"version": "v1.0.0",
			},
		})
	})

	// 404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "404 not found",
		})
	})

	// 路由分组
	v1 := r.Group("/v1")
	router.RouterGroup(v1)

	return r
}
