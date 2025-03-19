package custom

type RespComplaintsList struct {
	Data []struct {
		ComplaintID        string `json:"complaint_id"`
		ComplaintTime      string `json:"complaint_time"`
		ComplaintDetail    string `json:"complaint_detail"`
		ComplaintState     string `json:"complaint_state"`
		PayerPhone         string `json:"payer_phone"`
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
