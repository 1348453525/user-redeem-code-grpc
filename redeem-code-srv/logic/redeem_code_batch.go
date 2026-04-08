package logic

import (
	"context"

	"github.com/1348453525/user-redeem-code-grpc/redeem-code-srv/entity"
	"github.com/1348453525/user-redeem-code-grpc/redeem-code-srv/global"
	"github.com/1348453525/user-redeem-code-grpc/redeem-code-srv/model"
	"github.com/1348453525/user-redeem-code-grpc/redeem-code-srv/pkg/helper"
	proto "github.com/1348453525/user-redeem-code-grpc/redeem-code-srv/proto/redeem_code"
	"github.com/gofrs/uuid/v5"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RedeemCodeBatchLogic struct{}

func NewRedeemCodeBatchLogic() *RedeemCodeBatchLogic {
	return &RedeemCodeBatchLogic{}
}

func (l *RedeemCodeBatchLogic) Create(c context.Context, r *proto.CreateRedeemCodeBatchRequest) (*proto.RedeemCodeBatchResponse, error) {
	// 参数验证
	if r.Title == "" || r.TotalCount <= 0 || r.UsageLimit <= 0 {
		return nil, entity.ErrParam
	}

	// 解析日期
	startedAt, err := helper.ParseDatetime(r.StartedAt)
	if err != nil {
		zap.L().Error("解析开始时间失败：", zap.Error(err), zap.String("startedAt", r.StartedAt))
		return nil, entity.ErrParam
	}
	endedAt, err := helper.ParseDatetime(r.EndedAt)
	if err != nil {
		zap.L().Error("解析结束时间失败：", zap.Error(err), zap.String("endedAt", r.EndedAt))
		return nil, entity.ErrParam
	}

	// 验证时间范围
	if endedAt.Before(*startedAt) {
		zap.L().Error("结束时间早于开始时间", zap.String("startedAt", r.StartedAt), zap.String("endedAt", r.EndedAt))
		return nil, entity.ErrParam
	}

	// 使用事务确保原子性
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			zap.L().Error("事务panic", zap.Any("panic", r))
		}
	}()

	// 创建兑换码批次
	newRedeemCodeBatch := model.RedeemCodeBatch{
		Title:       r.Title,
		Description: r.Description,
		UsageLimit:  r.UsageLimit,
		TotalCount:  r.TotalCount,
		StartedAt:   *startedAt,
		EndedAt:     *endedAt,
		Status:      1,
		CreatorID:   r.CreatorId,
		CreatorName: r.CreatorName,
	}
	if result := tx.Create(&newRedeemCodeBatch); result.Error != nil {
		tx.Rollback()
		zap.L().Error("创建兑换码批次失败：", zap.Error(result.Error))
		return nil, entity.ErrOperationFailed
	}

	// 批量创建兑换码（分批次处理，避免内存问题）
	const batchSize = 1000
	totalCount := int(r.TotalCount)
	for i := 0; i < totalCount; i += batchSize {
		end := i + batchSize
		if end > totalCount {
			end = totalCount
		}

		var newRedeemCodes []model.RedeemCode
		for j := i; j < end; j++ {
			redeemCode := model.RedeemCode{
				RedeemCodeBatchID: newRedeemCodeBatch.ID,
				Title:             r.Title,
				Value:             uuid.Must(uuid.NewV7()).String(),
				UsageLimit:        r.UsageLimit,
				ExpirationAt:      *endedAt,
				IsDel:             2,
			}
			newRedeemCodes = append(newRedeemCodes, redeemCode)
		}

		if result := tx.Create(&newRedeemCodes); result.Error != nil {
			tx.Rollback()
			zap.L().Error("批量创建兑换码失败：", zap.Error(result.Error))
			return nil, entity.ErrOperationFailed
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		zap.L().Error("提交事务失败：", zap.Error(err))
		return nil, entity.ErrOperationFailed
	}

	// 返回数据
	return &proto.RedeemCodeBatchResponse{
		Id:          newRedeemCodeBatch.ID,
		Title:       newRedeemCodeBatch.Title,
		Description: newRedeemCodeBatch.Description,
		UsageLimit:  newRedeemCodeBatch.UsageLimit,
		TotalCount:  newRedeemCodeBatch.TotalCount,
		UsedCount:   newRedeemCodeBatch.UsedCount,
		StartedAt:   timestamppb.New(newRedeemCodeBatch.StartedAt),
		EndedAt:     timestamppb.New(newRedeemCodeBatch.EndedAt),
		Status:      newRedeemCodeBatch.Status,
		CreatorId:   newRedeemCodeBatch.CreatorID,
		CreatorName: newRedeemCodeBatch.CreatorName,
		CreatedAt:   timestamppb.New(newRedeemCodeBatch.CreatedAt),
		UpdatedAt:   timestamppb.New(newRedeemCodeBatch.UpdatedAt),
	}, nil
}

