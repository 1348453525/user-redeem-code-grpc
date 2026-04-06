package router

import (
	"github.com/1348453525/user-redeem-code-grpc/user-api/handler"
	"github.com/1348453525/user-redeem-code-grpc/user-api/middleware"
	"github.com/gin-gonic/gin"
)

func RouterGroup(r *gin.RouterGroup) {
	// user handler
	user := handler.NewUser()

	// auth 注册、登录、退出
	r.POST("/Register", user.Register)                            // 注册
	r.POST("/Login", user.Login)                                  // 登录
	r.GET("/Logout", middleware.JWTAuthMiddleware(), user.Logout) // 退出

	// 用户
	userGroup := r.Group("/User")
	userGroup.Use(middleware.JWTAuthMiddleware())
	{
		// userGroup.GET("/Logout", user.Logout)    // 退出
		userGroup.GET("/Info", user.Info)        // 获取用户信息
		userGroup.GET("/GetList", user.GetList)  // 获取用户列表
		userGroup.PUT("/Update", user.Update)    // 更新用户信息
		userGroup.DELETE("/Delete", user.Delete) // 删除用户
	}
}
