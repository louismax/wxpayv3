package wxpayv3

import (
	"fmt"
	"testing"
)

func TestClient_UploadImage(t *testing.T) {
	client, err := New("", "", "apiclient_key.pem", "apiclient_cert.pem")
	if err != nil {
		t.Log(fmt.Sprintf("V3客户端初始化失败,err:%+v", err))
		return
	}
	res, err := client.UploadImage("./微信图片.png")
	if err != nil {
		t.Log(fmt.Sprintf("RESP:%+v", res))
		t.Log(fmt.Sprintf("失败,err:%+v", err))
		return
	}

	fmt.Println(fmt.Sprintf("%+v", res))

}
