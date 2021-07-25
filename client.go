package wxpayv3

import (
	"context"
	"fmt"
	"github.com/louismax/wxpayv3/core"
)

//type ClientConfig struct {
//	MchId      string //商户号
//	ApiV3Key   string
//	ApiCert    *core.ApiCert //商户证书
//	PlatCert   *core.PlatformCert //平台证书
//	HttpClient *http.Client
//}


func NewClient(ctx context.Context, opts ...ClientOption) (core.Client, error) {
	settings := core.DialSettings{}
	for _, opt := range opts {
		if err := opt.Apply(&settings); err != nil {
			return nil, fmt.Errorf("初始化客户端设置错误:%v", err)
		}
	}
	if err := settings.Validate(); err != nil {
		return nil, fmt.Errorf("初始化客户端设置错误:%v", err)
	}
	client := core.InitClientWithSettings(ctx, &settings)
	return client, nil
}


