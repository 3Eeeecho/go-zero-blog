package tag

import (
	"net/http"

	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/api/internal/logic/tag"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 导出标签信息
func ExportTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExportTagRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := tag.NewExportTagLogic(r.Context(), svcCtx)
		resp, err := l.ExportTag(&req)
		result.HttpResult(r, w, resp, err)
	}
}
