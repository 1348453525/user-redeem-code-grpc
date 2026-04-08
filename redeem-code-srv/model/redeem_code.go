package model

import (
	"time"

	"github.com/1348453525/user-redeem-code-grpc/redeem-code-srv/global"
)

const TableNameRedeemCode = "redeem_code"

// RedeemCode mapped from table <redeem_code>
type RedeemCode struct {
	ID                int64      `json:"id"`
	RedeemCodeBatchID int64      `json:"redeem_code_batch_id"` // 兑换码批次 id
	Title             string     `json:"title"`                // 标题
	Value             string     `json:"value"`                // 兑换码
	UsageLimit        int32      `json:"usage_limit"`          // 可使用次数
	UsedCount         int32      `json:"used_count"`           // 已使用数量
	ExpirationAt      time.Time  `json:"expiration_at"`        // 过期时间
	IsDel             int32      `json:"is_del"`               // 是否删除：0，默认；1，已删除；2，未删除；
	DeletedAt         *time.Time `json:"deleted_at"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

// TableName RedeemCode's table name
func (*RedeemCode) TableName() string {
	return TableNameRedeemCode
}

func (m *RedeemCode) GetByID(id int64) error {
	result := global.DB.Where("id = ?", id).First(m)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *RedeemCode) GetList(page int32, pageSize int32) (list []*RedeemCode, count int64) {
	global.DB.Model(&RedeemCode{}).Count(&count).Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Order("id desc").Find(&list)
	return
}
