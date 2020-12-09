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

func TestClient_Transactions(t *testing.T) {
	client, err := New("", "", "", "")
	if err != nil {
		t.Log(fmt.Sprintf("V3客户端初始化失败,err:%+v", err))
		return
	}

	param := ReqTransactions{
		Out_Trade_No:        "",
		Sp_Mchid:            "",
		Sub_Mchid:           "",
		Business_Product_ID: 2,
	}

	res, err := client.Transactions(param)
	if err != nil {
		t.Log(fmt.Sprintf("RESP:%+v", res))
		t.Log(fmt.Sprintf("失败,err:%+v", err))
		return
	}
	t.Log(fmt.Sprintf("result:%+v", res))

}

func TestClient_TransactionDeduction(t *testing.T) {
	client, err := New("", "", "", "")
	if err != nil {
		t.Log(fmt.Sprintf("err:%+v", err))
		return
	}

	param := ReqTransactionDeduction{
		AuthCode: "",
		SpAppid:  "",
		SubAppid: "",
		SpMchid:  "",
		SubMchid: "",
		Amount: struct {
			Total    int64  `json:"total"`
			Currency string `json:"currency"`
		}{},
		SceneInfo: struct {
			DeviceIp string `json:"device_ip"`
		}{},
		GoodsTag:    "",
		Description: "",
		Attach:      "",
		SettleInfo: struct {
			ProfitSharing bool `json:"profit_sharing"`
		}{},
		OutTradeNo: "",
		Business: struct {
			BusinessProductId int `json:"business_product_id"`
			BusinessSceneId   int `json:"business_scene_id"`
		}{},
	}
	param.Amount.Total = 1
	param.Amount.Currency = "CNY"
	param.SceneInfo.DeviceIp = "127.0.0.1"
	param.SettleInfo.ProfitSharing = false
	param.Business.BusinessProductId = Business_scene_type_K12
	param.Business.BusinessSceneId = Business_scene_type_Mess

	res, err := client.TransactionDeduction(param)
	if err != nil {
		t.Log(fmt.Sprintf("RESP:%+v", res))
		t.Log(fmt.Sprintf("失败,err:%+v", err))
		return
	}
	t.Log(fmt.Sprintf("result:%+v", res))
}
