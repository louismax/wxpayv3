package custom

type ReqPaymentQueryOrder struct {
	Amount struct {
		Currency      string `json:"currency"`
		PayerCurrency string `json:"payer_currency"`
		PayerTotal    int    `json:"payer_total"`
		Total         int    `json:"total"`
	} `json:"amount"`
	Appid      string `json:"appid"`
	Attach     string `json:"attach"`
	BankType   string `json:"bank_type"`
	Mchid      string `json:"mchid"`
	OutTradeNo string `json:"out_trade_no"`
	Payer      struct {
		Openid string `json:"openid"`
	} `json:"payer"`
	PromotionDetail []interface{} `json:"promotion_detail"`
	SuccessTime     string        `json:"success_time"`
	TradeState      string        `json:"trade_state"`
	TradeStateDesc  string        `json:"trade_state_desc"`
	TradeType       string        `json:"trade_type"`
	TransactionId   string        `json:"transaction_id"`
}

type ReqPaymentRefund struct {
	SubMchid      string                `json:"sub_mchid"`
	TransactionId string                `json:"transaction_id,omitempty"`
	OutTradeNo    string                `json:"out_trade_no,omitempty"`
	OutRefundNo   string                `json:"out_refund_no"`
	Reason        string                `json:"reason,omitempty"`
	NotifyUrl     string                `json:"notify_url,omitempty"`
	FundsAccount  string                `json:"funds_account,omitempty"`
	Amount        PaymentRefundAmount   `json:"amount"`
	GoodsDetail   *[]PaymentGoodsDetail `json:"goods_detail,omitempty"`
}

type PaymentRefundAmount struct {
	Refund   int                  `json:"refund"`
	From     *[]PaymentRefundFrom `json:"from,omitempty"`
	Total    int                  `json:"total"`
	Currency string               `json:"currency"`
}

type PaymentRefundFrom struct {
	Account string `json:"account"`
	Amount  string `json:"amount"`
}

type PaymentGoodsDetail struct {
	MerchantGoodsId  string `json:"merchant_goods_id"`
	WechatpayGoodsId string `json:"wechatpay_goods_id,omitempty"`
	GoodsName        string `json:"goods_name,omitempty"`
	UnitPrice        int    `json:"unit_price"`
	RefundAmount     int    `json:"refund_amount"`
	RefundQuantity   int    `json:"refund_quantity"`
}

type RespPaymentRefund struct {
	RefundId            string `json:"refund_id"`
	OutRefundNo         string `json:"out_refund_no"`
	TransactionId       string `json:"transaction_id"`
	OutTradeNo          string `json:"out_trade_no"`
	Channel             string `json:"channel"`
	UserReceivedAccount string `json:"user_received_account"`
	SuccessTime         string `json:"success_time"`
	CreateTime          string `json:"create_time"`
	Status              string `json:"status"`
	FundsAccount        string `json:"funds_account"`
	Amount              struct {
		Total  int `json:"total"`
		Refund int `json:"refund"`
		From   []struct {
			Account string `json:"account"`
			Amount  int    `json:"amount"`
		} `json:"from"`
		PayerTotal       int    `json:"payer_total"`
		PayerRefund      int    `json:"payer_refund"`
		SettlementRefund int    `json:"settlement_refund"`
		SettlementTotal  int    `json:"settlement_total"`
		DiscountRefund   int    `json:"discount_refund"`
		Currency         string `json:"currency"`
	} `json:"amount"`
	PromotionDetail []struct {
		PromotionId  string `json:"promotion_id"`
		Scope        string `json:"scope"`
		Type         string `json:"type"`
		Amount       int    `json:"amount"`
		RefundAmount int    `json:"refund_amount"`
		GoodsDetail  []struct {
			MerchantGoodsId  string `json:"merchant_goods_id"`
			WechatpayGoodsId string `json:"wechatpay_goods_id"`
			GoodsName        string `json:"goods_name"`
			UnitPrice        int    `json:"unit_price"`
			RefundAmount     int    `json:"refund_amount"`
			RefundQuantity   int    `json:"refund_quantity"`
		} `json:"goods_detail"`
	} `json:"promotion_detail"`
}

