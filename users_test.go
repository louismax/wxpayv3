package wxpayv3

import (
	"fmt"
	"testing"
)

func TestClient_QueryUserInfo(t *testing.T) {
	client, err := New("15018xxxxxxx", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "apiclient_key.pem", "apiclient_cert.pem")
	if err != nil {
		t.Log(fmt.Sprintf("V3客户端初始化失败,err:%+v", err))
		return
	}

	param := QueryUserInfo{}
	param.Organization_id = "O36xxxxxxxxxxxxx"
	param.Out_user_id = "59xxxxx"
	res, err := client.QueryUserInfo(param)
	if err != nil {
		t.Log(fmt.Sprintf("RESP:%+v", res))
		t.Log(fmt.Sprintf("失败,err:%+v", err))
		return
	}
	t.Log(fmt.Sprintf("result:%+v", res))

}

func TestClient_QueryContracts(t *testing.T) {
	client, err := New("15018xxxxxxx", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "apiclient_key.pem", "apiclient_cert.pem")
	if err != nil {
		t.Log(fmt.Sprintf("V3客户端初始化失败,err:%+v", err))
		return
	}
	param2 := QueryContracts{}
	param2.Contract_id = "CIxxxxxxxxxxxxxxx"
	param2.Appid = "wx9xxxxxxxxxxxxx"

	res2, err := client.QueryContracts(param2)
	if err != nil {
		t.Log(fmt.Sprintf("RESP:%+v", res2))
		t.Log(fmt.Sprintf("失败,err:%+v", err))
		return
	}
	t.Log(fmt.Sprintf("result:%+v", res2))
}
