package core

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"github.com/louismax/wxpayv3/custom"
	"net/http"
	"time"
)

// DialSettings DialSettings
type DialSettings struct {
	MchId                string                       // 商户号
	ApiV3Key             string                       // apiV3密钥
	ApiSerialNo          string                       // API证书序列号
	ApiPrivateKey        *rsa.PrivateKey              // API证书私钥
	ApiCertificate       *x509.Certificate            // API证书(非必须，可获取证书序列号和商户API公钥)
	DefaultSerialNo      string                       // 默认平台证书序列号
	PlatformCertMap      map[string]*x509.Certificate // 平台证书集合
	WechatPayPublicKeyID string                       // 平台公钥ID
	WechatPayPublicKey   *rsa.PublicKey               // 平台公钥
	HttpClient           *http.Client                 // http客户端
}

// Validate 校验请求配置是否有效
func (ds *DialSettings) Validate() error {
	if ds.MchId == "" {
		return fmt.Errorf("商户号无效")
	}
	if ds.ApiPrivateKey == nil {
		return fmt.Errorf("商户API私钥无效")
	}
	if ds.ApiSerialNo == "" {
		if ds.ApiCertificate != nil {
			//通过商户证书获取证书编号
			ds.ApiSerialNo = GetCertificateSerialNumber(*ds.ApiCertificate)
		} else {
			return fmt.Errorf("商户API证书序列号无效")
		}
	}
	if ds.ApiCertificate != nil {
		if GetCertificateSerialNumber(*ds.ApiCertificate) != ds.ApiSerialNo {
			return fmt.Errorf("商户API证书序列号不匹配")
		}
		if IsCertExpired(*ds.ApiCertificate, time.Now()) {
			return fmt.Errorf("商户API证书已过期")
		}
	}

	if len(ds.PlatformCertMap) > 0 {
		for k, v := range ds.PlatformCertMap {
			if GetCertificateSerialNumber(*v) != k {
				return fmt.Errorf("微信平台证书序列号不匹配")
			}
			if IsCertExpired(*v, time.Now()) {
				return fmt.Errorf("平台证书已过期")
			}
		}
	}
	return nil
}

// BasicInformation 基础参数
type BasicInformation struct {
	MchID       string
	MchAPIv3Key string
	ApiCert     ApiCert
}

// Join as join
func (w BasicInformation) Join(o *DialSettings) error {
	o.MchId = w.MchID
	o.ApiV3Key = w.MchAPIv3Key
	o.ApiSerialNo = w.ApiCert.ApiSerialNo
	o.ApiPrivateKey = w.ApiCert.ApiPrivateKey
	o.ApiCertificate = w.ApiCert.ApiCertificate
	return nil
}

// PlatformCert 微信平台证书参数
type PlatformCert struct {
	DefaultSerialNo string // 默认平台证书序列号
	CertList        []custom.CertificateDataList
}

// Join as join
func (w PlatformCert) Join(o *DialSettings) error {
	o.DefaultSerialNo = w.DefaultSerialNo
	if o.PlatformCertMap == nil {
		o.PlatformCertMap = make(map[string]*x509.Certificate)
	}
	for _, v := range w.CertList {
		o.PlatformCertMap[v.SerialNo] = v.Certificate
	}
	return nil
}

// PlatformPubKey 微信平台公钥参数
type PlatformPubKey struct {
	PubKeyId string         // 公钥ID
	PubKey   *rsa.PublicKey // 平台公钥
}

// Join as join
func (w PlatformPubKey) Join(o *DialSettings) error {
	o.WechatPayPublicKeyID = w.PubKeyId
	o.WechatPayPublicKey = w.PubKey
	return nil
}
