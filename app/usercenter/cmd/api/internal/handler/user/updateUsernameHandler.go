package user

import (
	"net/http"

	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/api/internal/logic/user"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 修改用户名
func UpdateUsernameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUsernameRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewUpdateUsernameLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUsername(&req)
		result.HttpResult(r, w, resp, err)
	}
}
