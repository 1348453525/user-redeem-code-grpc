// Package result 通用返回结构
package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 结构体
type result struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回的数据
	Error   interface{} `json:"error"`   // 返回的错误数据
}

// Success 返回成功
// 参数顺序：data，code，message
func Success(c *gin.Context, params ...interface{}) {
	res := result{}

	// 默认值
	res.Code = 200
	res.Message = "success"
	res.Data = gin.H{}
	res.Error = gin.H{}

	// 根据传入参数数量按顺序赋值
	if len(params) >= 1 && params[0] != nil {
		res.Data = params[0]

	}
	if len(params) >= 2 && params[1] != nil {
		if code, ok := params[1].(int); ok {
			res.Code = code
		}

	}
	if len(params) >= 3 && params[2] != nil {
		if message, ok := params[2].(string); ok {
			res.Message = message
		}
	}

	c.JSON(http.StatusOK, res)
}

// Error 返回错误
// 参数顺序：code，message，data，error
func Error(c *gin.Context, params ...interface{}) {
	res := result{}

	// 默认值
	res.Code = 500
	res.Message = "error"
	res.Data = gin.H{}
	res.Error = gin.H{}

	// 根据传入参数数量按顺序赋值
	if len(params) >= 1 && params[0] != nil {
		if code, ok := params[0].(int); ok {
			res.Code = code
		}
	}
	if len(params) >= 2 && params[1] != nil {
		if message, ok := params[1].(string); ok {
			res.Message = message
		}
	}
	if len(params) >= 3 && params[2] != nil {
		res.Data = params[2]
	}
	if len(params) >= 4 && params[3] != nil {
		res.Error = params[3]
	}

	c.JSON(http.StatusOK, res)
}
