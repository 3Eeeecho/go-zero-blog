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

type GetUsersInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUsersInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersInfoLogic {
	return &GetUsersInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUsersInfoLogic) GetUsersInfo(in *userpb.BatchGetUsersInfoRequest) (*userpb.BatchGetUsersInfoResponse, error) {
	keys := make([]string, len(in.Ids))
	for i, id := range in.Ids {
		keys[i] = fmt.Sprintf("user:%d", id)
	}

	// 批量查 Redis
	cached, err := l.svcCtx.Redis.Mget(keys...)
	if err != nil {
		l.Logger.Errorf("cache hit failed,err: %v", err)
	}

	results := make([]*userpb.UserInfo, len(in.Ids))
	idMap := buildIndexMap(in.Ids)
	var missIDs []int64
	for i, val := range cached {
		if val != "" {
			var u userpb.UserInfo
			json.Unmarshal([]byte(val), &u)
			results[i] = &u
		} else {
			missIDs = append(missIDs, in.Ids[i])
		}
	}

	// 未命中的批量查数据库
	if len(missIDs) > 0 {
		users, err := l.svcCtx.UserModel.FindByUserIds(l.ctx, missIDs)
		if err != nil {
			return nil, xerr.NewErrCode(xerr.DB_ERROR)
		}
		for _, u := range users {
			userInfo := &userpb.UserInfo{Id: u.Id, Nickname: u.Nickname, Role: u.Role}
			if idx, ok := idMap[u.Id]; ok {
				results[idx] = userInfo
			}
			// 回写缓存
			data, err := json.Marshal(userInfo)
			if err != nil {
				l.Logger.Errorf("marshal user failed, id: %d, err: %v", u.Id, err)
				continue
			}
			if err := l.svcCtx.Redis.Set(fmt.Sprintf("user:%d", u.Id), string(data)); err != nil {
				l.Logger.Errorf("redis set failed, id: %d, err: %v", u.Id, err)
			}
		}
	}

	return &userpb.BatchGetUsersInfoResponse{Users: results}, nil
}

func buildIndexMap(ids []int64) map[int64]int {
	m := make(map[int64]int, len(ids))
	for i, id := range ids {
		m[id] = i
	}
	return m
}