type RespApplyTransactionBill struct {
	DownloadURL string `json:"download_url"`
	HashType    string `json:"hash_type"`
	HashValue   string `json:"hash_value"`
}

type ReqJSAPIOrdersForPartner struct {
	SpAppid     string  `json:"sp_appid"`
	SpMchid     string  `json:"sp_mchid"`
	SubAppid    string  `json:"sub_appid"`
	SubMchid    string  `json:"sub_mchid"`
	Description string  `json:"description"`
	OutTradeNo  string  `json:"out_trade_no"`
	TimeExpire  *string `json:"time_expire,omitempty"`
	Attach      string  `json:"attach"`
	NotifyUrl   string  `json:"notify_url"`
	GoodsTag    string  `json:"goods_tag"`
	Amount      struct {
		Total    int    `json:"total"`
		Currency string `json:"currency"`
	} `json:"amount"`
	Payer struct {
		SpOpenid  string `json:"sp_openid"`
		SubOpenid string `json:"sub_openid"`
	} `json:"payer"`
	Detail *struct {
		CostPrice   int    `json:"cost_price"`
		InvoiceId   string `json:"invoice_id"`
		GoodsDetail *[]struct {
			MerchantGoodsId  string `json:"merchant_goods_id"`
			WechatPayGoodsId string `json:"wechatpay_goods_id"`
			GoodsName        string `json:"goods_name"`
			Quantity         int    `json:"quantity"`
			UnitPrice        int    `json:"unit_price"`
		} `json:"goods_detail,omitempty"`
	} `json:"detail,omitempty"`
	SceneInfo *struct {
		PayerClientIp string `json:"payer_client_ip"`
		DeviceId      string `json:"device_id"`
		StoreInfo     struct {
			Id       string `json:"id"`
			Name     string `json:"name"`
			AreaCode string `json:"area_code"`
			Address  string `json:"address"`
		} `json:"store_info"`
	} `json:"scene_info,omitempty"`
	SettleInfo struct {
		ProfitSharing bool `json:"profit_sharing"`
	} `json:"settle_info"`
}

type ReqJSAPIOrders struct {
	Appid       string  `json:"appid"`
	Mchid       string  `json:"mchid"`
	Description string  `json:"description"`
	OutTradeNo  string  `json:"out_trade_no"`
	TimeExpire  *string `json:"time_expire,omitempty"`
	Attach      string  `json:"attach"`
	NotifyUrl   string  `json:"notify_url"`
	GoodsTag    string  `json:"goods_tag"`
	Amount      struct {
		Total    int    `json:"total"`
		Currency string `json:"currency"`
	} `json:"amount"`
	Payer struct {
		Openid string `json:"openid"`
	} `json:"payer"`
	Detail *struct {
		CostPrice   int    `json:"cost_price"`
		InvoiceId   string `json:"invoice_id"`
		GoodsDetail *[]struct {
			MerchantGoodsId  string `json:"merchant_goods_id"`
			WechatPayGoodsId string `json:"wechatpay_goods_id"`
			GoodsName        string `json:"goods_name"`
			Quantity         int    `json:"quantity"`
			UnitPrice        int    `json:"unit_price"`
		} `json:"goods_detail,omitempty"`
	} `json:"detail,omitempty"`
	SceneInfo *struct {
		PayerClientIp string `json:"payer_client_ip"`
		DeviceId      string `json:"device_id"`
		StoreInfo     struct {
			Id       string `json:"id"`
			Name     string `json:"name"`
			AreaCode string `json:"area_code"`
			Address  string `json:"address"`
		} `json:"store_info"`
	} `json:"scene_info,omitempty"`
	SettleInfo struct {
		ProfitSharing bool `json:"profit_sharing"`
	} `json:"settle_info"`
}

type RespJSAPIOrders struct {
	PrepayId string `json:"prepay_id"`
}
