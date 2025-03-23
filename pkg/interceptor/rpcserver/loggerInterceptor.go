package rpcserver

import (
	"context"
	"errors"

	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if err != nil {
		var codeErr *xerr.CodeError
		if errors.As(err, &codeErr) { // 使用 errors.As 检查错误类型
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)

			// 转成 grpc err
			err = status.Error(codes.Code(codeErr.GetErrCode()), codeErr.GetErrMsg())
		} else {
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
		}
	}

	return resp, err
}
