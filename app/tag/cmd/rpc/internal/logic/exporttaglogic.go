package logic

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExportTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExportTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportTagLogic {
	return &ExportTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ExportTag 导出标签信息
func (l *ExportTagLogic) ExportTag(in *pb.ExportTagRequest) (*pb.ExportTagResponse, error) {
	// 构造查询条件，支持按名称和状态过滤
	conditions := make(map[string]interface{})
	if in.Name != "" {
		conditions["name"] = in.Name
	}
	if in.State != 0 { // 假设 0 表示未指定状态
		conditions["state"] = in.State
	}

	// 获取标签列表
	tags, err := l.svcCtx.TagModel.GetOnCondition(l.ctx, 1, l.svcCtx.Config.App.PageSize, conditions)
	if err != nil {
		l.Logger.Errorf("get tags failed, page_num: %d, page_size: %d, conditions: %v, error: %v",
			1, l.svcCtx.Config.App.PageSize, conditions, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get tags failed: %v", err)
	}
	if len(tags) == 0 {
		l.Logger.Infof("no tags found, in: %+v", in)
		return nil, xerr.NewErrCode(xerr.ERROR_NOT_EXIST_TAG)
	}

	// 创建新的 Excel 文件
	file := excelize.NewFile()
	sheetName := "标签信息"
	_, err = file.NewSheet(sheetName)
	if err != nil {
		l.Logger.Errorf("failed to create sheet %s: %v", sheetName, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "create sheet failed: %v", err)
	}

	// 设置标题行
	titles := []string{"ID", "名称", "创建人", "创建时间", "修改人", "修改时间"}
	for i, title := range titles {
		cell := fmt.Sprintf("%c%d", 'A'+i, 1)
		if err := file.SetCellValue(sheetName, cell, title); err != nil {
			l.Logger.Errorf("failed to set title %s at %s: %v", title, cell, err)
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "set cell value failed: %v", err)
		}
	}

	// 填充标签数据
	for i, v := range tags {
		row := i + 2
		file.SetCellValue(sheetName, fmt.Sprintf("A%d", row), v.Id)
		file.SetCellValue(sheetName, fmt.Sprintf("B%d", row), v.Name)
		file.SetCellValue(sheetName, fmt.Sprintf("C%d", row), v.CreatedBy)
		file.SetCellValue(sheetName, fmt.Sprintf("D%d", row), time.Now().Unix())
		file.SetCellValue(sheetName, fmt.Sprintf("E%d", row), v.ModifiedBy)
		file.SetCellValue(sheetName, fmt.Sprintf("F%d", row), time.Now().Unix())
	}

	// 生成文件名并保存文件
	timeStamp := strconv.Itoa(int(time.Now().Unix()))
	filename := "tags-" + timeStamp + ".xlsx"
	filePath := l.svcCtx.Config.App.ExportSavePath + filename

	// 确保导出目录存在
	if err := os.MkdirAll(l.svcCtx.Config.App.ExportSavePath, os.ModePerm); err != nil {
		l.Logger.Errorf("failed to create export directory %s: %v", l.svcCtx.Config.App.ExportSavePath, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "create export directory failed: %v", err)
	}

	// 保存 Excel 文件
	if err = file.SaveAs(filePath); err != nil {
		l.Logger.Errorf("failed to save excel file %s: %v", filePath, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "save excel file failed: %v", err)
	}

	// 构造导出 URL 和本地保存路径
	exportURL := fmt.Sprintf("http://localhost:%d/%s", 8002, filePath)
	exportSaveURL := filePath

	// 记录成功日志并返回响应
	l.Logger.Infof("tags exported successfully, tag_count: %d, export_url: %s", len(tags), exportURL)
	return &pb.ExportTagResponse{
		Msg:           "导出标签成功",
		ExportUrl:     exportURL,
		ExportSaveUrl: exportSaveURL,
	}, nil
}
