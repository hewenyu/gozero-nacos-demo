package response

// DefaultCode 基本类型错误码
const DefaultCode = 1001

// DefaultCodeError 基本类型错误结构体
type DefaultCodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// CodeErrorResponse 错误类型返回结构体
type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// CodeError 打印错误信息
func (e *DefaultCodeError) Error() string {
	return e.Msg
}

// Data 错误数据处理
func (e *DefaultCodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
