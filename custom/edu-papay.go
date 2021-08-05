package custom

type ReqEduPaPayPresign struct {
	AppId           string `json:"appid"`
	SubMchId        string `json:"sub_mchid"`
	SubAppId        string `json:"sub_appid"`
	OpenId          string `json:"openid"`
	SubOpenId       string `json:"sub_openid"`
	PlanId          string `json:"plan_id"`
	UserId          string `json:"user_id"`
	PeriodStartDate string `json:"period_start_date"`
	TradeScene      string `json:"trade_scene"`
}

type RespEduPaPayPresign struct {
	PresignToken string `json:"presign_token"`
}

type RespEduPaPayContractQueryList struct {
	TotalCount int                         `json:"total_count"`
	Offset     int                         `json:"offset"`
	Limit      int                         `json:"limit"`
	Data       []RespEduPaPayContractQuery `json:"data"`
}

type RespEduPaPayContractQuery struct {
	SpMchId             string              `json:"sp_mchid"`
	AppId               string              `json:"appid"`
	SubMchId            string              `json:"sub_mchid"`
	SubAppId            string              `json:"sub_appid"`
	OpenId              string              `json:"openid"`
	SubOpenId           string              `json:"sub_openid"`
	PlanId              string              `json:"plan_id"`
	ContractInformation ContractInformation `json:"contract_information"`
}

type ContractInformation struct {
	ContractId     string `json:"contract_id"`
	ContractStatus string `json:"contract_status"`
	CreateTime     string `json:"create_time"`
}

type ReqSendEduPaPayNotifications struct {
	AppId    string `json:"appid"`
	SubMchId string `json:"sub_mchid"`
	SubAppId string `json:"sub_appid"`
}

type ReqEduPaPayTransactions struct {
	Appid             string                    `json:"appid"`
	SubMchId          string                    `json:"sub_mchid"`
	SubAppId          string                    `json:"sub_appid"`
	Body              string                    `json:"body"`
	Attach            string                    `json:"attach"`
	OutTradeNo        string                    `json:"out_trade_no"`
	GoodsTag          string                    `json:"goods_tag"`
	NotifyUrl         string                    `json:"notify_url"`
	ContractId        string                    `json:"contract_id"`
	TradeScene        string                    `json:"trade_scene"`
	Amount            EduPaPayAmount            `json:"amount"`
	DeviceInformation EduPaPayDeviceInformation `json:"device_information"`
}
type EduPaPayAmount struct {
	Total    int64  `json:"total"`
	Currency string `json:"currency"`
}

type EduPaPayDeviceInformation struct {
	DeviceId string `json:"device_id"`
	DeviceIp string `json:"device_ip"`
}

type RespEduPaPayQueryOrder struct {
	SpMchId               string `json:"sp_mchid"`
	AppId                 string `json:"appid"`
	SubMchId              string `json:"sub_mchid"`
	SubAppId              string `json:"sub_appid"`
	OutTradeNo            string `json:"out_trade_no"`
	TransactionId         string `json:"transaction_id"`
	Attach                string `json:"attach"`
	BankType              string `json:"bank_type"`
	SuccessTime           string `json:"success_time"`
	TradeState            string `json:"trade_state"`
	TradeStateDescription string `json:"trade_state_description"`
	Payer                 struct {
		OpenId    string `json:"openid"`
		SubOpenId string `json:"sub_openid"`
	} `json:"payer"`
	Amount struct {
		Total         int64  `json:"total"`
		PayerTotal    int64  `json:"payer_total"`
		DiscountTotal int64  `json:"discount_total"`
		Currency      string `json:"currency"`
	} `json:"amount"`
	DeviceInformation struct {
		DeviceId string `json:"device_id"`
		DeviceIp string `json:"device_ip"`
	} `json:"device_information"`
	PromotionDetail []struct {
		CouponId            string `json:"coupon_id"`
		Name                string `json:"name"`
		Scope               string `json:"scope"`
		Type                string `json:"type"`
		Amount              int64  `json:"amount"`
		StockId             string `json:"stock_id"`
		WechatpayContribute int64  `json:"wechatpay_contribute"`
		MerchantContribute  int64  `json:"merchant_contribute"`
		OtherContribute     int64  `json:"other_contribute"`
	} `json:"promotion_detail"`
}
