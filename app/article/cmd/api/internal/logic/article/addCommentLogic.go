package article

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/articleservice"
	"github.com/3Eeeecho/go-zero-blog/pkg/ctxdata"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCommentLogic) AddComment(req *types.CommentReq) (resp *types.CommentResp, err error) {
	if req.ArticleId < 0 || req.ParentId < 0 {
		l.Logger.Errorf("invalid param req.Id: %d ,err: %v", req.ArticleId, err)
		return nil, xerr.NewErrCode(xerr.REQUEST_PARAM_ERROR)
	}

	userId := ctxdata.GetUidFromCtx(l.ctx)
	addCommentResp, err := l.svcCtx.ArticleServiceRpc.AddComment(l.ctx, &articleservice.AddCommentRequest{
		ArticleId: req.ArticleId,
		Content:   req.Content,
		ParentId:  req.ParentId,
		UserId:    userId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	resp = &types.CommentResp{} // 初始化 resp
	err = copier.Copy(resp, addCommentResp)
	if err != nil {
		l.Logger.Errorf("failed to copy addCommentResp: %v", err)
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}
	return resp, nil
}
