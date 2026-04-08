package logic

import (
	"github.com/1348453525/user-redeem-code-grpc/user-api/entity"
	"github.com/1348453525/user-redeem-code-grpc/user-api/global"
	proto "github.com/1348453525/user-redeem-code-grpc/user-api/proto/redeem_code"
	uproto "github.com/1348453525/user-redeem-code-grpc/user-api/proto/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RedeemCodeBatchLogic struct{}

func NewRedeemCodeBatchLogic() *RedeemCodeBatchLogic {
	return &RedeemCodeBatchLogic{}
}

func (l *RedeemCodeBatchLogic) Create(c *gin.Context, userID int64, r *entity.CreateRedeemCodeBatchDto) (*proto.RedeemCodeBatchResponse, error) {
	// 获取用户信息
	userInfo, err := global.UserClient.Info(c, &uproto.IDRequest{Id: userID})
	if err != nil {
		zap.L().Error("获取用户信息失败：", zap.Error(err), zap.Int64("userID", userID))
		return nil, entity.ErrInternal
	}

	// 调用 grpc 服务
	rpcResp, err := global.RedeemCodeClient.CreateRedeemCodeBatch(c, &proto.CreateRedeemCodeBatchRequest{
		Title:       r.Title,
		Description: r.Description,
		StartedAt:   r.StartedAt,
		EndedAt:     r.EndedAt,
		Status:      r.Status,
		TotalCount:  r.TotalCount,
		UsageLimit:  r.UsageLimit,
		CreatorId:   userInfo.Id,
		CreatorName: userInfo.Nickname,
	})
	if err != nil {
		zap.L().Error("调用 grpc 服务失败：", zap.Error(err))
		return nil, entity.ErrInternal
	}
	return rpcResp, nil
}

func (l *RedeemCodeBatchLogic) Detail(c *gin.Context, id int64) (*proto.RedeemCodeBatchResponse, error) {
	// 调用 grpc 服务
	rpcResp, err := global.RedeemCodeClient.RedeemCodeBatchDetail(c, &proto.IDRequest{
		Id: id,
	})
	if err != nil {
		zap.L().Error("调用 grpc 服务失败：", zap.Error(err))
		return nil, entity.ErrInternal
	}
	return rpcResp, nil
}

func (l *RedeemCodeBatchLogic) GetList(c *gin.Context, r *entity.GetRedeemCodeBatchListDto) (*proto.GetRedeemCodeBatchListResponse, error) {
	// 调用 grpc 服务
	rpcResp, err := global.RedeemCodeClient.GetRedeemCodeBatchList(c, &proto.GetListRequest{
		Page:     r.Page,
		PageSize: r.PageSize,
	})
	if err != nil {
		zap.L().Error("调用 grpc 服务失败：", zap.Error(err))
		return nil, entity.ErrInternal
	}

	return rpcResp, nil
}

func (l *RedeemCodeBatchLogic) Update(c *gin.Context, r *entity.UpdateRedeemCodeBatchDto) error {
	// 调用 grpc 服务
	_, err := global.RedeemCodeClient.UpdateRedeemCodeBatch(c, &proto.UpdateRedeemCodeBatchRequest{
		Id:          r.ID,
		Title:       r.Title,
		Description: r.Description,
		StartedAt:   r.StartedAt,
		EndedAt:     r.EndedAt,
		Status:      r.Status,
	})
	if err != nil {
		zap.L().Error("调用 grpc 服务失败：", zap.Error(err))
		return entity.ErrInternal
	}

	return nil
}

func (l *RedeemCodeBatchLogic) Delete(c *gin.Context, id int64) error {
	// 调用 grpc 服务
	_, err := global.RedeemCodeClient.DeleteRedeemCodeBatch(c, &proto.IDRequest{
		Id: id,
	})
	if err != nil {
		zap.L().Error("调用 grpc 服务失败：", zap.Error(err))
		return entity.ErrInternal
	}

	return nil
}
