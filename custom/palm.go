package custom

type RespPalmServicePreBind struct {
	PrebindToken string `json:"prebind_token"`
}

type RespPalmServiceQueryBind struct {
	BindState        string `json:"bind_state"`
	ServiceId        string `json:"service_id"`
	MerchantTicketId string `json:"merchant_ticket_id"`
}
