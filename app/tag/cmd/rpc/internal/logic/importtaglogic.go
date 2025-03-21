package logic

import (
	"bytes"
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/app/tag/model"
	"github.com/xuri/excelize/v2"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewImportTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImportTagLogic {
	return &ImportTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 导入标签信息
func (l *ImportTagLogic) ImportTag(in *pb.ImportTagRequest) (*pb.ImportTagResponse, error) {
	if in.FileContent == nil {
		l.Logger.Errorf("file content is empty")
		return nil, errors.New("上传内容为空")
	}

	// 打印文件内容大小
	l.Logger.Infof("file content size: %d bytes", len(in.FileContent))

	excelFile, err := excelize.OpenReader(bytes.NewReader(in.FileContent))
	if err != nil {
		l.Logger.Errorf("failed to parse excel file: %v", err)
		return nil, err
	}
	defer excelFile.Close()

	// 获取指定 Sheet
	sheetName := "标签信息"
	rows, err := excelFile.GetRows(sheetName)
	if err != nil {
		l.Logger.Errorf("failed to get rows from sheet %s: %v", sheetName, err)
		return nil, err
	}

	// 检查表头
	if len(rows) < 1 {
		l.Logger.Errorf("excel file is empty")
		return nil, errors.New("标签列表为空")
	}

	// 解析并导入数据
	var tags []*model.BlogTag
	for i, row := range rows[1:] { // 从第二行开始（跳过表头）
		if len(row) < 6 {
			l.Logger.Errorf("skipping row %d: insufficient columns", i+2)
			continue
		}

		createdOn, err := parseTimestamp(row[3])
		if err != nil {
			l.Logger.Errorf("skipping row %d: invalid created_at %s: %v", i+2, row[3], err)
			continue
		}
		modifiedOn, err := parseTimestamp(row[5])
		if err != nil {
			l.Logger.Errorf("skipping row %d: invalid modified_at %s: %v", i+2, row[5], err)
			continue
		}

		tag := &model.BlogTag{
			Name:       row[1],
			CreatedBy:  row[2],
			CreatedOn:  createdOn.Unix(),
			ModifiedBy: row[4],
			ModifiedOn: modifiedOn.Unix(),
		}
		tags = append(tags, tag)
	}

	// 批量插入数据库
	err = l.svcCtx.TagModel.InsertBatch(l.ctx, tags)
	if err != nil {
		l.Logger.Errorf("failed to insert tags into database: %v", err)
		return nil, err
	}

	l.Logger.Infof("Successfully imported %d tags", len(tags))

	//TODO 填充fileurl,使用静态路由获取
	return &pb.ImportTagResponse{
		Msg:     "导入标签成功",
		FileUrl: "",
	}, nil
}

func parseTimestamp(ts string) (time.Time, error) {
	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(timestamp, 0), nil
}
