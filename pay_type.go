package wxpayv3

import (
	"encoding/json"
	"fmt"
)

const (
	Business_scene_type_K12 = 2  //K12项目
	Business_scene_type_QY  = 11 //企业团餐

	Business_scene_type_Mess        = 3 //食堂场景
	Business_scene_type_Supermarket = 4 //超市场景
	Business_scene_type_Infirmary   = 5 //校医院
	Business_scene_type_Dev         = 6 //测试场景

	Business_scene_type_K12Test = 124 //团餐测试
	Business_scene_type_QYBus   = 125 //企业食堂
	Business_scene_type_TXBus   = 126 //腾讯食堂

)

const (
	AUTH      = "AUTH"
	NOT_DEBT  = "NOT_DEBT"  //无垫资
	DEBT      = "DEBT"      //垫资支付
	REPAYMENT = "REPAYMENT" //已还款
	CNY       = "CNY"
)

type ReqCreatePayCredential struct {
	PayCredential   string `json:"pay_credential"` //支付凭证
	Mchid           string `json:"mchid"`          //商户号
	SubMchid        string `json:"sub_mchid"`
	Appid           string `json:"appid"`
	SubAppid        string `json:"sub_appid"`
	Amount          int64  `json:"amount"`
	DeviceIp        string `json:"device_ip"`
	Mac             string `json:"mac"`
	Description     string `json:"description"`  //商品信息
	Attach          string `json:"attach"`       //商户附加信息
	OutTradeNo      string `json:"out_trade_no"` //商户订单号
	BusinessSceneId int    `json:"business_scene_id"`
	GoodsTag        string `json:"goods_tag"` //优惠标记
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
	Trade_state_desc  string          `json:"trade_state_desc"`  //交易状态描述
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
	Bank_type         string `json:"bank_type"` //付款银行
	Core_payment_info struct {
		Amount   int64  `json:"amount"`   //支付金额
		Currency string `json:"currency"` //支付币种
	} `json:"core_payment_info"`
}

//ReqTransactions K12查单
type ReqTransactions struct {
	Out_Trade_No        string `json:"out_trade_no"`
	Sp_Mchid            string `json:"sp_mchid"`
	Sub_Mchid           string `json:"sub_mchid"`
	Business_Product_ID int    `json:"business_product_id"`
}

// APIUrl ReqTransactions APIURL
func (this ReqTransactions) APIUrl() string {
	return fmt.Sprintf("/v3/offlineface/transactions/out-trade-no/%s?sp_mchid=%s&sub_mchid=%s&business_product_id=%d", this.Out_Trade_No, this.Sp_Mchid, this.Sub_Mchid, this.Business_Product_ID)
}

// Method ReqTransactions Method
func (this ReqTransactions) Method() string {
	return "GET"
}

// Params ReqTransactions Params
func (this ReqTransactions) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

func (this ReqTransactions) RawJsonStr() string {
	return ""
}

type ReqTransactionDeduction struct {
	AuthCode string `json:"auth_code"` //支付凭证
	SpAppid  string `json:"sp_appid"`
	SubAppid string `json:"sub_appid"`
	SpMchid  string `json:"sp_mchid"`  //商户号
	SubMchid string `json:"sub_mchid"` //子商户号
	Amount   struct {
		Total    int64  `json:"total"`    //总金额
		Currency string `json:"currency"` //货币类型
	} `json:"amount"` //金额信息
	SceneInfo struct {
		DeviceIp string `json:"device_ip"` //设备IP
	} `json:"scene_info"` //支付场景信息
	GoodsTag    string `json:"goods_tag"`   //优惠标记
	Description string `json:"description"` //商品信息
	Attach      string `json:"attach"`      //商户附加信息
	SettleInfo  struct {
		ProfitSharing bool `json:"profit_sharing"` //是否支持分账
	} `json:"settle_info"` //结算信息
	OutTradeNo string `json:"out_trade_no"` // 商户单号
	Business   struct {
		BusinessProductId int `json:"business_product_id"` //平台产品ID
		BusinessSceneId   int `json:"business_scene_id"`   //平台场景ID
	} `json:"business"` //业务信息
}

// APIUrl ReqTransactionDeduction APIURL
func (this ReqTransactionDeduction) APIUrl() string {
	return "/v3/offlineface/transactions"
}

// Method ReqTransactionDeduction Method
func (this ReqTransactionDeduction) Method() string {
	return "POST"
}

// Params ReqTransactionDeduction Params
func (this ReqTransactionDeduction) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

func (this ReqTransactionDeduction) RawJsonStr() string {
	jsons, errs := json.Marshal(this)
	if errs != nil {
		fmt.Println(errs.Error())
		return "Error"
	}
	return string(jsons)
}

type RespTransactionDeduction struct {
	SpAppid  string `json:"sp_appid"`
	SubAppid string `json:"sub_appid"`
	SpMchid  string `json:"sp_mchid"`  //商户号
	SubMchid string `json:"sub_mchid"` //子商户号
	Payer    struct {
		SPOpenid  string `json:"sp_openid"`  //公众号下的openid
		SubOpenid string `json:"sub_openid"` //子公众号下的openid
	} `json:"payer"`
	Amount struct {
		Total       int64  `json:"total"`        //订单金额
		PayTotal    int64  `json:"pay_total"`    //用户支付金额
		Currency    string `json:"currency"`     //货币类型
		PayCurrency string `json:"pay_currency"` //用户支付货币类型
	} `json:"amount"`
	PromotionDetail []PromotionList `json:"promotion_detail"` //优惠信息
	SceneInfo       struct {
		DeviceIp string `json:"device_ip"` //设备IP
	} `json:"scene_info"` //支付场景信息
	BankType               string `json:"bank_type"`                //付款银行
	TradeType              string `json:"trade_type"`               //交易类型
	TradeState             string `json:"trade_state"`              //交易状态
	TradeStateDescription  string `json:"trade_state_description"`  //交易描述
	DebtState              string `json:"debt_state"`               //欠款状态
	Description            string `json:"description"`              //商品信息
	Attach                 string `json:"attach"`                   //商户附加信息
	SuccessTime            string `json:"success_time"`             //支付成功时间
	TransactionId          string `json:"transaction_id"`           //微信订单号
	RepaymentTransactionId string `json:"repayment_transaction_id"` //还款微信单号
	OutTradeNo             string `json:"out_trade_no"`             //商户单号
}
