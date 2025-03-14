package app

import (
	"github.com/3Eeeecho/go-zero-blog/internal/types"
	"github.com/3Eeeecho/go-zero-blog/pkg/e"
)

func Response(errCode int, data any) *types.Response {
	return &types.Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	}
}

func ResponseMsg(errCode int, msg string, data any) *types.Response {
	return &types.Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	}
}

func GetArticlesResponse(code int, data interface{}, total int64, pageNum, pageSize int) *types.GetArticlesResponse {
	return &types.GetArticlesResponse{
		Code:     code,
		Msg:      e.GetMsg(code),
		Data:     data,
		Total:    total,
		PageNum:  pageNum,
		PageSize: pageSize,
	}
}

func GetTagsResponse(code int, data interface{}, total int64, pageNum, pageSize int) *types.GetTagsResponse {
	return &types.GetTagsResponse{
		Code:     code,
		Msg:      e.GetMsg(code),
		Data:     data,
		Total:    total,
		PageNum:  pageNum,
		PageSize: pageSize,
	}
}

func LoginResponse(code int, token string) *types.LoginResponse {
	return &types.LoginResponse{
		Code:    code,
		Msg:     e.GetMsg(code),
		Token:   token,
		Expires: 0,
	}
}

func UpLoadImageResponse(code int, imageUrl string) *types.UpLoadImageResponse {
	return &types.UpLoadImageResponse{
		Code:     code,
		Msg:      e.GetMsg(code),
		ImageURL: imageUrl,
	}
}

func ExportTagsResponse(code int, ExportUrl, ExportSaveUrl string) *types.ExportTagResponse {
	return &types.ExportTagResponse{
		Code:          code,
		Msg:           e.GetMsg(code),
		ExportUrl:     ExportUrl,
		ExportSaveUrl: ExportSaveUrl,
	}
}
