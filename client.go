package wxpayv3

import (
	"fmt"
	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/core"
	"net/http"
)

func NewClient(opts ...core.ClientOption) (core.Client, error) {
	settings := &core.DialSettings{}
	for _, opt := range opts {
		if err := opt.Join(settings); err != nil {
			return nil, fmt.Errorf("初始化客户端设置错误:%v", err)
		}
	}
	if err := settings.Validate(); err != nil {
		return nil, err
	}
	if settings.HttpClient == nil {
		settings.HttpClient = &http.Client{
			Timeout: constant.DefaultTimeout,
		}
	}

	client := &core.PayClient{
		MchId:               settings.MchId,
		ApiV3Key:            settings.ApiV3Key,
		ApiSerialNo:         settings.ApiSerialNo,
		ApiPrivateKey:       settings.ApiPrivateKey,
		ApiCertificate:      settings.ApiCertificate,
		PlatformSerialNo:    settings.PlatformSerialNo,
		PlatformCertificate: settings.PlatformCertificate,
		HttpClient:          settings.HttpClient,
	}

	return client, nil
}
