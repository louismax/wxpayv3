package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/custom"
)

// ParseComplaintNotification 必须传入未经重编码的原始HTTP请求体，先验签再解密。
// SDK只返回通知内容，通知去重、可靠保存以及HTTP应答由调用方负责。
func (c *PayClient) ParseComplaintNotification(header http.Header, body []byte) (*custom.ComplaintNotification, *custom.ComplaintNotificationResource, error) {
	serial := header.Get("Wechatpay-Serial")
	if serial == "" {
		return nil, nil, fmt.Errorf("投诉通知缺少Wechatpay-Serial")
	}
	// 通知不能沿用缺少验签配置时仅告警的响应处理行为。
	if strings.HasPrefix(serial, "PUB_KEY_ID_") {
		if c.WechatPayPublicKeyID != serial || c.WechatPayPublicKey == nil {
			return nil, nil, fmt.Errorf("投诉通知的微信支付公钥未配置或ID不匹配")
		}
	} else if cert := c.PlatformCertMap[serial]; cert == nil {
		return nil, nil, fmt.Errorf("投诉通知对应的微信支付平台证书未配置")
	}
	if err := c.verify(serial, header.Get("Wechatpay-Signature"), header.Get("Wechatpay-Timestamp"), header.Get("Wechatpay-Nonce"), body); err != nil {
		return nil, nil, err
	}
	notification := custom.ComplaintNotification{}
	if err := json.Unmarshal(body, &notification); err != nil {
		return nil, nil, err
	}
	resource := notification.Resource
	if resource.Algorithm != constant.AlgorithmAEADAES256GCM || len(resource.Nonce) != 12 || len(c.ApiV3Key) != 32 {
		return nil, nil, fmt.Errorf("投诉通知加密算法、nonce或APIv3密钥长度不正确")
	}
	plaintext, err := c.Decrypt(resource.Algorithm, resource.Ciphertext, resource.AssociatedData, resource.Nonce)
	if err != nil {
		return nil, nil, err
	}
	result := custom.ComplaintNotificationResource{}
	if err = json.Unmarshal(plaintext, &result); err != nil {
		return nil, nil, err
	}
	if notification.ID == "" || result.ComplaintID == "" || result.ActionType == "" {
		return nil, nil, fmt.Errorf("投诉通知缺少通知ID、投诉单号或动作类型")
	}
	return &notification, &result, nil
}
