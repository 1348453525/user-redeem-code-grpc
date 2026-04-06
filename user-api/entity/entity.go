package entity

import "errors"

var (
	ErrInternal = errors.New("内部错误")

	ErrParam           = errors.New("参数错误")
	ErrOperationFailed = errors.New("操作失败")

	ErrUserNotLogin = errors.New("用户未登录")
)

type ID64 struct {
	ID int64 `json:"id" validate:"required,gte=1"`
}
