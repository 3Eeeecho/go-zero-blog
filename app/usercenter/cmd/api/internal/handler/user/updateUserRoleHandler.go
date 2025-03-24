package user

import (
	"net/http"

	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/api/internal/logic/user"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 更新用户权限
func UpdateUserRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserRoleRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewUpdateUserRoleLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUserRole(&req)
		result.HttpResult(r, w, resp, err)
	}
}
