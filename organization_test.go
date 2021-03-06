package wxpayv3

import (
	"fmt"
	"testing"
)

func TestClient_QueryOrganization(t *testing.T) {
	client, err := New("", "", "apiclient_key.pem", "apiclient_cert.pem")
	if err != nil {
		t.Log(fmt.Sprintf("V3客户端初始化失败,err:%+v", err))
		return
	}

	param := QueryOrganization{Organization_id: ""}

	res, err := client.QueryOrganization(param)
	if err != nil {
		t.Log(fmt.Sprintf("RESP:%+v", res))
		t.Log(fmt.Sprintf("失败,err:%+v", err))
		return
	}
	t.Log(fmt.Sprintf("result:%+v", res))
}
