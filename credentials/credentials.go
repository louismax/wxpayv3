package credentials

import (
	"context"
	"fmt"
	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/core"
	"github.com/louismax/wxpayv3/utils"
	"time"
)

// WechatPayCredentials 微信支付请求报文头 Authorization 信息生成器
type WechatPayCredentials struct {
	Signer core.Signer // 数字签名生成器
}

// GenerateAuthorizationHeader 生成请求报文头中的 Authorization 信息，详见：
// https://wechatpay-api.gitbook.io/wechatpay-api-v3/qian-ming-zhi-nan-1/qian-ming-sheng-cheng
func (c *WechatPayCredentials) GenerateAuthorizationHeader(ctx context.Context,
	method, canonicalURL, signBody string) (authorization string, err error) {
	if c.Signer == nil {
		return "", fmt.Errorf("签名器未初始化！")
	}
	nonce, err := utils.GenerateNonce()
	if err != nil {
		return "", err
	}
	timestamp := time.Now().Unix()
	message := fmt.Sprintf(constant.SignatureMessageFormat, method, canonicalURL, timestamp, nonce, signBody)
	signatureResult, err := c.Signer.Sign(ctx, message)
	if err != nil {
		return "", err
	}
	authorization = fmt.Sprintf(constant.HeaderAuthorizationFormat, c.getAuthorizationType(),
		signatureResult.MchID, nonce, timestamp, signatureResult.CertificateSerialNo, signatureResult.Signature)
	return authorization, nil
}

func (c *WechatPayCredentials) getAuthorizationType() string {
	return "WECHATPAY2-" + c.Signer.Algorithm()
}