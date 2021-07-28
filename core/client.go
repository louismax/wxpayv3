package core

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/custom"
	"io/ioutil"
	"net/http"
)

type Client interface {
	Authorization(httpMethod string, urlString string, body []byte) (string, error) // Authorization 获取签名Authorization，由认证类型和签名信息组成
	Certificate() (*custom.CertificateResp, error)                                  //获取平台证书
	SetClientPlatformCert(certificateStr string) error                              //设置平台证书

	QuerySettlementAccount(subMchid string) (*custom.SettlementAccount, error) //获取结算账户
}

type PayClient struct {
	MchId               string            // 商户号
	ApiV3Key            string            // apiV3密钥
	ApiSerialNo         string            // API证书序列号
	ApiPrivateKey       *rsa.PrivateKey   // API私钥
	ApiCertificate      *x509.Certificate // API证书
	PlatformSerialNo    string            // 平台证书序列号
	PlatformCertificate *x509.Certificate // 平台证书
	HttpClient          *http.Client
}

func (c *PayClient) doRequest(requestData interface{}, url string, httpMethod string) ([]byte, error) {
	var data []byte
	if requestData != nil {
		var err error
		data, err = json.Marshal(requestData)
		if err != nil {
			return nil, err
		}
	}
	authorization, err := c.Authorization(httpMethod, url, data)
	if err != nil {
		return nil, err
	}
	// 重试3次，避免因网络原因导致失败
	retryTimes := 3
	var resp *http.Response
	for i := 0; i < retryTimes; i++ {
		resp, err = SimpleRequest(c.HttpClient, url, httpMethod, authorization, data)
		if err != nil {
			continue
		}
		break
	}
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = c.VerifyResponse(resp.StatusCode, &resp.Header, body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *PayClient) Decrypt(algorithm string, cipherText string, associatedData string, nonce string) ([]byte, error) {
	// 默认使用AEAD_AES_256_GCM
	switch algorithm {
	default:
		fallthrough
	case constant.AlgorithmAEADAES256GCM:
		decodedCipherText, _ := base64.StdEncoding.DecodeString(cipherText)

		block, err := aes.NewCipher([]byte(c.ApiV3Key))
		if err != nil {
			return nil, err
		}

		aesGcm, err := cipher.NewGCM(block)
		if err != nil {
			return nil, err
		}

		plaintext, err := aesGcm.Open(nil, []byte(nonce), decodedCipherText, []byte(associatedData))
		if err != nil {
			return nil, err
		}
		return plaintext, nil
	}
}
