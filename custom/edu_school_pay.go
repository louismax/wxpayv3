package custom

type ReqEduSchoolPayPreSign struct {
	AppId                    string `json:"appid"`
	OpenId                   string `json:"openid"`
	PlanId                   string `json:"plan_id"`
	UserId                   string `json:"user_id"`
	SchoolId                 string `json:"school_id"`
	OutContractCode          string `json:"out_contract_code"`
	ContractMode             string `json:"contract_mode"`
	DowngradeDefaultContract bool   `json:"downgrade_default_contract"`
	Identity                 *struct {
		RealName       string `json:"real_name"`
		CredentialType string `json:"credential_type"`
		IdCardNumber   string `json:"id_card_number"`
	} `json:"identity,omitempty"`
	BankCard *struct {
		BankCardNo string `json:"bank_card_no"`
		ValidThru  string `json:"valid_thru"`
		Phone      string `json:"phone"`
		BankType   string `json:"bank_type"`
	} `json:"bank_card,,omitempty"`
}

type RespEduSchoolPayPreSign struct {
	PreSignToken string `json:"presign_token"`
}

type RespEduSchoolPayContractQuery struct {
	ContractId      string `json:"contract_id"`
	MchId           string `json:"mchid"`
	AppId           string `json:"appid"`
	OpenId          string `json:"openid"`
	PlanId          string `json:"plan_id"`
	ContractStatus  string `json:"contract_status"`
	CreateTime      string `json:"create_time"`
	OutContractCode string `json:"out_contract_code"`
}

type RespEduSchoolPayContractQueryPage struct {
	Data       []RespEduSchoolPayContractQuery `json:"data"`
	Offset     int                             `json:"offset"`
	Limit      int                             `json:"limit"`
	TotalCount int                             `json:"total_count"`
	Links      struct {
		Next string `json:"next"`
		Prev string `json:"prev"`
		Self string `json:"self"`
	} `json:"links"`
}

type ReqEduSchoolPayTransactions struct {
	Appid       string                 `json:"appid"`
	SubMchId    string                 `json:"sub_mchid"`
	SubAppId    string                 `json:"sub_appid"`
	Description string                 `json:"description"`
	Attach      string                 `json:"attach"`
	OutTradeNo  string                 `json:"out_trade_no"`
	GoodsTag    string                 `json:"goods_tag"`
	ContractId  string                 `json:"contract_id"`
	UserId      string                 `json:"user_id"`
	Amount      EduSchoolPayAmount     `json:"amount"`
	SceneInfo   EduSchoolPaySceneInfo  `json:"scene_info"`
	DeviceInfo  EduSchoolPayDeviceInfo `json:"device_info"`
	SettleInfo  EduSchoolPaySettleInfo `json:"settle_info"`
}

type EduSchoolPayAmount struct {
	Total    int64  `json:"total"`
	Currency string `json:"currency"`
}

type EduSchoolPaySceneInfo struct {
	StartTime string `json:"start_time"`
	SchoolId  string `json:"school_id"`
	SceneType string `json:"scene_type"`
}

type EduSchoolPayDeviceInfo struct {
	DeviceId string `json:"device_id"`
	DeviceIp string `json:"device_ip"`
}

type EduSchoolPaySettleInfo struct {
	ProfitSharing bool `json:"profit_sharing"`
}

type RespEduSchoolPayTransactions struct {
	MchId           string                      `json:"mchid"` //商户号
	Appid           string                      `json:"appid"`
	SubMchId        string                      `json:"sub_mchid"` //子商户号
	SubAppid        string                      `json:"sub_appid"`
	OutTradeNo      string                      `json:"out_trade_no"`     //商户单号
	TransactionId   string                      `json:"transaction_id"`   //微信订单号
	TradeType       string                      `json:"trade_type"`       //交易类型
	TradeState      string                      `json:"trade_state"`      //交易状态
	TradeStateDesc  string                      `json:"trade_state_desc"` //交易描述
	BankType        string                      `json:"bank_type"`        //付款银行
	Attach          string                      `json:"attach"`           //商户附加信息
	SuccessTime     string                      `json:"success_time"`     //支付成功时间
	Payer           EduSchoolPayPayer           `json:"payer"`
	Amount          RespEduSchoolPayAmount      `json:"amount"`
	DeviceInfo      EduSchoolPayDeviceInfo      `json:"device_info"`
	PromotionDetail []EduSchoolPayPromotionList `json:"promotion_detail"` //优惠信息

}
type EduSchoolPayPayer struct {
	Openid    string `json:"openid"`     //公众号下的openid
	SubOpenid string `json:"sub_openid"` //子公众号下的openid
}

type RespEduSchoolPayAmount struct {
	Total         int64  `json:"total"`       //订单金额
	PayerTotal    int64  `json:"payer_total"` //用户支付金额
	DiscountTotal int64  `json:"discount_total"`
	Currency      string `json:"currency"` //货币类型
}

type EduSchoolPayPromotionList struct {
	CouponId            string `json:"coupon_id"` //优惠ID
	Name                string `json:"name"`      //优惠名称
	Scope               string `json:"scope"`
	Type                string `json:"type"`
	Amount              int    `json:"amount"`
	StockId             string `json:"stock_id"`
	WechatPayContribute int64  `json:"wechatpay_contribute"` //微信出资金额
	MerchantContribute  int64  `json:"merchant_contribute"`  //商家出资金额
	OtherContribute     int64  `json:"other_contribute"`     //其他出资金额
}
