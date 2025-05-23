// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: usercenter.proto

package usercenter

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/userpb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BatchGetUsersInfoRequest  = userpb.BatchGetUsersInfoRequest
	BatchGetUsersInfoResponse = userpb.BatchGetUsersInfoResponse
	GenerateTokenReq          = userpb.GenerateTokenReq
	GenerateTokenResp         = userpb.GenerateTokenResp
	GetUserInfoRequest        = userpb.GetUserInfoRequest
	GetUserInfoResponse       = userpb.GetUserInfoResponse
	GetUserRoleRequest        = userpb.GetUserRoleRequest
	GetUserRoleResponse       = userpb.GetUserRoleResponse
	LoginRequest              = userpb.LoginRequest
	LoginResponse             = userpb.LoginResponse
	RegisterRequest           = userpb.RegisterRequest
	RegisterResponse          = userpb.RegisterResponse
	UpdateNicknameRequest     = userpb.UpdateNicknameRequest
	UpdateNicknameResponse    = userpb.UpdateNicknameResponse
	UpdatePasswordRequest     = userpb.UpdatePasswordRequest
	UpdatePasswordResponse    = userpb.UpdatePasswordResponse
	User                      = userpb.User
	UserInfo                  = userpb.UserInfo

	Usercenter interface {
		// 用户登录
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		// 用户注册
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		// 修改用户名
		UpdateNickname(ctx context.Context, in *UpdateNicknameRequest, opts ...grpc.CallOption) (*UpdateNicknameResponse, error)
		// 修改密码
		UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error)
		// 生成 token
		GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
		GetUserRole(ctx context.Context, in *GetUserRoleRequest, opts ...grpc.CallOption) (*GetUserRoleResponse, error)
		GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
		GetUsersInfo(ctx context.Context, in *BatchGetUsersInfoRequest, opts ...grpc.CallOption) (*BatchGetUsersInfoResponse, error)
	}

	defaultUsercenter struct {
		cli zrpc.Client
	}
)

func NewUsercenter(cli zrpc.Client) Usercenter {
	return &defaultUsercenter{
		cli: cli,
	}
}

// 用户登录
func (m *defaultUsercenter) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := userpb.NewUsercenterClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

// 用户注册
func (m *defaultUsercenter) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := userpb.NewUsercenterClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

// 修改用户名
func (m *defaultUsercenter) UpdateNickname(ctx context.Context, in *UpdateNicknameRequest, opts ...grpc.CallOption) (*UpdateNicknameResponse, error) {
	client := userpb.NewUsercenterClient(m.cli.Conn())
	return client.UpdateNickname(ctx, in, opts...)
}

// 修改密码
func (m *defaultUsercenter) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error) {
	client := userpb.NewUsercenterClient(m.cli.Conn())
	return client.UpdatePassword(ctx, in, opts...)
}

// 生成 token
func (m *defaultUsercenter) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	client := userpb.NewUsercenterClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserRole(ctx context.Context, in *GetUserRoleRequest, opts ...grpc.CallOption) (*GetUserRoleResponse, error) {
	client := userpb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserRole(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	client := userpb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUsersInfo(ctx context.Context, in *BatchGetUsersInfoRequest, opts ...grpc.CallOption) (*BatchGetUsersInfoResponse, error) {
	client := userpb.NewUsercenterClient(m.cli.Conn())
	return client.GetUsersInfo(ctx, in, opts...)
}
