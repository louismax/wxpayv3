package wxpayv3

import (
	"fmt"
	"github.com/louismax/wxpayv3/marketing"
	"testing"
)

func TestClient_SendCoupon(t *testing.T) {
	client, err := New("", "", "", "")
	if err != nil {
		t.Log(fmt.Sprintf("V3客户端初始化失败,err:%+v", err))
		return
	}

	param := marketing.LssueCoupons{}
	param.Stock_id = ""
	param.Openid = ""
	param.Out_request_no = GetGUID()
	param.Appid = ""
	param.Stock_creator_mchid = ""
	//param.Coupon_value = 50
	//param.Coupon_minimum = 51

	res, err := client.SendCoupon(param)
	if err != nil {
		t.Log(fmt.Sprintf("RESP:%+v", res))
		t.Log(fmt.Sprintf("失败,err:%+v", err))
		return
	}
	t.Log(fmt.Sprintf("result:%+v", res))

}
