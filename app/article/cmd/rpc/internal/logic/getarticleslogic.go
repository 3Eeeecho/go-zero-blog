package logic

import (
	"context"
	"errors"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticlesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticlesLogic {
	return &GetArticlesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取文章列表
func (l *GetArticlesLogic) GetArticles(in *pb.GetArticlesRequest) (*pb.GetArticlesResponse, error) {
	// 设置分页默认值
	pageNum := in.PageNum
	pageSize := in.PageSize
	if pageNum <= 0 {
		pageNum = 1 // 默认第1页
	}
	if pageSize <= 0 {
		pageSize = 10 // 默认每页10条
	}

	// 构造过滤条件
	maps := make(map[string]any)
	if in.TagId != 0 {
		maps["tag_id"] = in.TagId
	}
	if in.State != 0 {
		maps["state"] = in.State
	}

	articles, err := l.svcCtx.ArticleModel.GetArticles(l.ctx, int(pageNum), int(pageSize), maps)
	if err != nil {
		l.Logger.Errorf("get articles failed, page_num: %d, page_size: %d, maps: %v, error: %v",
			pageNum, pageSize, maps, err)
		return nil, err
	}

	// 文章总数
	total, err := l.svcCtx.ArticleModel.CountByCondition(l.ctx, maps)
	if err != nil {
		l.Logger.Errorf("count articles failed,condition:%v,error:%v", maps, err)
		return nil, errors.New("count articles failed")
	}

	data := make([]*pb.Article, 0, len(articles))
	//手动填充
	for _, article := range articles {
		data = append(data, &pb.Article{
			Id:         int64(article.Id),
			TagId:      int64(article.TagId),
			Title:      article.Title,
			Desc:       article.Desc,
			Content:    article.Content,
			State:      int32(article.State),
			CreatedBy:  article.CreatedBy,
			ModifiedBy: article.ModifiedBy,
		})
	}

	// 返回成功响应
	l.Logger.Infof("articles retrieved successfully, page_num: %d, page_size: %d", pageNum, pageSize)

	return &pb.GetArticlesResponse{
		Msg:      "获取文章列表成功",
		Data:     data,
		Total:    total,
		PageNum:  pageNum,
		PageSize: pageSize,
	}, nil
}
