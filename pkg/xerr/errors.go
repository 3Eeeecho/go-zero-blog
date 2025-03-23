package xerr

import "fmt"

// CodeError 包含错误码和错误消息的结构体
type CodeError struct {
	errCode uint32
	errMsg  string
}

// GetErrCode 返回错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

// GetErrMsg 返回错误消息
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

// Error 实现 error 接口，返回格式化的错误信息
func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d,ErrMsg:%s", e.errCode, e.errMsg)
}

// NewErrCodeMsg 创建自定义错误码和消息
func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}

// NewErrCode 根据错误码创建错误，使用预定义消息
func NewErrCode(errCode uint32) *CodeError {
	return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
}

// NewErrMsg 创建自定义消息的错误，默认使用通用错误码
func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: SERVER_COMMON_ERROR, errMsg: errMsg}
}
