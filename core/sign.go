package core

import (
	"fmt"
	"github.com/louismax/wxpayv3/utils"
	"time"
)

//CertificationType CertificationType
const CertificationType = "WECHATPAY2-SHA256-RSA2048"

//Authorization 获取WechatPayV3的header信息Authorization
func (c *PayClient) Authorization(httpMethod string, urlString string, body []byte) (string, error) {
	token, err := c.Token(httpMethod, urlString, body)
	if err != nil {
		return "", err
	}
	return CertificationType + " " + token, nil
}

//Token 获取签名信息 请求方法为GET时，报文主体为空;当请求方法为POST或PUT时，请使用真实发送的JSON报文;图片上传API，请使用meta对应的JSON报文
func (c *PayClient) Token(httpMethod string, rawUrl string, body []byte) (string, error) {
	nonce, err := utils.GenerateNonce()
	if err != nil {
		return "", err
	}
	timestamp := time.Now().Unix()
	message, err := utils.BuildMessage(httpMethod, rawUrl, body, nonce, timestamp)
	if err != nil {
		return "", err
	}
	signature, err := c.Sign(message)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(`mchid="%s",nonce_str="%s",signature="%s",timestamp="%v",serial_no="%s"`, c.MchId, nonce, signature, timestamp, c.ApiSerialNo), nil
}

//Sign Sign
func (c *PayClient) Sign(message []byte) (string, error) {
	return utils.Sign(message, c.ApiPrivateKey)
}
