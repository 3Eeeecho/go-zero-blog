package tag

import (
	"net/http"

	logic "github.com/3Eeeecho/go-zero-blog/internal/logic/tag"
	"github.com/3Eeeecho/go-zero-blog/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/internal/types"
	"github.com/3Eeeecho/go-zero-blog/pkg/e"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 导入标签信息
func ImportTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		maxBytes := int64(svcCtx.Config.App.ImageMaxSize << 20)
		err := r.ParseMultipartForm(maxBytes) // 限制最大 10MB
		if err != nil {
			httpx.OkJson(w, &types.Response{
				Code: e.INVALID_PARAMS,
				Msg:  "Failed to parse multipart form",
			})
			return
		}

		// 获取文件
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			httpx.OkJson(w, &types.Response{
				Code: e.INVALID_PARAMS,
				Msg:  "Failed to get file from request",
			})
			return
		}
		defer file.Close()

		var req types.ImportTagRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewImportTagLogic(r.Context(), svcCtx)
		resp, err := l.ImportTag(&req, fileHeader)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
