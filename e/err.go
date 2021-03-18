package e

const (
	SUCCESS                        = 200   //"ok",
	ERROR                          = 500   //"fail",
	INVALID_PARAMS                 = 400   //"请求参数错误",
	ERROR_EXIST_TAG                = 10001 //"已存在该标签名称",
	ERROR_NOT_EXIST_TAG            = 10002 //"该标签不存在",
	ERROR_NOT_EXIST_ARTICLE        = 10003 //"该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001 //"Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002 //"Token已超时",
	ERROR_AUTH_TOKEN               = 20003 //"Token生成失败",
	ERROR_AUTH                     = 20004 //"Token错误",
)
