package wxpayv3

import (
	"fmt"
	"testing"
)

func TestClient_QueryUserInfo(t *testing.T) {
	//client, err := New("1501889641", "1433BE83CE5A9D6022972F4B144A714C598CFADF", "apiclient_key.pem", "apiclient_cert.pem")
	//if err != nil {
	//	t.Log(fmt.Sprintf("V3客户端初始化失败,err:%+v", err))
	//	return
	//}

	client, err := New("1603541829", "7626C04E556AD527BD1844D9932FB54E91CAC378", "1603541829_key.pem", "1603541829_cert.pem")
	if err != nil {
		t.Log(fmt.Sprintf("V3客户端初始化失败,err:%+v", err))
		return
	}

	param := QueryUserInfo{}
	param.Organization_id = "O74B253Z38448e598c"
	param.Out_user_id = "945625"
	res, err := client.QueryUserInfo(param)
	if err != nil {
		t.Log(fmt.Sprintf("RESP:%+v", res))
		t.Log(fmt.Sprintf("失败,err:%+v", err))
		return
	}
	t.Log(fmt.Sprintf("result:%+v", res))

}

func TestClient_QueryContracts(t *testing.T) {

	client, err := New("1603541829", "7626C04E556AD527BD1844D9932FB54E91CAC378", "1603541829_key.pem", "1603541829_cert.pem")
	if err != nil {
		t.Log(fmt.Sprintf("V3客户端初始化失败,err:%+v", err))
		return
	}
	//client, err := New("1501889641", "1433BE83CE5A9D6022972F4B144A714C598CFADF", "apiclient_key.pem", "apiclient_cert.pem")
	//if err != nil {
	//	t.Log(fmt.Sprintf("V3客户端初始化失败,err:%+v", err))
	//	return
	//}
	param2 := QueryContracts{}
	param2.Contract_id = "CI4B86DFZ3a64e8415"
	param2.Appid = "wx23f78ccbdb9bfb6f"

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
