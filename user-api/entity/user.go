package entity

import (
	"errors"
)

var (
	ErrPasswordNotMatch = errors.New("密码不匹配")
	ErrUserExisted      = errors.New("用户已存在")
	ErrUserDisabled     = errors.New("账户已停用")
	ErrPasswordError    = errors.New("密码错误")
)

type RegisterDto struct {
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
	Nickname        string `json:"nickname" validate:"required"`
	Mobile          string `json:"mobile" validate:"required"`
	Gender          int32  `json:"gender" validate:"required"`
	Birthday        string `json:"birthday" validate:"required"`
}

type GetUserListDto struct {
	Page     int32 `form:"page" validate:"required,gte=1"`
	PageSize int32 `form:"page_size" validate:"required,gte=10"`
}

type UpdateUserDto struct {
	ID       int64  `json:"id" validate:"required,gte=1"`
	Username string `json:"username" validate:"required"`
	Nickname string `json:"nickname" validate:"required"`
	Mobile   string `json:"mobile" validate:"required"`
	Gender   int32  `json:"gender" validate:"required"`
	Birthday string `json:"birthday" validate:"required"`
}

type LoginDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