func (l *RedeemCodeBatchLogic) Detail(c context.Context, r *proto.IDRequest) (*proto.RedeemCodeBatchResponse, error) {
	// 参数验证
	if r.Id <= 0 {
		return nil, entity.ErrParam
	}

	var redeemCodeBatch model.RedeemCodeBatch
	if err := redeemCodeBatch.GetByID(r.Id); err != nil {
		zap.L().Error("获取兑换码批次详情失败：", zap.Error(err), zap.Int64("id", r.Id))
		return nil, entity.ErrInternal
	}
	// 返回数据
	return &proto.RedeemCodeBatchResponse{
		Id:          redeemCodeBatch.ID,
		Title:       redeemCodeBatch.Title,
		Description: redeemCodeBatch.Description,
		UsageLimit:  redeemCodeBatch.UsageLimit,
		TotalCount:  redeemCodeBatch.TotalCount,
		UsedCount:   redeemCodeBatch.UsedCount,
		StartedAt:   timestamppb.New(redeemCodeBatch.StartedAt),
		EndedAt:     timestamppb.New(redeemCodeBatch.EndedAt),
		Status:      redeemCodeBatch.Status,
		CreatorId:   redeemCodeBatch.CreatorID,
		CreatorName: redeemCodeBatch.CreatorName,
		CreatedAt:   timestamppb.New(redeemCodeBatch.CreatedAt),
		UpdatedAt:   timestamppb.New(redeemCodeBatch.UpdatedAt),
	}, nil
}

func (l *RedeemCodeBatchLogic) GetList(c context.Context, r *proto.GetListRequest) (*proto.GetRedeemCodeBatchListResponse, error) {
	// 验证分页参数
	if r.Page <= 0 {
		r.Page = 1
	}
	if r.PageSize <= 0 || r.PageSize > 100 {
		r.PageSize = 10
	}

	var redeemCodeBatch model.RedeemCodeBatch
	list, count := redeemCodeBatch.GetList(r.Page, r.PageSize)
	resp := &proto.GetRedeemCodeBatchListResponse{
		Page:     r.Page,
		PageSize: r.PageSize,
		Total:    count,
	}
	for _, v := range list {
		redeemCodeBatchTmp := &proto.RedeemCodeBatchResponse{
			Id:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			UsageLimit:  v.UsageLimit,
			TotalCount:  v.TotalCount,
			UsedCount:   v.UsedCount,
			StartedAt:   timestamppb.New(v.StartedAt),
			EndedAt:     timestamppb.New(v.EndedAt),
			Status:      v.Status,
			CreatorId:   v.CreatorID,
			CreatorName: v.CreatorName,
			CreatedAt:   timestamppb.New(v.CreatedAt),
			UpdatedAt:   timestamppb.New(v.UpdatedAt),
		}
		resp.Data = append(resp.Data, redeemCodeBatchTmp)
	}
	return resp, nil
}

