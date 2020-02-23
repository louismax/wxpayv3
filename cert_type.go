package wxpayv3

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// CertifiCates
type CertifiCates struct {
}

// APIUrl CertifiCates APIURL
func (this CertifiCates) APIUrl() string {
	return "/v3/certificates"
}

// Method CertifiCates Method
func (this CertifiCates) Method() string {
	return "GET"
}

// Params CertifiCates Params
func (this CertifiCates) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// RawJsonStr CertifiCates RawJsonStr
func (this CertifiCates) RawJsonStr() string {
	return ""
}

type SystemOauthTokenRsp struct {
	Data []struct {
		Serial_no           string `json:"serial_no"`
		Effective_time      string `json:"effective_time"`
		Expire_time         string `json:"expire_time"`
		Encrypt_certificate struct {
			Algorithm       string `json:"algorithm"`
			Nonce           string `json:"nonce"`
			Associated_data string `json:"associated_data"`
			Ciphertext      string `json:"ciphertext"`
		} `json:"encrypt_certificate"`
	} `json:"data"`
}

func (this SystemOauthTokenRsp) CertificateDecryption(apiv3key string) (interface{}, error) {
	if len(this.Data) < 1 {
		return nil, errors.New("没有找到平台密文证书")
	} else {
		cpinfo := this.Data[0]
		// 对编码密文进行base64解码
		decodeBytes, err := base64.StdEncoding.DecodeString(cpinfo.Encrypt_certificate.Ciphertext)
		if err != nil {
			return nil, err
		}

		c, err := aes.NewCipher([]byte(apiv3key))
		if err != nil {
			return nil, err
		}

		gcm, err := cipher.NewGCM(c)
		if err != nil {
			return nil, err
		}

		nonceSize := gcm.NonceSize()
		if len(decodeBytes) < nonceSize {
			return nil, errors.New("密文证书长度不够")
		}
		res := CertificateInfo{}
		res.Serial_no = cpinfo.Serial_no
		if cpinfo.Encrypt_certificate.Associated_data != "" {
			plaintext, err := gcm.Open(nil, []byte(cpinfo.Encrypt_certificate.Nonce), decodeBytes, []byte(cpinfo.Encrypt_certificate.Associated_data))
			if err != nil {
				return nil, err
			}
			res.Publickey = string(plaintext)
		} else {
			plaintext, err := gcm.Open(nil, []byte(cpinfo.Encrypt_certificate.Nonce), decodeBytes, nil)
			if err != nil {
				return nil, err
			}
			res.Publickey = string(plaintext)
		}
		return res, nil
	}
}

type CertificateInfo struct {
	Serial_no string `json:"serial_no"`
	Publickey string `json:"publickey"`
}
