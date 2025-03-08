package upload

import (
	"net/http"

	logic "github.com/3Eeeecho/go-zero-blog/internal/logic/upload"
	"github.com/3Eeeecho/go-zero-blog/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 上传图片
func UpLoadImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpLoadImageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpLoadImageLogic(r.Context(), svcCtx)
		resp, err := l.UpLoadImage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
