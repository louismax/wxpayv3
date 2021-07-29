package constant

import "time"

// 请求报文签名相关常量
const (
	SignatureMessageFormat = "%s\n%s\n%d\n%s\n%s\n" // 数字签名原文格式
	// HeaderAuthorizationFormat 请求头中的 Authorization 拼接格式
	HeaderAuthorizationFormat = "%s mchid=\"%s\",nonce_str=\"%s\",timestamp=\"%d\",serial_no=\"%s\",signature=\"%s\""
	AlgorithmAEADAES256GCM    = "AEAD_AES_256_GCM"
)
const (
	FiveMinute     = 5 * 60           // 回包校验最长时间（秒）
	DefaultTimeout = 30 * time.Second // HTTP 请求默认超时时间
)

const ApiDomain = "https://api.mch.weixin.qq.com/"

const (
	ApiCertification          = "/v3/certificates"                                   // 平台证书下载
	ApiUploadImage            = "/v3/merchant/media/upload"                          //图片上传
	ApiQuerySettlementAccount = "/v3/apply4sub/sub_merchants/{sub_mchid}/settlement" //查询结算账户
)
