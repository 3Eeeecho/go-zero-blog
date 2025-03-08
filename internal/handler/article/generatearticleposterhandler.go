package handler

import (
	"net/http"

	logic "github.com/3Eeeecho/go-zero-blog/internal/logic/article"
	"github.com/3Eeeecho/go-zero-blog/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 生成文章海报
func GenerateArticlePosterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGenerateArticlePosterLogic(r.Context(), svcCtx)
		resp, err := l.GenerateArticlePoster()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
