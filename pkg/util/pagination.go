package util

import (
	"net/http"
	"strconv"
)

func GetOffset(r *http.Request, defaultPageSize int) int {
	result := 0

	// 从查询参数获取 page
	pageStr := r.URL.Query().Get("page_num") // 对应 api 中的 PageNum
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1 // 默认第 1 页
	}

	//获取 page_size
	pageSize := GetPageSize(r, defaultPageSize)

	// 计算偏移量
	result = (page - 1) * pageSize
	if result < 0 {
		result = 0
	}

	return result
}

func GetPageSize(r *http.Request, defaultPageSize int) int {
	pageSizeStr := r.URL.Query().Get("page_size")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		return defaultPageSize
	}
	return pageSize
}
