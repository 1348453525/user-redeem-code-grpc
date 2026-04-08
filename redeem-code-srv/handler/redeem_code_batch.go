package handler

import (
	"context"

	"github.com/1348453525/user-redeem-code-grpc/redeem-code-srv/entity"
	"github.com/1348453525/user-redeem-code-grpc/redeem-code-srv/logic"
	proto "github.com/1348453525/user-redeem-code-grpc/redeem-code-srv/proto/redeem_code"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *RedeemCode) CreateRedeemCodeBatch(c context.Context, r *proto.CreateRedeemCodeBatchRequest) (*proto.RedeemCodeBatchResponse, error) {
	// 处理逻辑
	resp, err := logic.NewRedeemCodeBatchLogic().Create(c, r)
	if err != nil {
		return nil, entity.ToGrpcError(err)
	}
	return resp, nil
}

func (h *RedeemCode) RedeemCodeBatchDetail(c context.Context, r *proto.IDRequest) (*proto.RedeemCodeBatchResponse, error) {
	// 处理逻辑
	resp, err := logic.NewRedeemCodeBatchLogic().Detail(c, r)
	if err != nil {
		return nil, entity.ToGrpcError(err)
	}
	return resp, nil
}

func (h *RedeemCode) GetRedeemCodeBatchList(c context.Context, r *proto.GetListRequest) (*proto.GetRedeemCodeBatchListResponse, error) {
	// 处理逻辑
	resp, err := logic.NewRedeemCodeBatchLogic().GetList(c, r)
	if err != nil {
		return nil, entity.ToGrpcError(err)
	}
	return resp, nil
}

func (h *RedeemCode) UpdateRedeemCodeBatch(c context.Context, r *proto.UpdateRedeemCodeBatchRequest) (*emptypb.Empty, error) {
	// 处理逻辑
	err := logic.NewRedeemCodeBatchLogic().Update(c, r)
	if err != nil {
		return &emptypb.Empty{}, entity.ToGrpcError(err)
	}
	return &emptypb.Empty{}, nil
}

func (h *RedeemCode) DeleteRedeemCodeBatch(c context.Context, r *proto.IDRequest) (*emptypb.Empty, error) {
	// 处理逻辑
	err := logic.NewRedeemCodeBatchLogic().Delete(c, r)
	if err != nil {
		return &emptypb.Empty{}, entity.ToGrpcError(err)
	}
	return &emptypb.Empty{}, nil
}
