package custom

// ReqComplaintNotifications 创建、更新投诉通知回调地址的请求。
type ReqComplaintNotifications struct {
	URL string `json:"url"` //通知地址，必填，最长255字符，仅支持HTTPS
}

// RespComplaintNotifications 创建、查询、更新投诉通知回调地址的结果。
type RespComplaintNotifications struct {
	Mchid string `json:"mchid"` //配置回调地址的商户号；服务商模式为服务商商户号
	URL   string `json:"url"`   //投诉通知回调地址
}

type RespComplaintsList struct {
	Data []struct {
		ComplaintID        string `json:"complaint_id"`
		ComplaintTime      string `json:"complaint_time"`
		ComplaintDetail    string `json:"complaint_detail"`
		ComplaintState     string `json:"complaint_state"`
		ComplaintedMchid   string `json:"complainted_mchid"`
		PayerPhone         string `json:"payer_phone"`
		PayerOpenid        string `json:"payer_openid"`
		ComplaintOrderInfo []struct {
			TransactionID string `json:"transaction_id"`
			OutTradeNo    string `json:"out_trade_no"`
			Amount        int    `json:"amount"`
		} `json:"complaint_order_info"`
		ComplaintFullRefunded bool `json:"complaint_full_refunded"`
		IncomingUserResponse  bool `json:"incoming_user_response"`
		UserComplaintTimes    int  `json:"user_complaint_times"`
		ComplaintMediaList    []struct {
			MediaType string   `json:"media_type"`
			MediaURL  []string `json:"media_url"`
		} `json:"complaint_media_list"`
		ProblemDescription string   `json:"problem_description"`
		ProblemType        string   `json:"problem_type"`
		ApplyRefundAmount  int      `json:"apply_refund_amount"`
		UserTagList        []string `json:"user_tag_list"`
		ServiceOrderInfo   []struct {
			OrderID    string `json:"order_id"`
			OutOrderNo string `json:"out_order_no"`
			State      string `json:"state"`
		} `json:"service_order_info"`
		AdditionalInfo struct {
			Type           string `json:"type"`
			SharePowerInfo struct {
				ReturnTime        string `json:"return_time"`
				ReturnAddressInfo struct {
					ReturnAddress string `json:"return_address"`
					Longitude     string `json:"longitude"`
					Latitude      string `json:"latitude"`
				} `json:"return_address_info"`
				IsReturnedToSameMachine bool `json:"is_returned_to_same_machine"`
			} `json:"share_power_info"`
		} `json:"additional_info"`
		InPlatformService    bool `json:"in_platform_service"`
		NeedImmediateService bool `json:"need_immediate_service"`
	} `json:"data"`
	Limit      int `json:"limit"`
	Offset     int `json:"offset"`
	TotalCount int `json:"total_count"`
}
