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

	// 兑换码批次
	redeemCodeBatch := handler.NewRedeemCodeBatch()
	redeemCodeBatchGroup := r.Group("/RedeemCodeBatch")
	redeemCodeBatchGroup.Use(middleware.JWTAuthMiddleware())
	{
		redeemCodeBatchGroup.POST("/Create", redeemCodeBatch.Create)   // 创建兑换码批次
		redeemCodeBatchGroup.GET("/Detail", redeemCodeBatch.Detail)    // 获取兑换码批次详情
		redeemCodeBatchGroup.GET("/GetList", redeemCodeBatch.GetList)  // 获取兑换码批次列表
		redeemCodeBatchGroup.PUT("/Update", redeemCodeBatch.Update)    // 更新兑换码批次
		redeemCodeBatchGroup.DELETE("/Delete", redeemCodeBatch.Delete) // 删除兑换码批次
	}

	// 兑换码
	redeemCode := handler.NewRedeemCode()
	redeemCodeGroup := r.Group("/RedeemCode")
	redeemCodeGroup.Use(middleware.JWTAuthMiddleware())
	{
		redeemCodeGroup.GET("/Detail", redeemCode.Detail)    // 获取兑换码详情
		redeemCodeGroup.GET("/GetList", redeemCode.GetList)  // 获取兑换码列表
		redeemCodeGroup.PUT("/Update", redeemCode.Update)    // 更新兑换码
		redeemCodeGroup.DELETE("/Delete", redeemCode.Delete) // 删除兑换码
		redeemCodeGroup.POST("/Use", redeemCode.Use)         // 使用兑换码
	}
}
