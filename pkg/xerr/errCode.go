package xerr

// 成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码 (100xxx)
const (
	SERVER_COMMON_ERROR           uint32 = 100001 // 服务器通用错误
	REQUEST_PARAM_ERROR           uint32 = 100002 // 请求参数错误
	TOKEN_EXPIRE_ERROR            uint32 = 100003 // Token 过期
	TOKEN_GENERATE_ERROR          uint32 = 100004 // Token 生成失败
	DB_ERROR                      uint32 = 100005 // 数据库错误
	DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 100006 // 更新数据影响行数为0
	ERROR_INTERNAL                uint32 = 100007 // 内部错误
)

// 用户模块 (101xxx)
const (
	USER_NOT_FOUND         uint32 = 101001 // 用户不存在
	USER_ALREADY_EXISTS    uint32 = 101002 // 用户已存在
	ERROR_INVALID_PASSWORD uint32 = 101003 // 无效密码
	ERROR_FORBIDDEN        uint32 = 101004 //用户权限不足
)

// 文章模块 (102xxx)
const (
	ARTICLE_NOT_FOUND              uint32 = 102001 // 文章不存在
	ARTICLE_ALREADY_SUBMITTED      uint32 = 102002 // 文章已提交
	ERROR_ADD_ARTICLE_FAIL         uint32 = 102003 // 新增文章失败
	ERROR_DELETE_ARTICLE_FAIL      uint32 = 102004 // 删除文章失败
	ERROR_CHECK_EXIST_ARTICLE_FAIL uint32 = 102005 // 检查文章是否存在失败
	ERROR_EDIT_ARTICLE_FAIL        uint32 = 102006 // 修改文章失败
	ERROR_COUNT_ARTICLE_FAIL       uint32 = 102007 // 统计文章失败
	ERROR_GET_ARTICLES_FAIL        uint32 = 102008 // 获取多个文章失败
	ERROR_GET_ARTICLE_FAIL         uint32 = 102009 // 获取单个文章失败
	ERROR_GEN_ARTICLE_POSTER_FAIL  uint32 = 102010 // 生成文章海报失败
)

// 标签模块 (103xxx)
const (
	ERROR_EXIST_TAG       uint32 = 103001 // 已存在该标签名称
	ERROR_EXIST_TAG_FAIL  uint32 = 103002 // 获取已存在标签失败
	ERROR_NOT_EXIST_TAG   uint32 = 103003 // 该标签不存在
	ERROR_GET_TAGS_FAIL   uint32 = 103004 // 获取所有标签失败
	ERROR_COUNT_TAG_FAIL  uint32 = 103005 // 统计标签失败
	ERROR_ADD_TAG_FAIL    uint32 = 103006 // 新增标签失败
	ERROR_EDIT_TAG_FAIL   uint32 = 103007 // 修改标签失败
	ERROR_DELETE_TAG_FAIL uint32 = 103008 // 删除标签失败
	ERROR_EXPORT_TAG_FAIL uint32 = 103009 // 导出标签失败
	ERROR_IMPORT_TAG_FAIL uint32 = 103010 // 导入标签失败
)

// 认证模块 (104xxx)
const (
	ERROR_AUTH_CHECK_TOKEN_FAIL    uint32 = 104001 // Token 鉴权失败
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT uint32 = 104002 // Token 已超时
	ERROR_AUTH_GENERATE_TOKEN      uint32 = 104003 // Token 生成失败
	ERROR_AUTH_STORE_TOKEN_FAIL    uint32 = 104004 // Token 存储失败
	ERROR_AUTH                     uint32 = 104005 // Token 错误
)

// 文件上传模块 (105xxx)
const (
	ERROR_UPLOAD_SAVE_IMAGE_FAIL    uint32 = 105001 // 保存图片失败
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   uint32 = 105002 // 检查图片失败
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT uint32 = 105003 // 校验图片错误
	ERROR_UPLOAD_CREATE_DIR_FAIL    uint32 = 105004 // 创建上传目录失败
	ERROR_INVALID_FILE              uint32 = 105005 // 打开上传文件失败
)

// 其他模块 (106xxx)
const (
	ERROR_CREATE_NEWSHEAT_FAIL uint32 = 106001 // 导入标签时创建表格失败
)
