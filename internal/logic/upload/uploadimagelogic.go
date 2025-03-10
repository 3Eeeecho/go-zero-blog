package logic

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/3Eeeecho/go-zero-blog/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/internal/types"
	"github.com/3Eeeecho/go-zero-blog/pkg/app"
	"github.com/3Eeeecho/go-zero-blog/pkg/e"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpLoadImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 上传图片
func NewUpLoadImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpLoadImageLogic {
	return &UpLoadImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpLoadImageLogic) UpLoadImage(req *types.UpLoadImageRequest) (resp *types.UpLoadImageResponse, err error) {
	if req.Image == nil {
		l.Logger.Errorf("image data is empty,req:%+v", req)
		return app.UpLoadImageResponse(e.INVALID_PARAMS, ""), err
	}

	// 生成唯一的文件名（基于时间戳和随机数）
	fileName := fmt.Sprintf("%d_%d.jpg", time.Now().UnixNano(), req.UserId)
	uploadDir := "uploads/images" // 上传目录
	filePath := filepath.Join(uploadDir, fileName)

	// 确保上传目录存在
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		l.Logger.Errorf("failed to create upload directory: %v", err)
		return app.UpLoadImageResponse(e.ERROR_UPLOAD_CREATE_DIR_FAIL, ""), err
	}

	// 将图片数据写入文件
	err = os.WriteFile(filePath, req.Image, 0644)
	if err != nil {
		l.Logger.Errorf("failed to write image to file, path: %s, error: %v", filePath, err)
		return app.UpLoadImageResponse(e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, ""), err
	}

	// 构造图片访问路径（服务运行在本地）
	imageURL := fmt.Sprintf("http://localhost:%d/%s", l.svcCtx.Config.RestConf.Port, filePath)

	l.Logger.Infof("image uploaded successfully, user_id: %d, url: %s", req.UserId, imageURL)
	return app.UpLoadImageResponse(e.SUCCESS, imageURL), nil
}
