package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)

	// 通用模块
	message[OK] = "SUCCESS"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REQUEST_PARAM_ERROR] = "请求参数错误"
	message[TOKEN_EXPIRE_ERROR] = "token失效,请重新登陆"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "更新数据影响行数为0"
	message[ERROR_INTERNAL] = "内部错误"

	// 用户模块
	message[USER_NOT_FOUND] = "该用户不存在"
	message[USER_ALREADY_EXISTS] = "该用户名已存在"
	message[ERROR_INVALID_PASSWORD] = "无效密码"
	message[ERROR_FORBIDDEN] = "用户权限不足"

	// 文章模块
	message[ARTICLE_NOT_FOUND] = "该文章不存在"
	message[ARTICLE_ALREADY_SUBMITTED] = "文章已提交过"
	message[ERROR_ADD_ARTICLE_FAIL] = "新增文章失败"
	message[ERROR_DELETE_ARTICLE_FAIL] = "删除文章失败"
	message[ERROR_CHECK_EXIST_ARTICLE_FAIL] = "检查文章是否存在失败"
	message[ERROR_EDIT_ARTICLE_FAIL] = "修改文章失败"
	message[ERROR_COUNT_ARTICLE_FAIL] = "统计文章失败"
	message[ERROR_GET_ARTICLES_FAIL] = "获取多个文章失败"
	message[ERROR_GET_ARTICLE_FAIL] = "获取单个文章失败"
	message[ERROR_GEN_ARTICLE_POSTER_FAIL] = "生成文章海报失败"

	// 标签模块
	message[ERROR_EXIST_TAG] = "已存在该标签名称"
	message[ERROR_EXIST_TAG_FAIL] = "获取已存在标签失败"
	message[ERROR_NOT_EXIST_TAG] = "该标签不存在"
	message[ERROR_GET_TAGS_FAIL] = "获取所有标签失败"
	message[ERROR_COUNT_TAG_FAIL] = "统计标签失败"
	message[ERROR_ADD_TAG_FAIL] = "新增标签失败"
	message[ERROR_EDIT_TAG_FAIL] = "修改标签失败"
	message[ERROR_DELETE_TAG_FAIL] = "删除标签失败"
	message[ERROR_EXPORT_TAG_FAIL] = "导出标签失败"
	message[ERROR_IMPORT_TAG_FAIL] = "导入标签失败"

	// 认证模块
	message[ERROR_AUTH_CHECK_TOKEN_FAIL] = "Token鉴权失败"
	message[ERROR_AUTH_CHECK_TOKEN_TIMEOUT] = "Token已超时"
	message[ERROR_AUTH_GENERATE_TOKEN] = "Token生成失败"
	message[ERROR_AUTH_STORE_TOKEN_FAIL] = "Token存储失败"
	message[ERROR_AUTH] = "Token错误"

	// 文件上传模块
	message[ERROR_UPLOAD_SAVE_IMAGE_FAIL] = "保存图片失败"
	message[ERROR_UPLOAD_CHECK_IMAGE_FAIL] = "检查图片失败"
	message[ERROR_UPLOAD_CHECK_IMAGE_FORMAT] = "校验图片错误，图片格式或大小有问题"
	message[ERROR_UPLOAD_CREATE_DIR_FAIL] = "创建上传目录失败"
	message[ERROR_INVALID_FILE] = "打开上传文件失败"

	// 其他模块
	message[ERROR_CREATE_NEWSHEAT_FAIL] = "导入标签时创建表格失败"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
