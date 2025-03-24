package tag

import (
	"net/http"

	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/api/internal/logic/tag"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 导入标签信息
func ImportTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		maxBytes := int64(svcCtx.Config.App.ImageMaxSize << 20)
		err := r.ParseMultipartForm(maxBytes) // 限制最大 10MB
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, nil)
			return
		}

		// 获取文件
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, nil)
			return
		}
		defer file.Close()

		var req types.ImportTagRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := tag.NewImportTagLogic(r.Context(), svcCtx)
		resp, err := l.ImportTag(&req, fileHeader)
		result.HttpResult(r, w, resp, err)
	}
}
