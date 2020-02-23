package wxpayv3

import (
	"encoding/json"
	"fmt"
)

const (
	Business_scene_type_K12         = 2 //K12项目
	Business_scene_type_Mess        = 3 //食堂场景
	Business_scene_type_Supermarket = 4 //超市场景
	Business_scene_type_Infirmary   = 5 //校医院
	Business_scene_type_Dev         = 6 //测试场景
)

type ReqCreatePayCredential struct {
	Pay_credential    string `json:"pay_credential"` //支付凭证
	Mchid             string `json:"mchid"`          //商户号
	Sub_mchid         string `json:"sub_mchid"`
	Appid             string `json:"appid"`
	Sub_appid         string `json:"sub_appid"`
	Amount            int64  `json:"amount"`
	Device_ip         string `json:"device_ip"`
	Mac               string `json:"mac"`
	Description       string `json:"description"`  //商品信息
	Attach            string `json:"attach"`       //商户附加信息
	Out_trade_no      string `json:"out_trade_no"` //商户订单号
	Business_scene_id int    `json:"business_scene_id"`
}

type CreatePayCredential struct {
	Pay_credential    string          `json:"pay_credential"`    //支付凭证
	Merchant_info     MerchantInfo    `json:"merchant_info"`     //商户信息
	Trade_amount_info TradeAmountInfo `json:"trade_amount_info"` //金额信息
	Scene_info        SceneInfo       `json:"scene_info"`        //支付场景信息
	Device_info       DeviceInfo      `json:"device_info"`       //设备信息
	Goods_tag         string          `json:"goods_tag"`         //优惠标记
	Description       string          `json:"description"`       //商品信息
	Attach            string          `json:"attach"`            //商户附加信息
	Out_trade_no      string          `json:"out_trade_no"`      //商户订单号
	Business_info     BusinessInfo    `json:"business_info"`     //业务信息
}

type MerchantInfo struct {
	Mchid     string `json:"mchid"`     //商户号
	Sub_mchid string `json:"sub_mchid"` //子商户号
	Appid     string `json:"appid"`     //商户公众号
	Sub_appid string `json:"sub_appid"` //子商户公众号
}

type TradeAmountInfo struct {
	Amount   int64  `json:"amount"`   //总金额
	Currency string `json:"currency"` //货币类型
}

type SceneInfo struct {
	Device_ip string `json:"device_ip"` //设备IP
}

type DeviceInfo struct {
	Mac string `json:"mac"` //设备mc地址
}

type BusinessInfo struct {
	Business_product_id int `json:"business_product_id"` //平台产品ID
	Business_scene_id   int `json:"business_scene_id"`   //平台场景ID
}

// APIUrl CreatePayCredential APIURL
func (this CreatePayCredential) APIUrl() string {
	return "/v3/offlinefacemch/paycredential"
}

// Method CreatePayCredential Method
func (this CreatePayCredential) Method() string {
	return "POST"
}

// Params CreateOrganization Params
func (this CreatePayCredential) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

func (this CreatePayCredential) RawJsonStr() string {
	jsons, errs := json.Marshal(this) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
		return "Error"
	}
	return string(jsons)
}

type RespCreatePayCredential struct {
	Merchant_info     MerchantInfo    `json:"merchant_info"`     //商户信息
	Payer_info        PayerInfo       `json:"payer_info"`        //支付用户信息
	Trade_amount_info TradeAmountInfo `json:"trade_amount_info"` //金额信息
	Promotion_list    []PromotionList `json:"promotion_list"`    //优惠信息
	Scene_info        SceneInfo       `json:"scene_info"`        //支付场景信息
	Core_payment_info CorePaymentInfo `json:"core_payment_info"` //支付信息
	Trade_type        string          `json:"trade_type"`        //交易类型
	Trade_state       string          `json:"trade_state"`       //交易状态
	Body              string          `json:"body"`              //商品信息
	Attach            string          `json:"attach"`            //商户附加信息
	Payment_time      string          `json:"payment_time"`      //支付成功时间
	Transaction_id    string          `json:"transaction_id"`    //微信订单号
	Out_trade_no      string          `json:"out_trade_no"`      //商户订单号
}

type PayerInfo struct {
	Openid     string `json:"openid"`     //公众号下的openid
	Sub_openid string `json:"sub_openid"` //子公众号下的openid
}

type PromotionList struct {
	Promotion_id        string          `json:"promotion_id"`        //优惠ID
	Name                string          `json:"name"`                //优惠名称
	Amount_info         TradeAmountInfo `json:"amount_info"`         //优惠金额
	Wxpay_contribute    int64           `json:"wxpay_contribute"`    //微信出资金额
	Merchant_contribute int64           `json:"merchant_contribute"` //商家出资金额
	Other_contribute    int64           `json:"other_contribute"`    //其他出资金额
}

type CorePaymentInfo struct {
	Bank_type   string `json:"bank_type"`   //付款银行
	Amount      int64  `json:"amount"`      //支付金额
	Currency    string `json:"currency"`    //支付币种
	Advance_pay bool   `json:"advance_pay"` //是否发生垫资
}
