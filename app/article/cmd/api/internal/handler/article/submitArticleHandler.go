package article

import (
	"net/http"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/api/internal/logic/article"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SubmitArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SubmitArticleRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := article.NewSubmitArticleLogic(r.Context(), svcCtx)
		resp, err := l.SubmitArticle(&req)
		result.HttpResult(r, w, resp, err)
	}
}
