package logic

import (
	"context"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/3Eeeecho/go-zero-blog/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/internal/types"
	"github.com/3Eeeecho/go-zero-blog/model"
	"github.com/3Eeeecho/go-zero-blog/pkg/app"
	"github.com/3Eeeecho/go-zero-blog/pkg/e"
	"github.com/xuri/excelize/v2"

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
	// 记录请求开始
	l.Logger.Infof("Starting to import tags with request: %+v", req)

	if file == nil {
		return app.Response(e.INVALID_PARAMS, nil), err
	}

	f, err := file.Open()
	if err != nil {
		l.Logger.Errorf("failed to open uploaded file: %v", err)
		return app.Response(e.ERROR_INVALID_FILE, nil), err
	}
	defer f.Close()

	excelFile, err := excelize.OpenReader(f)
	if err != nil {
		l.Logger.Errorf("failed to parse excel file: %v", err)
		return app.Response(e.ERROR, nil), err
	}
	defer excelFile.Close()

	// 获取指定 Sheet
	sheetName := "标签信息"
	rows, err := excelFile.GetRows(sheetName)
	if err != nil {
		l.Logger.Errorf("failed to get rows from sheet %s: %v", sheetName, err)
		return app.Response(e.ERROR, nil), err
	}

	// 检查表头
	if len(rows) < 1 {
		l.Logger.Errorf("excel file is empty")
		return app.ResponseMsg(e.ERROR, "Invalid Excel file headers", nil), nil
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
		return app.ResponseMsg(e.ERROR, "Failed to import tags", nil), err
	}

	l.Logger.Infof("Successfully imported %d tags", len(tags))
	return app.Response(e.SUCCESS, nil), nil
}

func parseTimestamp(ts string) (time.Time, error) {
	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(timestamp, 0), nil
}
