package core

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"net/http"
	"time"
)

type DialSettings struct {
	MchId               string            // 商户号
	ApiV3Key            string            // apiV3密钥
	ApiSerialNo         string            // API证书序列号
	ApiPrivateKey       *rsa.PrivateKey   // API私钥
	ApiCertificate      *x509.Certificate // API证书
	PlatformSerialNo    string            // 平台证书序列号
	PlatformCertificate *x509.Certificate // 平台证书
	HttpClient          *http.Client
}

// Validate 校验请求配置是否有效
func (ds *DialSettings) Validate() error {
	if ds.MchId == "" {
		return fmt.Errorf("商户号无效")
	}
	if ds.ApiV3Key == "" {
		return fmt.Errorf("APIv3Key无效")
	}
	if ds.ApiSerialNo == "" {
		return fmt.Errorf("API证书序列号无效")
	}
	if ds.ApiPrivateKey == nil {
		return fmt.Errorf("API私钥无效")
	}
	if ds.ApiCertificate == nil {
		return fmt.Errorf("API证书无效")
	}
	if GetCertificateSerialNumber(*ds.ApiCertificate) != ds.ApiSerialNo {
		return fmt.Errorf("API证书序列号不匹配")
	}
	if IsCertExpired(*ds.ApiCertificate, time.Now()) {
		return fmt.Errorf("API证书已过期")
	}
	if ds.PlatformSerialNo != "" && ds.PlatformCertificate != nil {
		if GetCertificateSerialNumber(*ds.PlatformCertificate) != ds.PlatformSerialNo {
			return fmt.Errorf("平台证书序列号不匹配")
		}
		if IsCertExpired(*ds.PlatformCertificate, time.Now()) {
			return fmt.Errorf("平台证书已过期")
		}
	}

	return nil
}

type BasicInformation struct {
	MchID       string
	MchAPIv3Key string
	ApiCert     ApiCert
}

func (w BasicInformation) Join(o *DialSettings) error {
	o.MchId = w.MchID
	o.ApiV3Key = w.MchAPIv3Key
	o.ApiSerialNo = w.ApiCert.ApiSerialNo
	o.ApiPrivateKey = w.ApiCert.ApiPrivateKey
	o.ApiCertificate = w.ApiCert.ApiCertificate
	return nil
}

func (w PlatformCert) Join(o *DialSettings) error {
	o.PlatformSerialNo = w.PlatformSerialNo
	o.PlatformCertificate = w.PlatformCertificate
	return nil
}
