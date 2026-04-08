package handler

import (
	"context"

	"github.com/1348453525/user-redeem-code-grpc/redeem-code-srv/entity"
	"github.com/1348453525/user-redeem-code-grpc/redeem-code-srv/logic"
	proto "github.com/1348453525/user-redeem-code-grpc/redeem-code-srv/proto/redeem_code"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RedeemCode struct {
	proto.UnimplementedRedeemCodeServer
}

func (h *RedeemCode) RedeemCodeDetail(c context.Context, r *proto.IDRequest) (*proto.RedeemCodeResponse, error) {
	// 处理逻辑
	resp, err := logic.NewRedeemCodeLogic().Detail(c, r)
	if err != nil {
		return nil, entity.ToGrpcError(err)
	}
	return resp, nil
}
func (h *RedeemCode) GetRedeemCodeList(c context.Context, r *proto.GetListRequest) (*proto.GetRedeemCodeListResponse, error) {
	// 处理逻辑
	resp, err := logic.NewRedeemCodeLogic().GetList(c, r)
	if err != nil {
		return nil, entity.ToGrpcError(err)
	}
	return resp, nil
}
func (h *RedeemCode) UpdateRedeemCode(c context.Context, r *proto.UpdateRedeemCodeRequest) (*emptypb.Empty, error) {
	// 处理逻辑
	err := logic.NewRedeemCodeLogic().Update(c, r)
	if err != nil {
		return &emptypb.Empty{}, entity.ToGrpcError(err)
	}
	return &emptypb.Empty{}, nil
}
func (h *RedeemCode) DeleteRedeemCode(c context.Context, r *proto.IDRequest) (*emptypb.Empty, error) {
	// 处理逻辑
	err := logic.NewRedeemCodeLogic().Delete(c, r)
	if err != nil {
		return &emptypb.Empty{}, entity.ToGrpcError(err)
	}
	return &emptypb.Empty{}, nil
}
func (h *RedeemCode) UseRedeemCode(c context.Context, r *proto.UseRedeemCodeRequest) (*emptypb.Empty, error) {
	// 处理逻辑
	err := logic.NewRedeemCodeLogic().Use(c, r)
	if err != nil {
		return &emptypb.Empty{}, entity.ToGrpcError(err)
	}
	return &emptypb.Empty{}, nil
}
