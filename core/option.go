package core

import (
	"fmt"
	"net/http"
)

type DialSettings struct {
	HTTPClient *http.Client   // 自定义所使用的 HTTPClient 实例
	Signer    Signer    // 签名器
	Validator  Validator // 应答包签名校验器
	Cipher     Cipher  // 敏感字段加解密套件
}

// Validate 校验请求配置是否有效
func (ds *DialSettings) Validate() error {
	if ds.Validator == nil {
		return fmt.Errorf("validator is required for Client")
	}
	if ds.Signer == nil {
		return fmt.Errorf("signer is required for Client")
	}
	return nil
}