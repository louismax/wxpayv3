package core

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client interface {
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
