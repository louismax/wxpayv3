package core

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/custom"
	"github.com/louismax/wxpayv3/utils"
	"io/ioutil"
	"net/http"
	"time"
)

// ApiCert 商户证书参数
type ApiCert struct {
	ApiSerialNo    string            // 商户API证书序列号
	ApiPrivateKey  *rsa.PrivateKey   // 商户API私钥
	ApiCertificate *x509.Certificate // 商户API证书
}

// GetCertificate 获取微信平台证书(仅从微信获取最新的平台证书)
func (c *PayClient) GetCertificate() ([]custom.CertificateData, error) {
	body, err := c.doRequest(nil, utils.BuildUrl(nil, nil, constant.ApiCertification), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := &custom.CertificateResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}
	result := make([]custom.CertificateData, 0)

	for _, data := range resp.Data {
		item := custom.CertificateData{
			EncryptCertificate: data.EncryptCertificate,
			SerialNo:           data.SerialNo,
			EffectiveTime:      data.EffectiveTime,
			ExpireTime:         data.ExpireTime,
		}
		encryptCert := data.EncryptCertificate
		if encryptCert == nil {
			continue
		}
		decryptCert, err := c.Decrypt(encryptCert.Algorithm, encryptCert.Ciphertext, encryptCert.AssociatedData, encryptCert.Nonce)
		if err != nil {
			return nil, err
		}
		item.DecryptCertificate = string(decryptCert)

		result = append(result, item)
	}
	return result, nil
}

// GetAndSetCertificate 获取并设置微信平台证书（一步到位,从微信获取最新的平台证书，并设置到当前客户端）
func (c *PayClient) GetAndSetCertificate() ([]custom.CertificateData, error) {
	body, err := c.doRequest(nil, utils.BuildUrl(nil, nil, constant.ApiCertification), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := &custom.CertificateResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}
	result := make([]custom.CertificateData, 0)
	if c.PlatformCertMap == nil {
		c.PlatformCertMap = make(map[string]*x509.Certificate)
	}

	for _, data := range resp.Data {
		item := custom.CertificateData{
			EncryptCertificate: data.EncryptCertificate,
			SerialNo:           data.SerialNo,
			EffectiveTime:      data.EffectiveTime,
			ExpireTime:         data.ExpireTime,
		}
		encryptCert := data.EncryptCertificate
		if encryptCert == nil {
			continue
		}
		decryptCert, err := c.Decrypt(encryptCert.Algorithm, encryptCert.Ciphertext, encryptCert.AssociatedData, encryptCert.Nonce)
		if err != nil {
			return nil, err
		}
		item.DecryptCertificate = string(decryptCert)
		ct, err := LoadCertificate(item.DecryptCertificate)
		if err != nil {
			return nil, err
		}
		if c.DefaultPlatformSerialNo == "" {
			c.DefaultPlatformSerialNo = GetCertificateSerialNumber(*ct)
		}
		c.PlatformCertMap[GetCertificateSerialNumber(*ct)] = ct

		result = append(result, item)
	}
	return result, nil
}

// SetClientPlatformCert 设置客户端微信平台证书(通过证书字符串设置当前客户端的平台证书)
func (c *PayClient) SetClientPlatformCert(certificateStr []string) error {
	if c.PlatformCertMap == nil {
		c.PlatformCertMap = make(map[string]*x509.Certificate)
	}
	for _, v := range certificateStr {
		ct, err := LoadCertificate(v)
		if err != nil {
			return err
		}
		if c.DefaultPlatformSerialNo == "" {
			c.DefaultPlatformSerialNo = GetCertificateSerialNumber(*ct)
		}
		c.PlatformCertMap[GetCertificateSerialNumber(*ct)] = ct
	}
	return nil
}

// LoadCertificate 通过证书的文本内容加载证书
func LoadCertificate(certificateStr string) (certificate *x509.Certificate, err error) {
	block, _ := pem.Decode([]byte(certificateStr))
	if block == nil {
		return nil, fmt.Errorf("解码证书错误")
	}
	certificate, err = x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("解析证书错误:%s", err.Error())
	}
	return certificate, nil
}

// LoadPrivateKey 通过私钥的文本内容加载私钥
func LoadPrivateKey(privateKeyStr string) (privateKey *rsa.PrivateKey, err error) {
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		return nil, fmt.Errorf("解码私钥错误")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("解析私钥错误:%s", err.Error())
	}
	privateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("[%s]不是RSA私钥", privateKeyStr)
	}
	return privateKey, nil
}

// LoadPublicKey 通过公钥的文本内容加载公钥
func LoadPublicKey(publicKeyStr string) (publicKey *rsa.PublicKey, err error) {
	block, _ := pem.Decode([]byte(publicKeyStr))
	if block == nil {
		return nil, errors.New("解码公钥错误")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("解析公钥错误:%s", err.Error())
	}
	publicKey, ok := key.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("[%s]不是rsa公钥", publicKeyStr)
	}
	return publicKey, nil
}

// LoadCertificateWithPath  通过证书的文件路径加载证书
func LoadCertificateWithPath(path string) (certificate *x509.Certificate, err error) {
	certificateBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取证书pem文件错误:%s", err.Error())
	}
	return LoadCertificate(string(certificateBytes))
}

// LoadPrivateKeyWithPath 通过私钥的文件路径加载私钥
func LoadPrivateKeyWithPath(path string) (privateKey *rsa.PrivateKey, err error) {
	privateKeyBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取私有pem文件错误:%s", err.Error())
	}
	return LoadPrivateKey(string(privateKeyBytes))
}

// LoadPublicKeyWithPath 通过公钥的文件路径加载公钥
func LoadPublicKeyWithPath(path string) (publicKey *rsa.PublicKey, err error) {
	publicKeyBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取证书pem文件错误:%s", err.Error())
	}
	return LoadPublicKey(string(publicKeyBytes))
}

// GetCertificateSerialNumber 从证书中获取证书序列号
func GetCertificateSerialNumber(certificate x509.Certificate) string {
	return fmt.Sprintf("%X", certificate.SerialNumber)
}

// IsCertExpired 判定证书在特定时间是否过期
func IsCertExpired(certificate x509.Certificate, now time.Time) bool {
	return now.After(certificate.NotAfter)
}

// IsCertValid 判定证书在特定时间是否有效
func IsCertValid(certificate x509.Certificate, now time.Time) bool {
	return now.After(certificate.NotBefore) && now.Before(certificate.NotAfter)
}
