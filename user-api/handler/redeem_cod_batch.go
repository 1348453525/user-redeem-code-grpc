package handler

import (
	"errors"
	"strconv"

	"github.com/1348453525/user-redeem-code-grpc/user-api/entity"
	"github.com/1348453525/user-redeem-code-grpc/user-api/logic"
	"github.com/1348453525/user-redeem-code-grpc/user-api/pkg/helper"
	"github.com/1348453525/user-redeem-code-grpc/user-api/pkg/result"
	"github.com/1348453525/user-redeem-code-grpc/user-api/pkg/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RedeemCodeBatch struct{}

func NewRedeemCodeBatch() *RedeemCodeBatch {
	return &RedeemCodeBatch{}
}

func (h *RedeemCodeBatch) Create(c *gin.Context) {
	// 获取用户 id
	id, err := helper.GetUserID(c)
	if err != nil {
		result.Error(c, 400, entity.ErrParam.Error())
		return
	}

	// 接收参数
	var dto entity.CreateRedeemCodeBatchDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		result.Error(c, 400, entity.ErrParam.Error())
		return
	}

	// 验证参数
	if ok := util.Validate(c, dto); !ok {
		return
	}

	// 处理逻辑
	resp, err := logic.NewRedeemCodeBatchLogic().Create(c, id, &dto)
	if err != nil {
		result.Error(c, 500, err.Error())
		return
	}
	result.Success(c, resp)
}

func (h *RedeemCodeBatch) Detail(c *gin.Context) {
	// 接收参数
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		result.Error(c, 400, entity.ErrParam.Error())
		return
	}

	// 处理逻辑
	resp, err := logic.NewRedeemCodeBatchLogic().Detail(c, id)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			result.Error(c, 500, entity.ErrInternal.Error())
			return
		}
		result.Success(c)
		return
	}
	result.Success(c, resp)
}

func (h *RedeemCodeBatch) GetList(c *gin.Context) {
	// 接收参数
	var dto entity.GetRedeemCodeBatchListDto
	if err := c.ShouldBindQuery(&dto); err != nil {
		result.Error(c, 400, entity.ErrParam.Error())
		return
	}

	// 验证参数
	if ok := util.Validate(c, dto); !ok {
		return
	}

	// 处理逻辑
	resp, err := logic.NewRedeemCodeBatchLogic().GetList(c, &dto)
	if err != nil {
		result.Error(c, 500, err.Error())
		return
	}
	result.Success(c, resp)
}

func (h *RedeemCodeBatch) Update(c *gin.Context) {
	// 接收参数
	var dto entity.UpdateRedeemCodeBatchDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		result.Error(c, 400, entity.ErrParam.Error())
		return
	}

	// 验证参数
	if ok := util.Validate(c, dto); !ok {
		return
	}

	// 处理逻辑
	err := logic.NewRedeemCodeBatchLogic().Update(c, &dto)
	if err != nil {
		result.Error(c, 500, err.Error())
		return
	}
	result.Success(c)
}

func (h *RedeemCodeBatch) Delete(c *gin.Context) {
	// 接收参数
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		result.Error(c, 400, entity.ErrParam.Error())
		return
	}

	// 处理逻辑
	err = logic.NewRedeemCodeBatchLogic().Delete(c, id)
	if err != nil {
		result.Error(c, 500, err.Error())
		return
	}
	result.Success(c)
}
