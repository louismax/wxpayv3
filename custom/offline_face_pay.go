package custom

//RespOrganizationInfo RespOrganizationInfo
type RespOrganizationInfo struct {
	OrganizationId   string `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
}

//ReqObtainAuthToken ReqObtainAuthToken
type ReqObtainAuthToken struct {
	Scene       string      `json:"scene"`
	WebInitData WebInitData `json:"web_init_data"`
}

//WebInitData WebInitData
type WebInitData struct {
	OutUserId      string `json:"out_user_id"`     //商户侧用户ID
	OrganizationId string `json:"organization_id"` //机构ID
}

//RespObtainAuthToken RespObtainAuthToken
type RespObtainAuthToken struct {
	Token string `json:"token"`
}

//ReqPayCredential ReqPayCredential
type ReqPayCredential struct {
	PayCredential   string          `json:"pay_credential"`    //支付凭证
	MerchantInfo    MerchantInfo    `json:"merchant_info"`     //商户信息
	TradeAmountInfo TradeAmountInfo `json:"trade_amount_info"` //金额信息
	SceneInfo       SceneInfo       `json:"scene_info"`        //支付场景信息
	DeviceInfo      DeviceInfo      `json:"device_info"`       //设备信息
	GoodsTag        string          `json:"goods_tag"`         //优惠标记
	Description     string          `json:"description"`       //商品信息
	Attach          string          `json:"attach"`            //商户附加信息
	OutTradeNo      string          `json:"out_trade_no"`      //商户订单号
	BusinessInfo    BusinessInfo    `json:"business_info"`     //业务信息
}

//MerchantInfo MerchantInfo
type MerchantInfo struct {
	Mchid    string `json:"mchid"`     //商户号
	SubMchid string `json:"sub_mchid"` //子商户号
	Appid    string `json:"appid"`     //商户公众号
	SubAppid string `json:"sub_appid"` //子商户公众号
}

//TradeAmountInfo TradeAmountInfo
type TradeAmountInfo struct {
	Amount   int64  `json:"amount"`   //总金额
	Currency string `json:"currency"` //货币类型
}

//SceneInfo SceneInfo
type SceneInfo struct {
	DeviceIp string `json:"device_ip"` //设备IP
}

// DeviceInfo DeviceInfo
type DeviceInfo struct {
	Mac string `json:"mac"` //设备mc地址
}

// BusinessInfo BusinessInfo
type BusinessInfo struct {
	BusinessProductId int `json:"business_product_id"` //平台产品ID
	BusinessSceneId   int `json:"business_scene_id"`   //平台场景ID
}

//RespPayCredential RespPayCredential
type RespPayCredential struct {
	MerchantInfo    MerchantInfo    `json:"merchant_info"`     //商户信息
	PayerInfo       PayerInfo       `json:"payer_info"`        //支付用户信息
	TradeAmountInfo TradeAmountInfo `json:"trade_amount_info"` //金额信息
	PromotionList   []PromotionList `json:"promotion_list"`    //优惠信息
	SceneInfo       SceneInfo       `json:"scene_info"`        //支付场景信息
	CorePaymentInfo CorePaymentInfo `json:"core_payment_info"` //支付信息
	TradeType       string          `json:"trade_type"`        //交易类型
	TradeState      string          `json:"trade_state"`       //交易状态
	TradeStateDesc  string          `json:"trade_state_desc"`  //交易状态描述
	Body            string          `json:"body"`              //商品信息
	Attach          string          `json:"attach"`            //商户附加信息
	PaymentTime     string          `json:"payment_time"`      //支付成功时间
	TransactionId   string          `json:"transaction_id"`    //微信订单号
	OutTradeNo      string          `json:"out_trade_no"`      //商户订单号
}

//PayerInfo PayerInfo
type PayerInfo struct {
	Openid    string `json:"openid"`     //公众号下的openid
	SubOpenid string `json:"sub_openid"` //子公众号下的openid
}

//PromotionList PromotionList
type PromotionList struct {
	PromotionId        string          `json:"promotion_id"`        //优惠ID
	Name               string          `json:"name"`                //优惠名称
	AmountInfo         TradeAmountInfo `json:"amount_info"`         //优惠金额
	WxPayContribute    int64           `json:"wxpay_contribute"`    //微信出资金额
	MerchantContribute int64           `json:"merchant_contribute"` //商家出资金额
	OtherContribute    int64           `json:"other_contribute"`    //其他出资金额
}

//CorePaymentInfo CorePaymentInfo
type CorePaymentInfo struct {
	BankType        string `json:"bank_type"` //付款银行
	CorePaymentInfo struct {
		Amount   int64  `json:"amount"`   //支付金额
		Currency string `json:"currency"` //支付币种
	} `json:"core_payment_info"`
}

//RespQueryFaceUserInfo RespQueryFaceUserInfo
type RespQueryFaceUserInfo struct {
	UserId         string      `json:"user_id"`         //微信侧刷脸用户唯一ID
	OutUserId      string      `json:"out_user_id"`     //商户刷脸用户ID
	OrganizationId string      `json:"organization_id"` //机构ID
	UserName       string      `json:"user_name"`       //姓名
	UserType       string      `json:"user_type"`       //用户类型 学生：STUDENT教职工：STAFF
	StudentInfo    StudentInfo `json:"student_info"`    //学生信息
	StaffInfo      StaffInfo   `json:"staff_info"`      //教职工信息
	Status         string      `json:"status"`          //用户状态 NOMAL：正常状态 DISABLED：禁用状态，此时支付被限制
	ContractState  string      `json:"contract_state"`  //签约状态
	FaceImageOk    bool        `json:"face_image_ok"`   //人脸图片上传状态
	ContractId     string      `json:"contract_id"`     //签约ID
}

//StudentInfo StudentInfo
type StudentInfo struct {
	ClassName string `json:"class_name"`
}

//StaffInfo StaffInfo
type StaffInfo struct {
	Occupation string `json:"occupation"`
}

//ReqUpdateUserInfo ReqUpdateUserInfo
type ReqUpdateUserInfo struct {
	OrganizationId string               `json:"organization_id"`
	OutUserId      string               `json:"out_user_id"`
	RequestData    ReqUpdateRequestData `json:"request_data"`
}

//ReqUpdateRequestData ReqUpdateRequestData
type ReqUpdateRequestData struct {
	UserName    string       `json:"user_name"`
	UserType    string       `json:"user_type"`
	StudentInfo *StudentInfo `json:"student_info,omitempty"`
	StaffInfo   *StaffInfo   `json:"staff_info,omitempty"`
	Status      string       `json:"status"`
	Phone       string       `json:"phone"`
}

//ReqPresignToken ReqPresignToken
type ReqPresignToken struct {
	BusinessName  string         `json:"business_name"`             //业务类型
	FacePayUser   FacePayUser    `json:"facepay_user"`              //刷脸用户信息
	LimitBankCard *LimitBankCard `json:"limit_bank_card,omitempty"` //签约银行卡信息
	ContractMode  string         `json:"contract_mode,omitempty"`   //签约模式 LIMIT_BANK_CARD：指定卡签约；PRIORITY_BANK_CARD：优先卡签约；LIMIT_NONE：任意卡签约
}

//FacePayUser FacePayUser
type FacePayUser struct {
	OutUserId          string          `json:"out_user_id"`
	IdentificationName string          `json:"identification_name,omitempty"`
	OrganizationId     string          `json:"organization_id"`
	Identification     *Identification `json:"identification,omitempty"`
	Phone              string          `json:"phone,omitempty"`
}

//Identification Identification
type Identification struct {
	IdentificationType   string `json:"identification_type"`
	IdentificationNumber string `json:"identification_number"`
}

//LimitBankCard LimitBankCard
type LimitBankCard struct {
	BankCardNumber     string          `json:"bank_card_number"`
	IdentificationName string          `json:"identification_name"`
	Identification     *Identification `json:"identification,omitempty"`
	ValidThru          string          `json:"valid_thru,omitempty"`
	BankType           string          `json:"bank_type,omitempty"`
	Phone              string          `json:"phone,omitempty"`
}

//RespPresignToken RespPresignToken
type RespPresignToken struct {
	PresignToken string `json:"presign_token"`
}

//ReqOfflinefaceTransactions ReqOfflinefaceTransactions
type ReqOfflinefaceTransactions struct {
	AuthCode    string           `json:"auth_code"`             //支付凭证
	SpAppid     string           `json:"sp_appid"`              //服务商appid
	SubAppid    string           `json:"sub_appid,omitempty"`   //子商户appid
	SpMchid     string           `json:"sp_mchid"`              //商户号
	SubMchid    string           `json:"sub_mchid"`             //子商户号
	Amount      OTReqAmount      `json:"amount"`                //金额信息
	SceneInfo   OTReqSceneInfo   `json:"scene_info"`            //支付场景信息
	GoodsTag    string           `json:"goods_tag,omitempty"`   //优惠标记
	Description string           `json:"description"`           //商品信息
	Attach      string           `json:"attach"`                //商户附加信息
	SettleInfo  *OTReqSettleInfo `json:"settle_info,omitempty"` //结算信息
	OutTradeNo  string           `json:"out_trade_no"`          // 商户单号
	Business    OTReqBusiness    `json:"business"`              //业务信息
}

//OTReqAmount OTReqAmount
type OTReqAmount struct {
	Total    int64  `json:"total"`    //总金额
	Currency string `json:"currency"` //货币类型
}

//OTReqSceneInfo OTReqSceneInfo
type OTReqSceneInfo struct {
	DeviceIp string `json:"device_ip"` //设备IP
}

//OTReqSettleInfo OTReqSettleInfo
type OTReqSettleInfo struct {
	ProfitSharing bool `json:"profit_sharing,omitempty"` //是否支持分账
}

//OTReqBusiness OTReqBusiness
type OTReqBusiness struct {
	BusinessProductId int `json:"business_product_id"` //平台产品ID
	BusinessSceneId   int `json:"business_scene_id"`   //平台场景ID
}

//RespOfflinefaceTransactions RespOfflinefaceTransactions
type RespOfflinefaceTransactions struct {
	SpAppid                string          `json:"sp_appid"`
	SubAppid               string          `json:"sub_appid"`
	SpMchid                string          `json:"sp_mchid"`  //商户号
	SubMchid               string          `json:"sub_mchid"` //子商户号
	Payer                  OTRespPayer     `json:"payer"`
	Amount                 OTRespAmount    `json:"amount"`
	PromotionDetail        []PromotionList `json:"promotion_detail"`         //优惠信息
	SceneInfo              OTRespSceneInfo `json:"scene_info"`               //支付场景信息
	BankType               string          `json:"bank_type"`                //付款银行
	TradeType              string          `json:"trade_type"`               //交易类型
	TradeState             string          `json:"trade_state"`              //交易状态
	TradeStateDescription  string          `json:"trade_state_description"`  //交易描述
	ErrorType              string          `json:"error_type"`               //trade_state为PAYERROR时存在，“NOT_ENOUGH”和“NOTENOUGH”表示用户余额不足
	DebtState              string          `json:"debt_state"`               //欠款状态
	Description            string          `json:"description"`              //商品信息
	Attach                 string          `json:"attach"`                   //商户附加信息
	SuccessTime            string          `json:"success_time"`             //支付成功时间
	TransactionId          string          `json:"transaction_id"`           //微信订单号
	RepaymentTransactionId string          `json:"repayment_transaction_id"` //还款微信单号
	OutTradeNo             string          `json:"out_trade_no"`             //商户单号
}

//OTRespPayer OTRespPayer
type OTRespPayer struct {
	SPOpenid  string `json:"sp_openid"`  //公众号下的openid
	SubOpenid string `json:"sub_openid"` //子公众号下的openid
}

//OTRespAmount OTRespAmount
type OTRespAmount struct {
	Total       int64  `json:"total"`        //订单金额
	PayTotal    int64  `json:"pay_total"`    //用户支付金额
	Currency    string `json:"currency"`     //货币类型
	PayCurrency string `json:"pay_currency"` //用户支付货币类型
}

//OTRespSceneInfo OTRespSceneInfo
type OTRespSceneInfo struct {
	DeviceIp string `json:"device_ip"` //设备IP
}

//RespContractQuery RespContractQuery
type RespContractQuery struct {
	ContractId             string `json:"contract_id"`              //签约ID
	Mchid                  string `json:"mchid"`                    //商户号
	OrganizationId         string `json:"organization_id"`          //机构ID
	UserId                 string `json:"user_id"`                  //用户ID
	OpenId                 string `json:"openid"`                   //签约用户openid
	ContractState          string `json:"contract_state"`           //签约状态
	ContractSignedTime     string `json:"contract_signed_time"`     //签约时间
	ContractTerminatedTime string `json:"contract_terminated_time"` //解约时间
	ContractMode           string `json:"contract_mode"`            //签约模式 LIMIT_BANK_CARD：指定卡签约；PRIORITY_BANK_CARD：优先卡签约；LIMIT_NONE：任意卡签约
	ContractBankCardFrom   string `json:"contract_bank_card_from"`  //签约卡来源 MERCHANT_LIMITED_BANK_CARD：商户指定的签约卡；USER_SELECT_FREE：用户选择的签约卡
}

//FaceMessageCiphertext 签约解约报文数据
type FaceMessageCiphertext struct {
	ID           string `json:"id"`            //通知唯一ID
	CreateTime   string `json:"create_time"`   //通知创建的时间
	EventType    string `json:"event_type"`    //通知的类型
	ResourceType string `json:"resource_type"` //通知的资源数据类型
	Resource     struct {
		Algorithm      string `json:"algorithm"`       //加密算法类型
		Ciphertext     string `json:"ciphertext"`      //数据密文
		OriginalType   string `json:"original_type"`   //原始回调类型
		AssociatedData string `json:"associated_data"` //附加数据
		Nonce          string `json:"nonce"`           //随机串
	} `json:"resource"` //通知资源数据
	Summary string `json:"summary"` //回调摘要
}

// FaceMessagePlaintext 签约解约解密明文
type FaceMessagePlaintext struct {
	UserId           string `json:"user_id"`            //微信刷脸用户唯一标识
	OutUserId        string `json:"out_user_id"`        //商户刷脸用户唯一标识
	OrganizationId   string `json:"organization_id"`    //机构编号
	MchId            string `json:"mch_id"`             //微信支付分配的商户号
	NotifyCreateTime string `json:"notify_create_time"` //通知创建时间
	Appid            string `json:"appid"`              //微信APPID
	Openid           string `json:"openid"`             //微信openid
}

//RespQueryRepurchaseUsersList RespQueryRepurchaseUsersList
type RespQueryRepurchaseUsersList struct {
	FaceCollections []FaceCollections `json:"face_collections"`
}

// FaceCollections FaceCollections
type FaceCollections struct {
	CollectionId            string `json:"collection_id"`
	UserId                  string `json:"user_id"`
	OrganizationId          string `json:"organization_id"`
	CollectionState         string `json:"collection_state"`
	RegisterPhotoUploadTime string `json:"register_photo_upload_time"`
	ConfirmTime             string `json:"confirm_time"`
}

// ReqGetAuthInfo ReqGetAuthInfo
type ReqGetAuthInfo struct {
	SpAppid        string `json:"sp_appid"`
	SubAppid       string `json:"sub_appid"`
	SubMchid       string `json:"sub_mchid"`
	DeviceId       string `json:"device_id"`
	RawData        string `json:"raw_data"`
	OrganizationId string `json:"organization_id"`
}

// RespGetAuthInfo RespGetAuthInfo
type RespGetAuthInfo struct {
	AuthInfo string `json:"authinfo"`
}

// ReqGetRepaymentUrl ReqGetRepaymentUrl
type ReqGetRepaymentUrl struct {
	OutUserId      string `json:"out_user_id"`
	OrganizationId string `json:"organization_id"`
}

// RespGetRepaymentUrl RespGetRepaymentUrl
type RespGetRepaymentUrl struct {
	RepaymentUrl string `json:"repayment_url"`
	ExpireAt     string `json:"expire_at"`
}
