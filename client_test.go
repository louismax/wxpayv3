package wxpayv3

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	client,err := NewClient(
		InjectWechatPayParameterUseCertPath("1501889641","anzhixiaopay2019anzhixiaopay9876","apiclient_key.pem","apiclient_cert.pem"),
		)
	if err != nil{
		t.Log(err)
	}
	t.Log(client)
}