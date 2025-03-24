package tag

import (
	"context"
	"io"
	"mime/multipart"

	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/tagservice"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 导入标签信息
func NewImportTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImportTagLogic {
	return &ImportTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportTagLogic) ImportTag(req *types.ImportTagRequest, file *multipart.FileHeader) (resp *types.Response, err error) {
	// 打开上传的文件
	fileReader, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fileReader.Close()

	// 读取文件内容
	fileContent, err := io.ReadAll(fileReader)
	if err != nil {
		return nil, err
	}

	editTagResp, err := l.svcCtx.TagServiceRpc.ImportTag(l.ctx, &tagservice.ImportTagRequest{
		FileContent: fileContent,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	resp = &types.Response{} // 初始化 resp
	err = copier.Copy(resp, editTagResp)
	if err != nil {
		l.Logger.Errorf("failed to copy editTagResp: %v", err)
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}
	return resp, nil
}
