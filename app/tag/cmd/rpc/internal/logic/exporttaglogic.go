package logic

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/pb"
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

// 导出标签信息
func (l *ExportTagLogic) ExportTag(in *pb.ExportTagRequest) (*pb.ExportTagResponse, error) {
	// 获取所有tags
	// 构造查询条件，支持按名称和状态过滤
	conditions := make(map[string]any)
	if in.Name != "" {
		conditions["name"] = in.Name
	}
	if in.State != 0 { // 假设 0 表示未指定状态
		conditions["state"] = in.State
	}

	// 获取标签列表
	tags, err := l.svcCtx.TagModel.GetAll(l.ctx, 1, l.svcCtx.Config.App.PageSize, conditions)
	if err != nil {
		l.Logger.Errorf("get tags failed, page_num: %d, page_size: %d, conditions: %v, error: %v",
			1, l.svcCtx.Config.App.PageSize, conditions, err)
		return nil, err
	}
	if len(tags) == 0 {
		l.Logger.Infof("no tags found, in: %+v", in)
		return nil, errors.New("标签列表为空")
	}

	//导出标签信息
	file := excelize.NewFile()
	sheetName := "标签信息"
	_, err = file.NewSheet(sheetName)
	if err != nil {
		return nil, err
	}

	// 设置标题行
	titles := []string{"ID", "名称", "创建人", "创建时间", "修改人", "修改时间"}
	for i, title := range titles {
		cell := fmt.Sprintf("%c%d", 'A'+i, 1) // 修正为字符类型计算
		if err := file.SetCellValue(sheetName, cell, title); err != nil {
			l.Logger.Errorf("failed to set title %s at %s: %v", title, cell, err)
			return nil, err
		}
	}

	for i, v := range tags {
		row := i + 2
		file.SetCellValue(sheetName, fmt.Sprintf("A%d", row), v.Id)
		file.SetCellValue(sheetName, fmt.Sprintf("B%d", row), v.Name)
		file.SetCellValue(sheetName, fmt.Sprintf("C%d", row), v.CreatedBy)
		file.SetCellValue(sheetName, fmt.Sprintf("D%d", row), time.Now().Unix())
		file.SetCellValue(sheetName, fmt.Sprintf("E%d", row), v.ModifiedBy)
		file.SetCellValue(sheetName, fmt.Sprintf("F%d", row), time.Now().Unix())
	}

	timeStamp := strconv.Itoa(int(time.Now().Unix()))
	filename := "tags-" + timeStamp + ".xlsx"
	filePath := l.svcCtx.Config.App.ExportSavePath + filename

	// 确保 export 目录存在
	if err := os.MkdirAll(l.svcCtx.Config.App.ExportSavePath, os.ModePerm); err != nil {
		return nil, err
	}

	if err = file.SaveAs(filePath); err != nil {
		return nil, err
	}

	exportURL := fmt.Sprintf("http://localhost:%d/%s", 8002, filePath)
	exportSaveURL := filePath // 本地保存路径

	l.Logger.Infof("tags exported successfully, tag_count: %d, export_url: %s", len(tags), exportURL)
	return &pb.ExportTagResponse{
		Msg:           "导出标签成功",
		ExportUrl:     exportURL,
		ExportSaveUrl: exportSaveURL,
	}, nil
}
