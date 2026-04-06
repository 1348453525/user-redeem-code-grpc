package handler

import (
	"errors"

	"github.com/1348453525/user-redeem-code-grpc/user-api/entity"
	"github.com/1348453525/user-redeem-code-grpc/user-api/logic"
	"github.com/1348453525/user-redeem-code-grpc/user-api/pkg/helper"
	"github.com/1348453525/user-redeem-code-grpc/user-api/pkg/result"
	"github.com/1348453525/user-redeem-code-grpc/user-api/pkg/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (h *User) Register(c *gin.Context) {
	// 接收参数
	var dto entity.RegisterDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		result.Error(c, 400, entity.ErrParam.Error())
		return
	}

	// 验证参数
	if ok := util.Validate(c, dto); !ok {
		return
	}

	// 处理逻辑
	resp, err := logic.NewUserLogic().Register(c, &dto)
	if err != nil {
		result.Error(c, 500, err.Error())
		return
	}
	result.Success(c, resp)
}

func (h *User) Login(c *gin.Context) {
	// 接收参数
	var dto entity.LoginDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		result.Error(c, 400, entity.ErrParam.Error())
		return
	}

	// 验证参数
	if ok := util.Validate(c, dto); !ok {
		return
	}

	// 处理逻辑
	resp, err := logic.NewUserLogic().Login(c, &dto)
	if err != nil {
		result.Error(c, 500, err.Error())
		return
	}
	result.Success(c, resp)
}

func (h *User) Logout(c *gin.Context) {
	result.Success(c)
}

func (h *User) Info(c *gin.Context) {
	// 接收参数
	// id, _ := strconv.Atoi(c.Query("id"))
	// if id <= 0 {
	// 	result.Error(c, 400, entity.ErrParam.Error())
	// 	return
	// }
	id, err := helper.GetUserID(c)
	if err != nil {
		result.Error(c, 400, entity.ErrParam.Error())
		return
	}

	// 处理逻辑
	resp, err := logic.NewUserLogic().Info(c, id)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			result.Error(c, 500, entity.ErrInternal.Error())
		}
		result.Success(c)
		return
	}
	result.Success(c, resp)
}

func (h *User) GetList(c *gin.Context) {
	// 接收参数
	var dto entity.GetUserListDto
	if err := c.ShouldBindQuery(&dto); err != nil {
		result.Error(c, 400, entity.ErrParam.Error())
		return
	}

	// 验证参数
	if ok := util.Validate(c, dto); !ok {
		return
	}

	// 处理逻辑
	resp, err := logic.NewUserLogic().GetList(c, &dto)
	if err != nil {
		result.Error(c, 500, err.Error())
		return
	}
	result.Success(c, resp)
}

func (h *User) Update(c *gin.Context) {
	id, err := helper.GetUserID(c)
	if err != nil {
		result.Error(c, 400, entity.ErrParam.Error())
		return
	}

	// 接收参数
	var dto entity.UpdateUserDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		result.Error(c, 400, entity.ErrParam.Error())
		return
	}

	// 验证参数
	if ok := util.Validate(c, dto); !ok {
		return
	}
	if id != dto.ID {
		result.Error(c, 400, entity.ErrParam.Error())
		return
	}

	// 处理逻辑
	err = logic.NewUserLogic().Update(c, &dto)
	if err != nil {
		result.Error(c, 500, err.Error())
		return
	}
	result.Success(c)
}

func (h *User) Delete(c *gin.Context) {
	// 接收参数
	// id, _ := strconv.Atoi(c.Query("id"))
	// if id <= 0 {
	// 	result.Error(c, 400, entity.ErrParam.Error())
	// 	return
	// }
	id, err := helper.GetUserID(c)
	if err != nil {
		result.Error(c, 400, entity.ErrParam.Error())
		return
	}

	// 处理逻辑
	err = logic.NewUserLogic().Delete(c, id)
	if err != nil {
		result.Error(c, 500, err.Error())
		return
	}
	result.Success(c)
}
