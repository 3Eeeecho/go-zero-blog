package app

import (
	"github.com/3Eeeecho/go-zero-blog/internal/types"
	"github.com/3Eeeecho/go-zero-blog/pkg/e"
)

func Response(errCode int, data any, err error) (*types.Response, error) {
	return &types.Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	}, err
}
