package middleware

import (
	"net/http"
	"strings"

	"github.com/1348453525/user-redeem-code-grpc/user-api/pkg/jwt"
	"github.com/1348453525/user-redeem-code-grpc/user-api/pkg/result"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware Gin JWT认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Header获取 Authorization: Bearer <token>
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			result.Error(
				c,
				http.StatusUnauthorized,
				"未提供token",
			)
			c.Abort()
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// 解析
		claims, err := jwt.ParseToken(tokenStr)
		if err != nil {
			result.Error(
				c,
				http.StatusUnauthorized,
				"token无效或已过期",
			)
			c.Abort()
			return
		}

		// 校验用户状态
		// var user model.User
		// if err = user.GetByID(claims.ID); err != nil {
		// 	result.Error(
		// 		c,
		// 		http.StatusUnauthorized,
		// 		"用户不存在",
		// 	)
		// 	c.Abort()
		// 	return
		// }
		// if user.IsDel == 1 {
		// 	result.Error(
		// 		c,
		// 		http.StatusUnauthorized,
		// 		entity.ErrUserDisabled.Error(),
		// 	)
		// 	c.Abort()
		// 	return
		// }

		// 将用户信息存入上下文
		c.Set("userID", claims.ID)
		// c.Set("user", &user)
		c.Next()
	}
}
