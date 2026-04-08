package entity

import (
	"errors"
)

var (
	ErrRedeemCodeUsedUp  = errors.New("兑换码已用完")
	ErrRedeemCodeExpired = errors.New("兑换码已过期")
	ErrRedeemCodeInvalid = errors.New("兑换码无效")
)

type GetRedeemCodeListDto struct {
	// ID       int64 `json:"id" validate:"required,gte=1"` // batch id
	Page     int32 `form:"page" validate:"required,gte=1"`
	PageSize int32 `form:"page_size" validate:"required,gte=10"`
}

type UpdateRedeemCodeDto struct {
	ID    int64  `json:"id" validate:"required,gte=1"`
	Title string `json:"title"` // 标题
	// UsedCount    int32  `json:"used_count"`    // 已使用数量
	ExpirationAt string `json:"expiration_at"` // 过期时间
	IsDel        int32  `json:"is_del"`        // 是否删除：0，默认；1，已删除；2，未删除；
}

type UseRedeemCodeDto struct {
	RedeemCodeID int64 `json:"redeem_code_id" validate:"required,gte=1"`
	UserID       int64 `json:"user_id" validate:"required,gte=1"`
}
