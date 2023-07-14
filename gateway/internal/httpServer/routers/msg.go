package routers

const (
	SUCCESS                        = "success"
	FAILED                         = "failed"
	ERROR                          = "error"
	INVALID_PARAMS                 = "请求参数错误"
	ERROR_EXIST_TAG                = "已存在该标签名称"
	ERROR_NOT_EXIST_TAG            = "该标签不存在"
	ERROR_NOT_EXIST_ARTICLE        = "该文章不存在"
	ERROR_AUTH_CHECK_TOKEN_FAIL    = "Token鉴权失败"
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = "Token已超时"
	ERROR_AUTH_TOKEN               = "Token生成失败"
	ERROR_AUTH                     = "Token错误"
	ERROR_SERVER_GRPC              = "服务器rpc调用错误"
	ERROR_MESSAGE_TYPE             = "错误的消息类型"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
