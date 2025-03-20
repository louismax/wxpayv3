package wxpayv3

import (
	"github.com/louismax/wxpayv3/core"
	"testing"
)

func TestNewClient1(t *testing.T) {
	client, err := NewClient(
		InjectWxPayMchParamExtra("MCH_ID", "API_V3_KEY", "private_key.pem file path", "private_cert.pem file path"),
	)
	if err != nil {
		t.Log(err)
		return
	}

	//获取微信平台证书
	resp, err := client.GetCertificate()
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(resp)
}

func TestNewClient2(t *testing.T) {
	pk, err := core.LoadPrivateKey(`-----BEGIN PRIVATE KEY-----
XXXX商户私钥文本内容...
-----END PRIVATE KEY-----`)
	if err != nil {
		t.Log(err)
		return
	}
	client, err := NewClient(
		InjectWxPayMchParamFull("1607549129", core.ApiCert{
			ApiSerialNo:   "ApiSerialNo",
			ApiPrivateKey: pk,
		}, "API_V3_KEY"),
	)
	if err != nil {
		t.Log(err)
		return
	}

	resp, err := client.PaymentQueryOrderByOutTradeNo("OUT_TRADE_NO", "MCH_ID", "SUB_MCH_ID")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(resp)
}

func TestLoadPlatformCertStr(t *testing.T) {
	client, err := NewClient(
		InjectWxPayMchParamExtra("MCH_ID", "API_V3_KEY", "private_key.pem file path", "private_cert.pem file path"),
		InjectWxPayPlatformCert([]string{`-----BEGIN CERTIFICATE-----
平台证书文本内容...
-----END CERTIFICATE-----`}),
	)
	if err != nil {
		t.Log(err)
		return
	}

	resp, err := client.PaymentQueryOrderByOutTradeNo("OUT_TRADE_NO", "MCH_ID", "SUB_MCH_ID")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(resp)

}

func TestLoadPlatformCertFile(t *testing.T) {
	client, err := NewClient(
		InjectWxPayMchParamExtra("MCH_ID", "API_V3_KEY", "private_key.pem file path", "private_cert.pem file path"),
		InjectWxPayPlatformCertUseCertPath([]string{"pub_cert.pem  file path"}), //使用文件路径注入平台证书
	)
	if err != nil {
		t.Log(err)
		return
	}

	resp, err := client.PaymentQueryOrderByOutTradeNo("OUT_TRADE_NO", "MCH_ID", "SUB_MCH_ID")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(resp)
}

func TestWxPayPlatformPubKey(t *testing.T) {
	client, err := NewClient(
		InjectWxPayMchParamExtra("MCH_ID", "API_V3_KEY", "private_key.pem file path", "private_cert.pem file path"),
		InjectWxPayPlatformPubKey("PUB_KEY_ID_0000000000000000000000000000001", `-----BEGIN PUBLIC KEY-----
平台公钥文本内容...
-----END PUBLIC KEY-----
`),
	)
	if err != nil {
		t.Log(err)
		return
	}

	resp, err := client.PaymentQueryOrderByOutTradeNo("OUT_TRADE_NO", "MCH_ID", "SUB_MCH_ID")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(resp)
}

func TestWxPayPlatformPubKeyFile(t *testing.T) {
	client, err := NewClient(
		InjectWxPayMchParamExtra("MCH_ID", "API_V3_KEY", "private_key.pem file path", "private_cert.pem file path"),
		InjectWxPayPlatformPubKeyUsePath("PUB_KEY_ID_0000000000000000000000000000001", "./pub_key.pem"), //使用平台公钥本地文件
	)
	if err != nil {
		t.Log(err)
		return
	}

	resp, err := client.PaymentQueryOrderByOutTradeNo("OUT_TRADE_NO", "MCH_ID", "SUB_MCH_ID")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(resp)
}
