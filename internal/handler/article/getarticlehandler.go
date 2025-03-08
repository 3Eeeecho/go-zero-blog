package handler

import (
	"net/http"

	logic "github.com/3Eeeecho/go-zero-blog/internal/logic/article"
	"github.com/3Eeeecho/go-zero-blog/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取单篇文章的详细信息
func GetArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetArticleLogic(r.Context(), svcCtx)
		resp, err := l.GetArticle()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
