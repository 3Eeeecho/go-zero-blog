package tag

import (
	"net/http"

	logic "github.com/3Eeeecho/go-zero-blog/internal/logic/tag"
	"github.com/3Eeeecho/go-zero-blog/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 删除文章标签
func DeleteTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDeleteTagLogic(r.Context(), svcCtx)
		resp, err := l.DeleteTag()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
