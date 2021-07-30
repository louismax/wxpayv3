package custom

type SettlementAccount struct {
	AccountType   string `json:"account_type"`
	AccountBank   string `json:"account_bank"`
	BankName      string `json:"bank_name"`
	BankBranchId  string `json:"bank_branch_id"`
	AccountNumber string `json:"account_number"`
	VerifyResult  string `json:"verify_result"`
}

type RespGetStatusRepairOrder struct {
	BusinessCode      string        `json:"business_code"`
	ApplymentId       int64         `json:"applyment_id"`
	SubMchid          string        `json:"sub_mchid"`
	SignUrl           string        `json:"sign_url"`
	ApplymentState    string        `json:"applyment_state"`
	ApplymentStateMsg string        `json:"applyment_state_msg"`
	AuditDetail       []AuditDetail `json:"audit_detail"`
}
type AuditDetail struct {
	Field        string `json:"field"`
	FieldName    string `json:"field_name"`
	RejectReason string `json:"reject_reason"`
}
