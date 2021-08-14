package custom

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
