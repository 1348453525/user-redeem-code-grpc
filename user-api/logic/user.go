package logic

import (
	"github.com/1348453525/user-redeem-code-grpc/user-api/entity"
	"github.com/1348453525/user-redeem-code-grpc/user-api/global"
	proto "github.com/1348453525/user-redeem-code-grpc/user-api/proto/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserLogic struct{}

func NewUserLogic() *UserLogic {
	return &UserLogic{}
}

func (l *UserLogic) Register(c *gin.Context, r *entity.RegisterDto) (*proto.UserInfoResponse, error) {
	// 调用 grpc 服务
	rpcResp, err := global.UserClient.Register(c, &proto.RegisterRequest{
		Username:        r.Username,
		Password:        r.Password,
		ConfirmPassword: r.ConfirmPassword,
		Nickname:        r.Nickname,
		Mobile:          r.Mobile,
		Gender:          r.Gender,
		Birthday:        r.Birthday,
	})
	if err != nil {
		zap.L().Error("调用 grpc 服务失败：", zap.Error(err))
		return nil, entity.ErrInternal
	}

	// 返回数据
	return rpcResp, nil
}

func (l *UserLogic) Login(c *gin.Context, r *entity.LoginDto) (*proto.LoginResponse, error) {
	// 调用 grpc 服务
	rpcResp, err := global.UserClient.Login(c, &proto.LoginRequest{
		Username: r.Username,
		Password: r.Password,
	})
	if err != nil {
		zap.L().Error("调用 grpc 服务失败：", zap.Error(err))
		return nil, entity.ErrInternal
	}

	// 返回数据
	return rpcResp, nil
}

func (l *UserLogic) Info(c *gin.Context, id int64) (*proto.UserInfoResponse, error) {
	// 调用 grpc 服务
	rpcResp, err := global.UserClient.Info(c, &proto.IDRequest{
		Id: id,
	})
	if err != nil {
		zap.L().Error("调用 grpc 服务失败：", zap.Error(err))
		return nil, entity.ErrInternal
	}

	// 返回数据
	return rpcResp, nil
}

func (l *UserLogic) GetList(c *gin.Context, r *entity.GetUserListDto) (*proto.GetUserListResponse, error) {
	// 调用 grpc 服务
	rpcResp, err := global.UserClient.GetList(c, &proto.GetUserListRequest{
		Page:     r.Page,
		PageSize: r.PageSize,
	})
	if err != nil {
		zap.L().Error("调用 grpc 服务失败：", zap.Error(err))
		return nil, entity.ErrInternal
	}

	return rpcResp, nil
}

func (l *UserLogic) Update(c *gin.Context, r *entity.UpdateUserDto) error {
	// 调用 grpc 服务
	_, err := global.UserClient.Update(c, &proto.UpdateUserRequest{
		Id:       r.ID,
		Username: r.Username,
		Nickname: r.Nickname,
		Mobile:   r.Mobile,
		Gender:   r.Gender,
		Birthday: r.Birthday,
	})
	if err != nil {
		zap.L().Error("调用 grpc 服务失败：", zap.Error(err))
		return entity.ErrInternal
	}

	return nil
}

func (l *UserLogic) Delete(c *gin.Context, id int64) error {
	// 调用 grpc 服务
	_, err := global.UserClient.Delete(c, &proto.IDRequest{
		Id: id,
	})
	if err != nil {
		zap.L().Error("调用 grpc 服务失败：", zap.Error(err))
		return entity.ErrInternal
	}
	return nil
}
