package custom

type ReqSmartGuideRegister struct {
	SubMchId    string `json:"sub_mchid"`
	CorpId      string `json:"corpid"`
	StoreId     int    `json:"store_id"`
	UserId      string `json:"userid"`
	Name        string `json:"name"`
	Mobile      string `json:"mobile"`
	QrCode      string `json:"qr_code"`
	Avatar      string `json:"avatar"`
	GroupQrcode string `json:"group_qrcode,omitempty"`
}

type RespSmartGuideRegister struct {
	GuideId string `json:"guide_id"`
}

type ReqSmartGuideAssign struct {
	OutTradeNo string `json:"out_trade_no"`
	SubMchid   string `json:"sub_mchid"`
}

type RespSmartGuideQuery struct {
	TotalCount int `json:"total_count"`
	Limit      int `json:"limit"`
	Offset     int `json:"offset"`
	Data       []struct {
		GuideId string `json:"guide_id"`
		StoreId int    `json:"store_id"`
		Name    string `json:"name"`
		Mobile  string `json:"mobile"`
		Userid  string `json:"userid"`
		WorkId  string `json:"work_id"`
	} `json:"data"`
}

type ReqSmartGuideUpdate struct {
	SubMchId    string `json:"sub_mchid"`
	Name        string `json:"name"`
	Mobile      string `json:"mobile"`
	QrCode      string `json:"qr_code"`
	Avatar      string `json:"avatar"`
	GroupQrcode string `json:"group_qrcode"`
}
