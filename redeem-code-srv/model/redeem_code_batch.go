package model

import (
	"time"

	"github.com/1348453525/user-redeem-code-grpc/redeem-code-srv/global"
)

const TableNameRedeemCodeBatch = "redeem_code_batch"

// RedeemCodeBatch mapped from table <redeem_code_batch>
type RedeemCodeBatch struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`        // 批次名称
	Description string    `json:"description"`  // 批次描述
	UsageLimit  int32     `json:"usage_limit"`  // 可使用次数
	TotalCount  int32     `json:"total_count"`  // 生成数量
	UsedCount   int32     `json:"used_count"`   // 已使用数量
	StartedAt   time.Time `json:"started_at"`   // 开始时间
	EndedAt     time.Time `json:"ended_at"`     // 结束时间
	Status      int32     `json:"status"`       // 状态：0，默认；1，启用；2，停用；
	CreatorID   int64     `json:"creator_id"`   // 创建人 ID
	CreatorName string    `json:"creator_name"` // 创建人名称
	CreatedAt   time.Time `json:"created_at"`   // 创建时间
	UpdatedAt   time.Time `json:"updated_at"`   // 更新时间
}

// TableName RedeemCodeBatch's table name
func (*RedeemCodeBatch) TableName() string {
	return TableNameRedeemCodeBatch
}

func (m *RedeemCodeBatch) GetByID(id int64) error {
	result := global.DB.Where("id = ?", id).First(m)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *RedeemCodeBatch) GetList(page int32, pageSize int32) (list []*RedeemCodeBatch, count int64) {
	global.DB.Model(&RedeemCodeBatch{}).Count(&count).Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Order("id desc").Find(&list)
	return
}
