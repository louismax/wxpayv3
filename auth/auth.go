package auth

import (
	"context"
	"net/http"
)

// SignatureResult 数字签名结果
type SignatureResult struct {
	MchID               string // 商户号
	CertificateSerialNo string // 签名对应的证书序列号
	Signature           string // 签名内容
}

// Deprecated:Credential 请求报文头 Authorization 信息生成器
type Credential interface {
	GenerateAuthorizationHeader(ctx context.Context, method, canonicalURL, signBody string) (authorization string, err error)
}

// Deprecated:Signer 数字签名生成器
type Signer interface {
	Sign(ctx context.Context, message string) (*SignatureResult, error) // 对信息进行签名
	Algorithm() string                                                  // 返回使用的签名算法
}

// Deprecated:Validator  应答报文验证器
type Validator interface {
	Validate(ctx context.Context, response *http.Response) error // 对 HTTP 应答报文进行验证
}
