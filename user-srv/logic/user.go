package logic

import (
	"context"
	"crypto/sha512"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/1348453525/user-redeem-code-grpc/user-srv/entity"
	"github.com/1348453525/user-redeem-code-grpc/user-srv/global"
	"github.com/1348453525/user-redeem-code-grpc/user-srv/model"
	"github.com/1348453525/user-redeem-code-grpc/user-srv/pkg/helper"
	"github.com/1348453525/user-redeem-code-grpc/user-srv/pkg/jwt"
	proto "github.com/1348453525/user-redeem-code-grpc/user-srv/proto/user"
	"github.com/anaskhan96/go-password-encoder"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 密码加密选项常量
const (
	PasswordSaltLen    = 16
	PasswordIterations = 10000 // 提高迭代次数增强安全性
	PasswordKeyLen     = 32
)

// getPasswordOptions 获取密码加密选项
func getPasswordOptions() *password.Options {
	return &password.Options{
		SaltLen:      PasswordSaltLen,
		Iterations:   PasswordIterations,
		KeyLen:       PasswordKeyLen,
		HashFunction: sha512.New,
	}
}

// encodePassword 加密密码
func encodePassword(pwd string) string {
	salt, encodedPwd := password.Encode(pwd, getPasswordOptions())
	return fmt.Sprintf("pbkdf2-sha512$%s$%s", salt, encodedPwd)
}

// verifyPassword 验证密码
func verifyPassword(pwd, encrypted string) bool {
	parts := strings.Split(encrypted, "$")
	if len(parts) != 3 {
		return false
	}
	return password.Verify(pwd, parts[1], parts[2], getPasswordOptions())
}

// parseBirthday 解析生日字符串为时间指针
func parseBirthday(birthdayStr string) *time.Time {
	if birthdayStr == "" {
		return nil
	}
	if t, err := time.Parse("2006-01-02", birthdayStr); err == nil {
		return &t
	}
	return nil
}

// buildUserInfoDvo 构建用户信息
func buildUserInfoDvo(user *model.User) proto.UserInfoResponse {
	return proto.UserInfoResponse{
		Id:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Mobile:   user.Mobile,
		Gender:   user.Gender,
		Birthday: helper.FormatDate(user.Birthday),
	}
}

type UserLogic struct{}

func NewUserLogic() *UserLogic {
	return &UserLogic{}
}

func (l *UserLogic) Register(c context.Context, r *proto.RegisterRequest) (*proto.UserInfoResponse, error) {
	// 验证密码是否一致
	if r.Password != r.ConfirmPassword {
		return nil, entity.ErrPasswordNotMatch
	}

	// 验证用户名长度
	if len(r.Username) < 3 || len(r.Username) > 20 {
		return nil, entity.ErrParam
	}

	// 验证密码长度
	if len(r.Password) < 6 {
		return nil, entity.ErrParam
	}

	// 验证手机号格式
	if r.Mobile != "" && !helper.IsValidMobile(r.Mobile) {
		return nil, entity.ErrParam
	}

	// 查询用户是否存在
	var count int64
	result := global.DB.Model(&model.User{}).Where("username = ?", r.Username).Or("mobile = ?", r.Mobile).Count(&count)
	if result.Error != nil {
		zap.L().Error("查询用户失败：", zap.Error(result.Error))
		return nil, entity.ErrInternal
	}
	if count > 0 {
		return nil, entity.ErrUserExisted
	}

	// 生成密码
	pwd := encodePassword(r.Password)

	// 处理生日
	birthday := parseBirthday(r.Birthday)

	// 创建用户
	newUser := model.User{
		Username: r.Username,
		Password: pwd,
		Nickname: r.Nickname,
		Mobile:   r.Mobile,
		Gender:   r.Gender,
		Birthday: birthday,
		IsDel:    2,
	}
	result = global.DB.Create(&newUser)
	if result.Error != nil {
		return nil, entity.ErrOperationFailed
	}

	// 返回数据
	resp := buildUserInfoDvo(&newUser)
	return &resp, nil
}

func (l *UserLogic) Login(c context.Context, r *proto.LoginRequest) (*proto.LoginResponse, error) {
	// 验证用户名长度
	if len(r.Username) < 3 || len(r.Username) > 20 {
		return nil, entity.ErrParam
	}

	// 验证密码长度
	if len(r.Password) < 6 {
		return nil, entity.ErrParam
	}

	// 查询用户
	var user model.User
	result := global.DB.Where("username = ?", r.Username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, entity.ErrParam
		}
		zap.L().Error("查询用户失败：", zap.Error(result.Error))
		return nil, entity.ErrInternal
	}

	// 检验状态
	if user.IsDel == 1 {
		return nil, entity.ErrUserDisabled
	}

	// 校验密码
	if !verifyPassword(r.Password, user.Password) {
		return nil, entity.ErrPasswordError
	}

	// 生成 token
	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		zap.L().Error("生成token失败：", zap.Error(err))
		return nil, entity.ErrInternal
	}

	// 返回数据
	userInfo := buildUserInfoDvo(&user)
	resp := &proto.LoginResponse{
		Info:  &userInfo,
		Token: token,
	}
	return resp, nil
}

func (l *UserLogic) Info(c context.Context, id int64) (*proto.UserInfoResponse, error) {
	// 查询用户
	var user model.User
	if err := user.GetByID(id); err != nil {
		return nil, err
	}

	// 返回数据
	resp := buildUserInfoDvo(&user)
	return &resp, nil
}

func (l *UserLogic) GetList(c context.Context, r *proto.GetUserListRequest) (*proto.GetUserListResponse, error) {
	var user model.User
	list, count := user.GetList(r.Page, r.PageSize)
	resp := &proto.GetUserListResponse{
		Page:     r.Page,
		PageSize: r.PageSize,
		Total:    count,
	}
	for _, v := range list {
		userInfoDvo := buildUserInfoDvo(v)
		resp.Data = append(resp.Data, &userInfoDvo)
	}
	return resp, nil
}

func (l *UserLogic) Update(c context.Context, r *proto.UpdateUserRequest) error {
	// 查询用户是否存在
	var count int64
	result := global.DB.Model(&model.User{}).Where("id = ?", r.Id).Count(&count)
	if result.Error != nil {
		zap.L().Error("查询用户失败：", zap.Error(result.Error))
		return entity.ErrInternal
	}
	if count == 0 {
		return entity.ErrParam
	}

	// 处理生日
	birthday := parseBirthday(r.Birthday)

	// 构建更新数据
	updateData := map[string]interface{}{
		"username": r.Username,
		"nickname": r.Nickname,
		"mobile":   r.Mobile,
		"gender":   r.Gender,
		"birthday": birthday,
	}
	result = global.DB.Model(&model.User{}).Where("id=?", r.Id).Updates(updateData)
	if result.Error != nil {
		zap.L().Error("更新用户失败：", zap.Error(result.Error))
		return entity.ErrInternal
	}

	// 检查是否更新成功
	if result.RowsAffected == 0 {
		return entity.ErrOperationFailed
	}
	return nil
}

func (l *UserLogic) Delete(c context.Context, id int64) error {
	// 删除用户
	result := global.DB.Model(&model.User{}).Where("id=?", id).Update("is_del", 1)
	if result.Error != nil {
		zap.L().Error("删除用户失败：", zap.Error(result.Error))
		return entity.ErrInternal
	}

	// 检查是否删除成功
	if result.RowsAffected == 0 {
		return entity.ErrOperationFailed
	}
	return nil
}
