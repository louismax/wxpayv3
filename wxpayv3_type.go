package wxpayv3

const (
	ServerUrl   = "https://api.mch.weixin.qq.com"
	ContentType = "application/json"
	AcceptType  = "*/*"
	AuthType    = "WECHATPAY2-SHA256-RSA2048"
)

type Param interface {
	// 用于提供访问的url后缀
	APIUrl() string
	// 用于提供访问的 method
	Method() string
	// 返回参数列表
	Params() map[string]string
	// RawJsonStr 请求JSON
	RawJsonStr() string
}

type SysError struct {
	Code    string `json:"code"`
	Message string `json:"message"`

	Detail           Detail_bak      `json:"detail"`
	Field            string          `json:"field"`
	Location         string          `json:"location"`
	Sign_information Signinformation `json:"sign_information"`
}

type Detail_bak struct {
	Detail Detail `json:"detail"`
}

type Detail struct {
	Issue string `json:"issue"`
}

type Signinformation struct {
	Method                 string `json:"method"`
	Sign_message_length    string `json:"sign_message_length"`
	Truncated_sign_message string `json:"truncated_sign_message"`
	Url                    string `json:"url"`
}
