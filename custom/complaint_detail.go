package custom

import "encoding/json"

// RespComplaintDetail 投诉详情，时间字段保留渠道原始RFC3339格式。
type RespComplaintDetail struct {
	ComplaintID           string                   `json:"complaint_id"`
	ComplaintTime         string                   `json:"complaint_time"`
	ComplaintDetail       string                   `json:"complaint_detail"`
	ComplaintState        string                   `json:"complaint_state"`
	ComplaintedMchid      string                   `json:"complainted_mchid"`
	PayerPhone            string                   `json:"payer_phone"`
	PayerOpenid           string                   `json:"payer_openid"`
	ComplaintOrderInfo    []ComplaintOrderInfo     `json:"complaint_order_info"`
	ComplaintFullRefunded bool                     `json:"complaint_full_refunded"`
	IncomingUserResponse  bool                     `json:"incoming_user_response"`
	UserComplaintTimes    int                      `json:"user_complaint_times"`
	ComplaintMediaList    []ComplaintMedia         `json:"complaint_media_list"`
	ProblemDescription    string                   `json:"problem_description"`
	ProblemType           string                   `json:"problem_type"`
	ApplyRefundAmount     int64                    `json:"apply_refund_amount"`
	UserTagList           []string                 `json:"user_tag_list"`
	ServiceOrderInfo      []ComplaintServiceOrder  `json:"service_order_info"`
	AdditionalInfo        *ComplaintAdditionalInfo `json:"additional_info,omitempty"`
	InPlatformService     bool                     `json:"in_platform_service"`
	NeedImmediateService  bool                     `json:"need_immediate_service"`
}

type ComplaintOrderInfo struct {
	TransactionID string `json:"transaction_id"`
	OutTradeNo    string `json:"out_trade_no"`
	Amount        int64  `json:"amount"` // 单位分
}

type ComplaintMedia struct {
	MediaType string   `json:"media_type"`
	MediaURL  []string `json:"media_url"`
}

type ComplaintServiceOrder struct {
	OrderID    string `json:"order_id"`
	OutOrderNo string `json:"out_order_no"`
	State      string `json:"state"`
}

type ComplaintAdditionalInfo struct {
	Type           string                   `json:"type"`
	SharePowerInfo *ComplaintSharePowerInfo `json:"share_power_info,omitempty"`
}

type ComplaintSharePowerInfo struct {
	ReturnTime              string                      `json:"return_time"`
	ReturnAddressInfo       *ComplaintReturnAddressInfo `json:"return_address_info,omitempty"`
	IsReturnedToSameMachine bool                        `json:"is_returned_to_same_machine"`
}

type ComplaintReturnAddressInfo struct {
	ReturnAddress string `json:"return_address"`
	Longitude     string `json:"longitude"`
	Latitude      string `json:"latitude"`
}

type RespComplaintNegotiationHistory struct {
	Data       []ComplaintNegotiationRecord `json:"data"`
	Limit      int                          `json:"limit"`
	Offset     int                          `json:"offset"`
	TotalCount int                          `json:"total_count"`
}

type ComplaintNegotiationRecord struct {
	LogID                                    string            `json:"log_id"`
	Operator                                 string            `json:"operator"`
	OperateTime                              string            `json:"operate_time"`
	OperateType                              string            `json:"operate_type"`
	OperateDetails                           string            `json:"operate_details"`
	ImageList                                []string          `json:"image_list"`
	ComplaintMediaList                       *ComplaintMedia   `json:"complaint_media_list,omitempty"`
	UserAppyPlatformServiceReason            string            `json:"user_appy_platform_service_reason"`
	UserAppyPlatformServiceReasonDescription string            `json:"user_appy_platform_service_reason_description"`
	NormalMessage                            *ComplaintMessage `json:"normal_message,omitempty"`
	ClickMessage                             json.RawMessage   `json:"click_message,omitempty"`
}

// ReqResponseComplaint 普通回复，图片须先经UploadComplaintImage上传。
type ReqResponseComplaint struct {
	ComplaintedMchid    string                    `json:"complainted_mchid,omitempty"`
	ResponseContent     string                    `json:"response_content"`
	ResponseImages      []string                  `json:"response_images,omitempty"`
	JumpURL             string                    `json:"jump_url,omitempty"`
	JumpURLText         string                    `json:"jump_url_text,omitempty"`
	MiniProgramJumpInfo *ComplaintMiniProgramJump `json:"mini_program_jump_info,omitempty"`
}

type ReqCompleteComplaint struct {
	ComplaintedMchid string `json:"complainted_mchid,omitempty"`
}

