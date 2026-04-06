package entity

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInternal = errors.New("内部错误")

	ErrParam           = errors.New("参数错误")
	ErrOperationFailed = errors.New("操作失败")
)

func ToGrpcError(err error) error {
	if errors.Is(err, ErrInternal) {
		return status.Error(codes.Internal, err.Error())
	} else if errors.Is(err, ErrParam) {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return status.Error(codes.Unknown, err.Error())
}
