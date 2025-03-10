package upload

import (
	"io"
	"net/http"
	"strconv"

	logic "github.com/3Eeeecho/go-zero-blog/internal/logic/upload"
	"github.com/3Eeeecho/go-zero-blog/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/internal/types"
	"github.com/3Eeeecho/go-zero-blog/pkg/e"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 上传图片
func UpLoadImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 解析 multipart/form-data 请求
		err := r.ParseMultipartForm(10 << 20) // 限制最大 10MB
		if err != nil {
			httpx.OkJson(w, &types.UpLoadImageResponse{
				Code: e.INVALID_PARAMS,
				Msg:  "Failed to parse multipart form: " + err.Error(),
			})
			return
		}

		// 获取图片文件
		file, _, err := r.FormFile("image")
		if err != nil {
			httpx.OkJson(w, &types.UpLoadImageResponse{
				Code: e.INVALID_PARAMS,
				Msg:  "Failed to get image: " + err.Error(),
			})
			return
		}
		defer file.Close()

		// 读取图片数据
		imageData, err := io.ReadAll(file)
		if err != nil {
			httpx.OkJson(w, &types.UpLoadImageResponse{
				Code: e.INVALID_PARAMS,
				Msg:  "Failed to read image: " + err.Error(),
			})
			return
		}

		// 获取用户ID（可选）
		userId := r.FormValue("user_id") // 如果需要，从表单中获取
		req := &types.UpLoadImageRequest{
			UserId: 0, // 默认值，可根据实际需求解析
			Image:  imageData,
		}
		var id int
		if userId != "" {
			id, err = strconv.Atoi(userId)
			if err != nil {
				httpx.OkJson(w, &types.UpLoadImageResponse{
					Code: e.INVALID_PARAMS,
					Msg:  "Failed to parse user_id: " + err.Error(),
				})
				return
			}
			req.UserId = int64(id)
		}

		l := logic.NewUpLoadImageLogic(r.Context(), svcCtx)
		resp, err := l.UpLoadImage(req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
