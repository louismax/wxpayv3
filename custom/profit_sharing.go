package custom

//ReqInitiateProfitSharing 发起分账请求参数
type ReqInitiateProfitSharing struct {
	SubMchid        string                      `json:"sub_mchid"`
	Appid           string                      `json:"appid"`
	SubAppid        string                      `json:"sub_appid"`
	TransactionID   string                      `json:"transaction_id"`
	OutOrderNo      string                      `json:"out_order_no"`
	Receivers       []ReqProfitSharingReceivers `json:"receivers"` //分账接收方列表
	UnfreezeUnsplit bool                        `json:"unfreeze_unsplit"`
}

//ReqProfitSharingReceivers 分账接收方信息
type ReqProfitSharingReceivers struct {
	Type        string `json:"type"`
	Account     string `json:"account"`
	Name        string `json:"name"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

//RespInitiateProfitSharing 发起分账返回参数
type RespInitiateProfitSharing struct {
	SubMchid      string                       `json:"sub_mchid"`
	TransactionID string                       `json:"transaction_id"`
	OutOrderNo    string                       `json:"out_order_no"`
	OrderID       string                       `json:"order_id"`
	State         string                       `json:"state"`
	Receivers     []RespProfitSharingReceivers `json:"receivers"`
}

//RespProfitSharingReceivers is RespProfitSharingReceivers
type RespProfitSharingReceivers struct {
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Account     string `json:"account"`
	Result      string `json:"result"`
	DetailID    string `json:"detail_id"`
	FailReason  string `json:"fail_reason"`
	CreateTime  string `json:"create_time"`
	FinishTime  string `json:"finish_time"`
}

//RespQueryProfitSharingResult is RespQueryProfitSharingResult
type RespQueryProfitSharingResult struct {
	SubMchid      string                       `json:"sub_mchid"`
	TransactionID string                       `json:"transaction_id"`
	OutOrderNo    string                       `json:"out_order_no"`
	OrderID       string                       `json:"order_id"`
	State         string                       `json:"state"`
	Receivers     []RespProfitSharingReceivers `json:"receivers"`
}

//ReqInitiateProfitSharingReturnOrders 请求分账回退请求参数
type ReqInitiateProfitSharingReturnOrders struct {
	SubMchid    string `json:"sub_mchid"`
	OrderID     string `json:"order_id"`
	OutReturnNo string `json:"out_return_no"`
	ReturnMchid string `json:"return_mchid"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

//RespInitiateProfitSharingReturnOrders 请求分账回退返回结果
type RespInitiateProfitSharingReturnOrders struct {
	SubMchid    string `json:"sub_mchid"`
	OrderID     string `json:"order_id"`
	OutOrderNo  string `json:"out_order_no"`
	OutReturnNo string `json:"out_return_no"`
	ReturnID    string `json:"return_id"`
	ReturnMchid string `json:"return_mchid"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Result      string `json:"result"`
	FailReason  string `json:"fail_reason"`
	CreateTime  string `json:"create_time"`
	FinishTime  string `json:"finish_time"`
}

//RespQueryProfitSharingReturnOrders 查询分账回退结果返回参数
type RespQueryProfitSharingReturnOrders struct {
	SubMchid    string `json:"sub_mchid"`
	OrderID     string `json:"order_id"`
	OutOrderNo  string `json:"out_order_no"`
	OutReturnNo string `json:"out_return_no"`
	ReturnID    string `json:"return_id"`
	ReturnMchid string `json:"return_mchid"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Result      string `json:"result"`
	FailReason  string `json:"fail_reason"`
	CreateTime  string `json:"create_time"`
	FinishTime  string `json:"finish_time"`
}

//ReqUnfreezeRemainingFunds 解冻剩余资金请求参数
type ReqUnfreezeRemainingFunds struct {
	SubMchid      string `json:"sub_mchid"`
	TransactionID string `json:"transaction_id"`
	OutOrderNo    string `json:"out_order_no"`
	Description   string `json:"description"`
}

//RespUnfreezeRemainingFunds 解冻剩余资金返回参数
type RespUnfreezeRemainingFunds struct {
	SubMchid      string                       `json:"sub_mchid"`
	TransactionID string                       `json:"transaction_id"`
	OutOrderNo    string                       `json:"out_order_no"`
	OrderID       string                       `json:"order_id"`
	State         string                       `json:"state"`
	Receivers     []RespProfitSharingReceivers `json:"receivers"`
}

//RespQueryRemainingFrozenAmount 查询剩余待分金额返回参数
type RespQueryRemainingFrozenAmount struct {
	TransactionID string `json:"transaction_id"`
	UnsplitAmount int    `json:"unsplit_amount"`
}

//RespQueryMaximumSplitRatio 查询最大分账比例API
type RespQueryMaximumSplitRatio struct {
	SubMchid string `json:"sub_mchid"`
	MaxRatio int    `json:"max_ratio"`
}

//ReqAddProfitSharingReceiver 添加分账接收方请求参数
type ReqAddProfitSharingReceiver struct {
	SubMchid       string `json:"sub_mchid"`
	Appid          string `json:"appid"`
	SubAppid       string `json:"sub_appid"`
	Type           string `json:"type"`
	Account        string `json:"account"`
	Name           string `json:"name"`
	RelationType   string `json:"relation_type"`
	CustomRelation string `json:"custom_relation"`
}

//RespAddProfitSharingReceiver 添加分账接收方返回参数
type RespAddProfitSharingReceiver struct {
	SubMchid       string `json:"sub_mchid"`
	Type           string `json:"type"`
	Account        string `json:"account"`
	Name           string `json:"name"`
	RelationType   string `json:"relation_type"`
	CustomRelation string `json:"custom_relation"`
}

//ReqDeleteProfitSharingReceiver 删除分账接收方请求参数
type ReqDeleteProfitSharingReceiver struct {
	SubMchid string `json:"sub_mchid"`
	Appid    string `json:"appid"`
	SubAppid string `json:"sub_appid"`
	Type     string `json:"type"`
	Account  string `json:"account"`
}

//RespDeleteProfitSharingReceiver 删除分账接收方返回参数
type RespDeleteProfitSharingReceiver struct {
	SubMchid string `json:"sub_mchid"`
	Type     string `json:"type"`
	Account  string `json:"account"`
}
