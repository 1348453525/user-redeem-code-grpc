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
