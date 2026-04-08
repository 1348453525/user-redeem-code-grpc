package entity

import (
	"errors"
)

var (
	ErrRedeemCodeUsedUp  = errors.New("兑换码已用完")
	ErrRedeemCodeExpired = errors.New("兑换码已过期")
	ErrRedeemCodeInvalid = errors.New("兑换码无效")
)
