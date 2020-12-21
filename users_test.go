package wxpayv3

import (
	"fmt"
	"testing"
)

func TestClient_QueryUserInfo(t *testing.T) {
	client, err := New("", "", "", "")
	if err != nil {
		t.Log(fmt.Sprintf("err:%+v", err))
		return
	}

	param := QueryUserInfo{}
	param.Organization_id = ""
	param.Out_user_id = ""
	res, err := client.QueryUserInfo(param)
	if err != nil {
		t.Log(fmt.Sprintf("RESP:%+v", res))
		t.Log(fmt.Sprintf("失败,err:%+v", err))
		return
	}
	t.Log(fmt.Sprintf("result:%+v", res))

}

func TestClient_QueryContracts(t *testing.T) {
	client, err := New("", "", "", "")
	if err != nil {
		t.Log(fmt.Sprintf("err:%+v", err))
		return
	}
	param2 := QueryContracts{}
	param2.Contract_id = ""
	param2.Appid = ""

	res2, err := client.QueryContracts(param2)
	if err != nil {
		t.Log(fmt.Sprintf("RESP:%+v", res2))
		t.Log(fmt.Sprintf("失败,err:%+v", err))
		return
	}
	t.Log(fmt.Sprintf("result:%+v", res2))
}

func TestClient_QueryHeavyWeight(t *testing.T) {
	client, err := New("", "", "", "")
	if err != nil {
		t.Log(fmt.Sprintf("V3客户端初始化失败,err:%+v", err))
		return
	}
	param2 := QueryHeavyWeight{}
	param2.Organization_id = ""

	res2, err := client.QueryHeavyWeight(param2)
	if err != nil {
		t.Log(fmt.Sprintf("RESP:%+v", res2))
		t.Log(fmt.Sprintf("失败,err:%+v", err))
		return
	}
	t.Log(fmt.Sprintf("result:%+v", res2))
}
