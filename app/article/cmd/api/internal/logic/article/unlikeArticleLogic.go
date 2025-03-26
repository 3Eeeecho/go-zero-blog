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

type UnlikeArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 取消点赞文章
func NewUnlikeArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnlikeArticleLogic {
	return &UnlikeArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnlikeArticleLogic) UnlikeArticle(req *types.LikeArticleRequest) (resp *types.ArticleCommonResponse, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	unlikeArticleResp, err := l.svcCtx.ArticleServiceRpc.UnlikeArtilce(l.ctx, &articleservice.UnlikeArticleRequest{
		ArticleId: req.Article_id,
		UserId:    userId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	resp = &types.ArticleCommonResponse{}
	err = copier.Copy(resp, unlikeArticleResp)
	if err != nil {
		l.Logger.Errorf("failed to copy unlikeArticleResp: %v", err)
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}
	return resp, nil
}
