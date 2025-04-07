package user

import (
	"net/http"

	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/api/internal/logic/user"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/pkg/result"
	"github.com/go-playground/validator"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 修改用户名
func UpdateNicknameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateNicknameRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		// 校验请求参数
		validate := validator.New()
		if err := validate.Struct(&req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewUpdateNicknameLogic(r.Context(), svcCtx)
		resp, err := l.UpdateNickname(&req)
		result.HttpResult(r, w, resp, err)
	}
}
