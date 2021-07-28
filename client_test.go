package wxpayv3

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient(
		InjectWechatPayParameterUseCertPath("1607549129", "b5E016760S289934074655Q752T81ZS8", "apiclient_key.pem", "apiclient_cert.pem"),
	)
	if err != nil {
		t.Log(err)
		return
	}
	//resp, err := client.Certificate()
	//if err != nil {
	//	t.Log(err)
	//	return
	//}
	//
	//for _, v := range resp.Data {
	//	t.Logf("%+v", *v)
	//	t.Logf("%+v", *v.EncryptCertificate)
	//}

	resp, err := client.QuerySettlementAccount("1609337198")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(resp)

}
