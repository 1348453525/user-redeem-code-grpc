package model

import (
	"time"
)

const TableNameRedeemCodeRecord = "redeem_code_record"

// RedeemCodeRecord mapped from table <redeem_code_record>
type RedeemCodeRecord struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"user_id"`        // 用户 id
	RedeemCodeID int64     `json:"redeem_code_id"` // 兑换码 id
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// TableName RedeemCodeRecord's table name
func (*RedeemCodeRecord) TableName() string {
	return TableNameRedeemCodeRecord
}
