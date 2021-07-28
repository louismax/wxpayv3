package custom

type SettlementAccount struct {
	AccountType   string `json:"account_type"`
	AccountBank   string `json:"account_bank"`
	BankName      string `json:"bank_name"`
	BankBranchId  string `json:"bank_branch_id"`
	AccountNumber string `json:"account_number"`
	VerifyResult  string `json:"verify_result"`
}
