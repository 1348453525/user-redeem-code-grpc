package logic

import (
	"context"
	"time"

	"github.com/1348453525/user-redeem-code-grpc/redeem-code-srv/entity"
	"github.com/1348453525/user-redeem-code-grpc/redeem-code-srv/global"
	"github.com/1348453525/user-redeem-code-grpc/redeem-code-srv/model"
	"github.com/1348453525/user-redeem-code-grpc/redeem-code-srv/pkg/helper"
	proto "github.com/1348453525/user-redeem-code-grpc/redeem-code-srv/proto/redeem_code"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type RedeemCodeLogic struct{}

func NewRedeemCodeLogic() *RedeemCodeLogic {
	return &RedeemCodeLogic{}
}

func (l *RedeemCodeLogic) Detail(c context.Context, r *proto.IDRequest) (*proto.RedeemCodeResponse, error) {
	// 查询数据
	var redeemCode model.RedeemCode
	if err := redeemCode.GetByID(r.Id); err != nil {
		return nil, err
	}

	// deleted_at=null，timestamppb.New(*redeemCode.DeletedAt)
	// panic: runtime error: invalid memory address or nil pointer dereference
	// 安全处理 DeletedAt
	var deletedAtPb *timestamppb.Timestamp
	if redeemCode.DeletedAt != nil {
		deletedAtPb = timestamppb.New(*redeemCode.DeletedAt)
	}

	// 返回数据
	return &proto.RedeemCodeResponse{
		Id:                redeemCode.ID,
		RedeemCodeBatchId: redeemCode.RedeemCodeBatchID,
		Title:             redeemCode.Title,
		Value:             redeemCode.Value,
		UsageLimit:        redeemCode.UsageLimit,
		UsedCount:         redeemCode.UsedCount,
		ExpirationAt:      timestamppb.New(redeemCode.ExpirationAt),
		IsDel:             redeemCode.IsDel,
		DeletedAt:         deletedAtPb,
		CreatedAt:         timestamppb.New(redeemCode.CreatedAt),
		UpdatedAt:         timestamppb.New(redeemCode.UpdatedAt),
	}, nil
}

func (l *RedeemCodeLogic) GetList(c context.Context, r *proto.GetListRequest) (*proto.GetRedeemCodeListResponse, error) {
	var redeemCodeModel model.RedeemCode
	list, count := redeemCodeModel.GetList(r.Page, r.PageSize)
	resp := &proto.GetRedeemCodeListResponse{
		Page:     r.Page,
		PageSize: r.PageSize,
		Total:    count,
	}
	for _, v := range list {
		// 安全处理 DeletedAt
		var deletedAtPb *timestamppb.Timestamp
		if v.DeletedAt != nil {
			deletedAtPb = timestamppb.New(*v.DeletedAt)
		}
		redeemCode := &proto.RedeemCodeResponse{
			Id:                v.ID,
			RedeemCodeBatchId: v.RedeemCodeBatchID,
			Title:             v.Title,
			Value:             v.Value,
			UsageLimit:        v.UsageLimit,
			UsedCount:         v.UsedCount,
			ExpirationAt:      timestamppb.New(v.ExpirationAt),
			IsDel:             v.IsDel,
			DeletedAt:         deletedAtPb,
			CreatedAt:         timestamppb.New(v.CreatedAt),
			UpdatedAt:         timestamppb.New(v.UpdatedAt),
		}
		resp.Data = append(resp.Data, redeemCode)
	}
	return resp, nil
}

func (l *RedeemCodeLogic) Update(c context.Context, r *proto.UpdateRedeemCodeRequest) error {
	redeemCodeModel := model.RedeemCode{
		ID: r.Id,
	}
	if r.Title != "" {
		redeemCodeModel.Title = r.Title
	}
	if r.ExpirationAt != "" {
		expirationAt, err := helper.ParseDatetime(r.ExpirationAt)
		if err != nil {
			zap.L().Error("解析过期时间失败：", zap.Error(err))
			return entity.ErrParam
		}
		redeemCodeModel.ExpirationAt = *expirationAt
	}
	if r.IsDel != 0 {
		redeemCodeModel.IsDel = r.IsDel
	}
	if result := global.DB.Model(&model.RedeemCode{}).Where("id=?", r.Id).Updates(&redeemCodeModel); result.Error != nil {
		zap.L().Error("更新兑换码失败：", zap.Error(result.Error), zap.Int64("id", r.Id))
		return entity.ErrInternal
	}
	return nil
}

func (l *RedeemCodeLogic) Delete(c context.Context, r *proto.IDRequest) error {
	if result := global.DB.Model(&model.RedeemCode{}).Where("id=?", r.Id).Update("is_del", 1); result.Error != nil {
		zap.L().Error("删除兑换码失败：", zap.Error(result.Error), zap.Int64("id", r.Id))
		return entity.ErrInternal
	}
	return nil
}

func (l *RedeemCodeLogic) Use(c context.Context, r *proto.UseRedeemCodeRequest) error {
	zap.L().Info("兑换码使用开始：", zap.Int64("user_id", r.UserId), zap.Int64("redeem_code_id", r.RedeemCodeId))
	// 开启事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			zap.L().Error("事务panic", zap.Any("panic", r))
		}
	}()

	// 获取兑换码信息
	var redeemCode model.RedeemCode
	if err := tx.Where("id = ?", r.RedeemCodeId).First(&redeemCode).Error; err != nil {
		tx.Rollback()
		zap.L().Error("获取兑换码失败：", zap.Error(err))
		return entity.ErrInternal
	}

	// 检查兑换码状态
	if redeemCode.IsDel == 1 {
		tx.Rollback()
		return entity.ErrRedeemCodeInvalid
	}

	// 检查兑换码是否已过期
	if time.Now().After(redeemCode.ExpirationAt) {
		tx.Rollback()
		return entity.ErrRedeemCodeExpired
	}

	// 检查兑换码是否已达到使用上限
	if redeemCode.UsedCount >= redeemCode.UsageLimit {
		tx.Rollback()
		return entity.ErrRedeemCodeUsedUp
	}

	// 创建使用记录
	redeemCodeRecord := model.RedeemCodeRecord{
		UserID:       r.UserId,
		RedeemCodeID: r.RedeemCodeId,
	}
	if err := tx.Create(&redeemCodeRecord).Error; err != nil {
		tx.Rollback()
		zap.L().Error("创建使用记录失败：", zap.Error(err))
		return entity.ErrInternal
	}

	// 更新兑换码已使用数量（乐观锁：使用 updated_at 作为版本控制）
	result := tx.Model(&model.RedeemCode{}).
		Where("id = ? AND updated_at = ? AND used_count < usage_limit", r.RedeemCodeId, redeemCode.UpdatedAt).
		Update("used_count", gorm.Expr("used_count + 1"))
	if result.Error != nil {
		tx.Rollback()
		zap.L().Error("更新兑换码使用数量失败：", zap.Error(result.Error), zap.Int64("user_id", r.UserId), zap.Int64("redeem_code_id", r.RedeemCodeId))
		return entity.ErrInternal
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		zap.L().Warn("兑换码使用失败：乐观锁失败，可能已被其他用户使用", zap.Int64("user_id", r.UserId), zap.Int64("redeem_code_id", r.RedeemCodeId))
		return entity.ErrRedeemCodeUsedUp
	}

	// 如果达到使用上限，更新批次已使用数量
	if redeemCode.UsedCount+1 >= redeemCode.UsageLimit {
		if err := tx.Model(&model.RedeemCodeBatch{}).Where("id = ?", redeemCode.RedeemCodeBatchID).Update("used_count", gorm.Expr("used_count + 1")).Error; err != nil {
			tx.Rollback()
			zap.L().Error("更新批次使用数量失败：", zap.Error(err), zap.Int64("user_id", r.UserId), zap.Int64("redeem_code_id", r.RedeemCodeId), zap.Int64("batch_id", redeemCode.RedeemCodeBatchID))
			return entity.ErrInternal
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		zap.L().Error("提交事务失败：", zap.Error(err))
		if err := tx.Rollback().Error; err != nil {
			zap.L().Error("事务回滚失败：", zap.Error(err))
		}
		return entity.ErrInternal
	}

	zap.L().Info("兑换码使用成功：", zap.Int64("user_id", r.UserId), zap.Int64("redeem_code_id", r.RedeemCodeId))
	return nil
}
