package constant

import "time"

// 请求报文签名相关常量
const (
	SignatureMessageFormat = "%s\n%s\n%d\n%s\n%s\n" // 数字签名原文格式
	// HeaderAuthorizationFormat 请求头中的 Authorization 拼接格式
	HeaderAuthorizationFormat = "%s mchid=\"%s\",nonce_str=\"%s\",timestamp=\"%d\",serial_no=\"%s\",signature=\"%s\""
)
const (
	FiveMinute     = 5 * 60           // 回包校验最长时间（秒）
	DefaultTimeout = 30 * time.Second // HTTP 请求默认超时时间
)
