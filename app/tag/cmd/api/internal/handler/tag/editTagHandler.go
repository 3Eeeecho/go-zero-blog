package tag

import (
	"net/http"

	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/api/internal/logic/tag"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 修改文章标签
func EditTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EditTagRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := tag.NewEditTagLogic(r.Context(), svcCtx)
		resp, err := l.EditTag(&req)
		result.HttpResult(r, w, resp, err)
	}
}
