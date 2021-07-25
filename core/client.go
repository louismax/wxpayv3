package core

import (
	"context"
	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/credentials"
	"net/http"
)

// PayClient 微信支付API v3 基础 Client
type PayClient struct {
	httpClient *http.Client
	credential Credential
	validator  Validator
	signer     Signer
	cipher     Cipher
}

func InitClientWithSettings(_ context.Context, settings *DialSettings) *PayClient {
	client := &PayClient{
		signer:     settings.Signer,
		validator:  settings.Validator,
		credential: &credentials.WechatPayCredentials{Signer: settings.Signer},
		httpClient: settings.HTTPClient,
		cipher:     settings.Cipher,
	}

	if client.httpClient == nil {
		client.httpClient = &http.Client{
			Timeout: constant.DefaultTimeout,
		}
	}
	return client
}

//func (c *PayClient) Authorization(httpMethod string, urlString string, body []byte) (string, error) {
//	return "", nil
//}


type Client interface {
	//Authorization(httpMethod string, urlString string, body []byte) (string, error)
}