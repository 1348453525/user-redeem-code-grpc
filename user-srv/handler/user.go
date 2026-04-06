package handler

import (
	"context"

	"github.com/1348453525/user-redeem-code-grpc/user-srv/entity"
	"github.com/1348453525/user-redeem-code-grpc/user-srv/logic"
	proto "github.com/1348453525/user-redeem-code-grpc/user-srv/proto/user"
	"google.golang.org/protobuf/types/known/emptypb"
)

type User struct {
	proto.UnimplementedUserServer
}

func (h *User) Register(c context.Context, r *proto.RegisterRequest) (*proto.UserInfoResponse, error) {
	// 处理逻辑
	resp, err := logic.NewUserLogic().Register(c, r)
	if err != nil {
		return nil, entity.ToGrpcError(err)
	}
	return resp, nil
}

func (h *User) Login(c context.Context, r *proto.LoginRequest) (*proto.LoginResponse, error) {
	// 处理逻辑
	resp, err := logic.NewUserLogic().Login(c, r)
	if err != nil {
		return nil, entity.ToGrpcError(err)
	}
	return resp, nil
}

func (h *User) Info(c context.Context, r *proto.IDRequest) (*proto.UserInfoResponse, error) {
	// 处理逻辑
	resp, err := logic.NewUserLogic().Info(c, r.Id)
	if err != nil {
		return nil, entity.ToGrpcError(err)
	}
	return resp, nil
}

func (h *User) GetList(c context.Context, r *proto.GetUserListRequest) (*proto.GetUserListResponse, error) {
	// 处理逻辑
	resp, err := logic.NewUserLogic().GetList(c, r)
	if err != nil {
		return nil, entity.ToGrpcError(err)
	}
	return resp, nil
}

func (h *User) Update(c context.Context, r *proto.UpdateUserRequest) (*emptypb.Empty, error) {
	// 处理逻辑
	err := logic.NewUserLogic().Update(c, r)
	if err != nil {
		return &emptypb.Empty{}, entity.ToGrpcError(err)
	}
	return &emptypb.Empty{}, nil
}

func (h *User) Delete(c context.Context, r *proto.IDRequest) (*emptypb.Empty, error) {
	// 处理逻辑
	err := logic.NewUserLogic().Delete(c, r.Id)
	if err != nil {
		return &emptypb.Empty{}, entity.ToGrpcError(err)
	}
	return &emptypb.Empty{}, nil
}
