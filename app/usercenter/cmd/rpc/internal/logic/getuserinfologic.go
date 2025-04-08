package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/userpb"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *userpb.GetUserInfoRequest) (*userpb.GetUserInfoResponse, error) {
	//尝试redis缓存获取用户信息
	cachedKey := fmt.Sprintf("user:%d", in.Id)
	cached, err := l.svcCtx.Redis.Get(cachedKey)
	if err != nil {
		l.Logger.Errorf("cache hit userid failed")
	}
	if cached != "" {
		var user userpb.UserInfo
		if err := json.Unmarshal([]byte(cached), &user); err != nil {
			l.Logger.Errorf("unmarshal cache failed, id: %d, err: %v", in.Id, err)
			// 缓存失效，继续查数据库
		} else {
			return &userpb.GetUserInfoResponse{User: &user}, nil
		}
	}

	//缓存未命中,从数据库获取
	user, err := l.svcCtx.UserModel.FindByUserId(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("find user failed, id: %d, err: %v", in.Id, err)
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	if user == nil {
		l.Logger.Error("user not found, id: %d", in.Id)
		return nil, xerr.NewErrCode(xerr.USER_NOT_FOUND)
	}

	// 构造响应
	userInfo := &userpb.UserInfo{
		Id:       user.Id,
		Nickname: user.Nickname,
		Role:     user.Role,
	}

	// 回写缓存
	data, err := json.Marshal(userInfo)
	if err != nil {
		l.Logger.Errorf("marshal user failed, id: %d, err: %v", in.Id, err)
	} else if err := l.svcCtx.Redis.Set(cachedKey, string(data)); err != nil {
		l.Logger.Errorf("redis set failed, id: %d, err: %v", in.Id, err)
	}

	return &userpb.GetUserInfoResponse{User: userInfo}, nil
}
