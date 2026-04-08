package logic

import (
	"github.com/1348453525/user-redeem-code-grpc/user-api/entity"
	"github.com/1348453525/user-redeem-code-grpc/user-api/global"
	proto "github.com/1348453525/user-redeem-code-grpc/user-api/proto/redeem_code"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RedeemCodeLogic struct{}

func NewRedeemCodeLogic() *RedeemCodeLogic {
	return &RedeemCodeLogic{}
}

func (l *RedeemCodeLogic) Detail(c *gin.Context, id int64) (*proto.RedeemCodeResponse, error) {
	// 调用 grpc 服务
	rpcResp, err := global.RedeemCodeClient.RedeemCodeDetail(c, &proto.IDRequest{
		Id: id,
	})
	if err != nil {
		zap.L().Error("调用 grpc 服务失败：", zap.Error(err))
		return nil, entity.ErrInternal
	}
	// 返回数据
	return rpcResp, nil
}

func (l *RedeemCodeLogic) GetList(c *gin.Context, r *entity.GetRedeemCodeListDto) (*proto.GetRedeemCodeListResponse, error) {
	// 调用 grpc 服务
	rpcResp, err := global.RedeemCodeClient.GetRedeemCodeList(c, &proto.GetListRequest{
		Page:     r.Page,
		PageSize: r.PageSize,
	})
	if err != nil {
		zap.L().Error("调用 grpc 服务失败：", zap.Error(err))
		return nil, entity.ErrInternal
	}

	return rpcResp, nil
}

func (l *RedeemCodeLogic) Update(c *gin.Context, r *entity.UpdateRedeemCodeDto) error {
	// 调用 grpc 服务
	_, err := global.RedeemCodeClient.UpdateRedeemCode(c, &proto.UpdateRedeemCodeRequest{
		Id:           r.ID,
		Title:        r.Title,
		ExpirationAt: r.ExpirationAt,
		IsDel:        r.IsDel,
	})
	if err != nil {
		zap.L().Error("调用 grpc 服务失败：", zap.Error(err))
		return entity.ErrInternal
	}

	return nil
}

func (l *RedeemCodeLogic) Delete(c *gin.Context, id int64) error {
	// 调用 grpc 服务
	_, err := global.RedeemCodeClient.DeleteRedeemCode(c, &proto.IDRequest{
		Id: id,
	})
	if err != nil {
		zap.L().Error("调用 grpc 服务失败：", zap.Error(err))
		return entity.ErrInternal
	}

	return nil
}

func (l *RedeemCodeLogic) Use(c *gin.Context, r *entity.UseRedeemCodeDto) error {
	// 调用 grpc 服务
	_, err := global.RedeemCodeClient.UseRedeemCode(c, &proto.UseRedeemCodeRequest{
		RedeemCodeId: r.RedeemCodeID,
		UserId:       r.UserID,
	})
	if err != nil {
		zap.L().Error("调用 grpc 服务失败：", zap.Error(err))
		return entity.ErrInternal
	}

	return nil
}