func (l *RedeemCodeBatchLogic) Update(c context.Context, r *proto.UpdateRedeemCodeBatchRequest) error {
	// 参数验证
	if r.Id <= 0 || r.Title == "" {
		return entity.ErrParam
	}

	// 解析日期
	startedAt, err := helper.ParseDatetime(r.StartedAt)
	if err != nil {
		zap.L().Error("解析开始时间失败：", zap.Error(err), zap.String("startedAt", r.StartedAt))
		return entity.ErrParam
	}
	endedAt, err := helper.ParseDatetime(r.EndedAt)
	if err != nil {
		zap.L().Error("解析结束时间失败：", zap.Error(err), zap.String("endedAt", r.EndedAt))
		return entity.ErrParam
	}

	// 验证时间范围
	if endedAt.Before(*startedAt) {
		zap.L().Error("结束时间早于开始时间", zap.String("startedAt", r.StartedAt), zap.String("endedAt", r.EndedAt))
		return entity.ErrParam
	}

	// 验证批次是否存在
	var existingBatch model.RedeemCodeBatch
	if err := existingBatch.GetByID(r.Id); err != nil {
		zap.L().Error("获取兑换码批次失败：", zap.Error(err), zap.Int64("id", r.Id))
		return entity.ErrInternal
	}

	redeemCodeBatch := model.RedeemCodeBatch{
		ID:          r.Id,
		Title:       r.Title,
		Description: r.Description,
		// UsageLimit:  r.UsageLimit, // 不允许修改使用限制
		// TotalCount:  r.TotalCount, // 不允许修改总数
		StartedAt: *startedAt,
		EndedAt:   *endedAt,
		Status:    r.Status,
	}

	// 更新批次信息
	result := global.DB.Model(&model.RedeemCodeBatch{}).Where("id=?", r.Id).Updates(&redeemCodeBatch)
	if result.Error != nil {
		zap.L().Error("更新兑换码批次失败：", zap.Error(result.Error), zap.Int64("id", r.Id))
		return entity.ErrInternal
	}

	// 如果结束时间发生变化，更新关联的兑换码过期时间
	if existingBatch.EndedAt != *endedAt {
		if err := global.DB.Model(&model.RedeemCode{}).Where("redeem_code_batch_id=?", r.Id).Update("expiration_at", *endedAt).Error; err != nil {
			zap.L().Error("更新兑换码过期时间失败：", zap.Error(err), zap.Int64("batchID", r.Id))
		}
	}

	return nil
}

func (l *RedeemCodeBatchLogic) Delete(c context.Context, r *proto.IDRequest) error {
	// 参数验证
	if r.Id <= 0 {
		return entity.ErrParam
	}

	// 验证批次是否存在
	var existingBatch model.RedeemCodeBatch
	if err := existingBatch.GetByID(r.Id); err != nil {
		zap.L().Error("获取兑换码批次失败：", zap.Error(err), zap.Int64("id", r.Id))
		return entity.ErrInternal
	}

	// 使用事务确保原子性
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			zap.L().Error("事务panic", zap.Any("panic", r))
		}
	}()

	// 更新批次状态为删除
	if err := tx.Model(&model.RedeemCodeBatch{}).Where("id=?", r.Id).Update("status", 2).Error; err != nil {
		tx.Rollback()
		zap.L().Error("删除兑换码批次失败：", zap.Error(err), zap.Int64("id", r.Id))
		return entity.ErrInternal
	}

	// 更新关联的兑换码状态
	if err := tx.Model(&model.RedeemCode{}).Where("redeem_code_batch_id=?", r.Id).Update("is_del", 1).Error; err != nil {
		tx.Rollback()
		zap.L().Error("更新兑换码状态失败：", zap.Error(err), zap.Int64("batchID", r.Id))
		return entity.ErrInternal
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		zap.L().Error("提交事务失败：", zap.Error(err))
		return entity.ErrOperationFailed
	}

	return nil
}
