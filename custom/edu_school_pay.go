package custom

type ReqEduSchoolPayPreSign struct {
	AppId                    string `json:"appid"`
	OpenId                   string `json:"openid"`
	PlanId                   string `json:"plan_id"`
	UserId                   string `json:"user_id"`
	SchoolId                 string `json:"school_id"`
	OutContractCode          string `json:"out_contract_code"`
	ContractMode             string `json:"contract_mode"`
	DowngradeDefaultContract bool   `json:"downgrade_default_contract"`
	identity                 *struct {
		RealName string `json:"real_name"`
	} `json:"identity,omitempty"`
	TradeScene string `json:"trade_scene"`
}
