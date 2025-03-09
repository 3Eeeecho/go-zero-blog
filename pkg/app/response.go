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