type ReqResponseComplaintImmediateService struct {
	ComplaintedMchid string           `json:"complainted_mchid,omitempty"`
	Message          ComplaintMessage `json:"message"`
}

type RespResponseComplaintImmediateService struct {
	LogID string `json:"log_id"`
}

type ComplaintMessage struct {
	Blocks         []ComplaintMessageBlock `json:"blocks,omitempty"`
	SenderIdentity string                  `json:"sender_identity,omitempty"`
	CustomData     string                  `json:"custom_data,omitempty"`
}

type ComplaintMessageBlock struct {
	Type        string                  `json:"type,omitempty"`
	Text        *ComplaintMessageText   `json:"text,omitempty"`
	Image       *ComplaintMessageImage  `json:"image,omitempty"`
	Link        *ComplaintMessageLink   `json:"link,omitempty"`
	FAQList     *ComplaintFAQList       `json:"faq_list,omitempty"`
	Button      *ComplaintMessageButton `json:"button,omitempty"`
	ButtonGroup *ComplaintButtonGroup   `json:"button_group,omitempty"`
}

type ComplaintMessageText struct {
	Text   string `json:"text,omitempty"`
	Color  string `json:"color,omitempty"`
	IsBold bool   `json:"is_bold,omitempty"`
}

type ComplaintMessageImage struct {
	MediaID        string `json:"media_id,omitempty"`
	ImageStyleType string `json:"image_style_type,omitempty"`
}

type ComplaintMessageLink struct {
	Text        string                    `json:"text,omitempty"`
	Action      *ComplaintClickAction     `json:"action,omitempty"`
	InvalidInfo *ComplaintMessageValidity `json:"invalid_info,omitempty"`
}

type ComplaintFAQList struct {
	FAQs []ComplaintFAQ `json:"faqs,omitempty"`
}

type ComplaintFAQ struct {
	FAQID    string                `json:"faq_id,omitempty"`
	FAQTitle string                `json:"faq_title,omitempty"`
	Action   *ComplaintClickAction `json:"action,omitempty"`
}

type ComplaintMessageButton struct {
	Text        string                    `json:"text,omitempty"`
	Action      *ComplaintClickAction     `json:"action,omitempty"`
	InvalidInfo *ComplaintMessageValidity `json:"invalid_info,omitempty"`
}

type ComplaintButtonGroup struct {
	Buttons      []ComplaintInnerButton    `json:"buttons,omitempty"`
	ButtonLayout string                    `json:"button_layout,omitempty"`
	InvalidInfo  *ComplaintMessageValidity `json:"invalid_info,omitempty"`
}

type ComplaintInnerButton struct {
	Text   string                `json:"text,omitempty"`
	Action *ComplaintClickAction `json:"action,omitempty"`
}

type ComplaintClickAction struct {
	ActionType          string                    `json:"action_type,omitempty"`
	JumpURL             string                    `json:"jump_url,omitempty"`
	MiniProgramJumpInfo *ComplaintMiniProgramJump `json:"mini_program_jump_info,omitempty"`
	MessageInfo         *ComplaintMessageInfo     `json:"message_info,omitempty"`
	ActionID            string                    `json:"action_id,omitempty"`
}

type ComplaintMiniProgramJump struct {
	Appid string `json:"appid"`
	Path  string `json:"path"`
	Text  string `json:"text,omitempty"`
}

type ComplaintMessageInfo struct {
	Content    string `json:"content"`
	CustomData string `json:"custom_data,omitempty"`
}

type ComplaintMessageValidity struct {
	ExpiredTime    string `json:"expired_time,omitempty"`
	MultiClickable bool   `json:"multi_clickable,omitempty"`
}

// ComplaintNotification 是原始通知信封，Resource中的密文需先验签后解密。
type ComplaintNotification struct {
	ID           string                     `json:"id"`
	CreateTime   string                     `json:"create_time"`
	EventType    string                     `json:"event_type"`
	ResourceType string                     `json:"resource_type"`
	Summary      string                     `json:"summary"`
	Resource     ComplaintEncryptedResource `json:"resource"`
}

type ComplaintEncryptedResource struct {
	Algorithm      string `json:"algorithm"`
	Ciphertext     string `json:"ciphertext"`
	AssociatedData string `json:"associated_data"`
	Nonce          string `json:"nonce"`
	OriginalType   string `json:"original_type"`
}

type ComplaintNotificationResource struct {
	ComplaintID string `json:"complaint_id"`
	ActionType  string `json:"action_type"`
}
