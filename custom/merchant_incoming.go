package custom

//SettlementAccount SettlementAccount
type SettlementAccount struct {
	AccountType   string `json:"account_type"`
	AccountBank   string `json:"account_bank"`
	BankName      string `json:"bank_name"`
	BankBranchId  string `json:"bank_branch_id"`
	AccountNumber string `json:"account_number"`
	VerifyResult  string `json:"verify_result"`
}

//RespGetStatusRepairOrder RespGetStatusRepairOrder
type RespGetStatusRepairOrder struct {
	BusinessCode      string        `json:"business_code"`
	ApplymentId       int64         `json:"applyment_id"`
	SubMchid          string        `json:"sub_mchid"`
	SignUrl           string        `json:"sign_url"`
	ApplymentState    string        `json:"applyment_state"`
	ApplymentStateMsg string        `json:"applyment_state_msg"`
	AuditDetail       []AuditDetail `json:"audit_detail"`
}

//AuditDetail AuditDetail
type AuditDetail struct {
	Field        string `json:"field"`
	FieldName    string `json:"field_name"`
	RejectReason string `json:"reject_reason"`
}

type ReqIncomingSubmitApplication struct {
	BusinessCode    string            `json:"business_code"`
	ContactInfo     ContactInfo       `json:"contact_info"`
	SubjectInfo     SubjectInfo       `json:"subject_info"`
	BusinessInfo    ApplyBusinessInfo `json:"business_info"`
	SettlementInfo  SettlementInfo    `json:"settlement_info"`
	BankAccountInfo BankAccountInfo   `json:"bank_account_info"`
	AdditionInfo    AdditionInfo      `json:"addition_info"`
}

// s1 ------------------------------------------------------

type ContactInfo struct {
	ContactName     string `json:"contact_name"`
	ContactIDNumber string `json:"contact_id_number,omitempty"`
	Openid          string `json:"openid,omitempty"`
	MobilePhone     string `json:"mobile_phone"`
	ContactEmail    string `json:"contact_email"`
}

type SubjectInfo struct {
	SubjectType           string               `json:"subject_type"`
	MicroBizInfo          *MicroBizInfo        `json:"micro_biz_info,omitempty"`          //小微辅助证明材料（小微专用）
	BusinessLicenseInfo   *BusinessLicenseInfo `json:"business_license_info,omitempty"`   //营业执照(特约)
	CertificateInfo       *CertificateInfo     `json:"certificate_info,omitempty"`        //登记证书 (特约)
	OrganizationInfo      *OrganizationInfo    `json:"organization_info,omitempty"`       //组织机构代码证(特约)
	CertificateLetterCopy string               `json:"certificate_letter_copy,omitempty"` //单位证明函照片(特约)
	IdentityInfo          IdentityInfo         `json:"identity_info"`                     //经营者/法人身份证件（公共）
	UboInfo               *UboInfo             `json:"ubo_info,omitempty"`                //最终受益人信息(UBO) （特约）
}

type ApplyBusinessInfo struct {
	MerchantShortname string     `json:"merchant_shortname"`
	ServicePhone      string     `json:"service_phone"`
	SalesInfo         *SalesInfo `json:"sales_info,omitempty"` //特约商户进件需要
}

type SettlementInfo struct {
	SettlementID        string    `json:"settlement_id"`
	QualificationType   string    `json:"qualification_type"`
	Qualifications      []string `json:"qualifications,omitempty"`
	ActivitiesID        string    `json:"activities_id,omitempty"`
	ActivitiesRate      string    `json:"activities_rate,omitempty"`
	ActivitiesAdditions []string `json:"activities_additions,omitempty"`
}

type BankAccountInfo struct {
	BankAccountType string `json:"bank_account_type"`
	AccountName     string `json:"account_name"`
	AccountBank     string `json:"account_bank"`
	BankAddressCode string `json:"bank_address_code"`
	BankBranchID    string `json:"bank_branch_id,omitempty"`
	BankName        string `json:"bank_name,omitempty"`
	AccountNumber   string `json:"account_number"`
}

type AdditionInfo struct {
	LegalPersonCommitment string    `json:"legal_person_commitment,omitempty"`
	LegalPersonVideo      string    `json:"legal_person_video,omitempty"`
	BusinessAdditionPics  []string `json:"business_addition_pics,omitempty"`
	BusinessAdditionMsg   string    `json:"business_addition_msg,omitempty"`
}

// e1 ------------------------------------------------------

// s2 ------------------------------------------------------

type MicroBizInfo struct {
	MicroBizType    string           `json:"micro_biz_type"`
	MicroStoreInfo  *MicroStoreInfo  `json:"micro_store_info,omitempty"`  //门店场所
	MicroMobileInfo *MicroMobileInfo `json:"micro_mobile_info,omitempty"` //流动经营/便民服务
	MicroOnlineInfo *MicroOnlineInfo `json:"micro_online_info,omitempty"` //线上商品/服务交易
}

type BusinessLicenseInfo struct {
	LicenseCopy   string `json:"license_copy"`
	LicenseNumber string `json:"license_number"`
	MerchantName  string `json:"merchant_name"`
	LegalPerson   string `json:"legal_person"`
}
type CertificateInfo struct {
	CertCopy       string `json:"cert_copy"`
	CertType       string `json:"cert_type"`
	CertNumber     string `json:"cert_number"`
	MerchantName   string `json:"merchant_name"`
	CompanyAddress string `json:"company_address"`
	LegalPerson    string `json:"legal_person"`
	PeriodBegin    string `json:"period_begin"`
	PeriodEnd      string `json:"period_end"`
}

type OrganizationInfo struct {
	OrganizationCopy string `json:"organization_copy"`
	OrganizationCode string `json:"organization_code"`
	OrgPeriodBegin   string `json:"org_period_begin"`
	OrgPeriodEnd     string `json:"org_period_end"`
}

type IdentityInfo struct {
	IDDocType  string      `json:"id_doc_type"`
	IDCardInfo *IDCardInfo `json:"id_card_info,omitempty"` //身份证信息
	IDDocInfo  *IDDocInfo   `json:"id_doc_info,omitempty"` //其他证件信息
	Owner      bool        `json:"owner"`
}

type UboInfo struct {
	IDType         string `json:"id_type"`
	IDCardCopy     string `json:"id_card_copy"`
	IDCardNational string `json:"id_card_national"`
	IDDocCopy      string `json:"id_doc_copy"`
	Name           string `json:"name"`
	IDNumber       string `json:"id_number"`
	IDPeriodBegin  string `json:"id_period_begin"`
	IDPeriodEnd    string `json:"id_period_end"`
}

type SalesInfo struct {
	SalesScenesType []string         `json:"sales_scenes_type"`
	BizStoreInfo    *BizStoreInfo    `json:"biz_store_info,omitempty"`
	MpInfo          *MpInfo          `json:"mp_info,omitempty"`
	MiniProgramInfo *MiniProgramInfo `json:"mini_program_info,omitempty"`
	AppInfo         *AppInfo         `json:"app_info,omitempty"`
	WebInfo         *WebInfo         `json:"web_info,omitempty"`
	WeworkInfo      *WeworkInfo      `json:"wework_info,omitempty"`
}

// e2 ------------------------------------------------------

// s3 ------------------------------------------------------

type MicroStoreInfo struct {
	MicroName        string `json:"micro_name"`
	MicroAddressCode string `json:"micro_address_code"`
	MicroAddress     string `json:"micro_address"`
	StoreEntrancePic string `json:"store_entrance_pic"`
	MicroIndoorCopy  string `json:"micro_indoor_copy"`
	StoreLongitude   string `json:"store_longitude,omitempty"`
	StoreLatitude    string `json:"store_latitude,omitempty"`
}

type MicroMobileInfo struct {
	MicroMobileName    string   `json:"micro_mobile_name"`
	MicroMobileCity    string   `json:"micro_mobile_city"`
	MicroMobileAddress string   `json:"micro_mobile_address"`
	MicroMobilePics    []string `json:"micro_mobile_pics"`
}
type MicroOnlineInfo struct {
	MicroOnlineStore string `json:"micro_online_store"`
	MicroEcName      string `json:"micro_ec_name"`
	MicroQrcode      string `json:"micro_qrcode,omitempty"`
	MicroLink        string `json:"micro_link,omitempty"`
}

type IDCardInfo struct {
	IDCardCopy      string `json:"id_card_copy"`
	IDCardNational  string `json:"id_card_national"`
	IDCardName      string `json:"id_card_name"`
	IDCardNumber    string `json:"id_card_number"`
	CardPeriodBegin string `json:"card_period_begin"`
	CardPeriodEnd   string `json:"card_period_end"`
}

type IDDocInfo struct {
	IDDocCopy      string `json:"id_doc_copy"`
	IDDocName      string `json:"id_doc_name"`
	IDDocNumber    string `json:"id_doc_number"`
	DocPeriodBegin string `json:"doc_period_begin"`
	DocPeriodEnd   string `json:"doc_period_end"`
}

type BizStoreInfo struct {
	BizStoreName     string   `json:"biz_store_name"`
	BizAddressCode   string   `json:"biz_address_code"`
	BizStoreAddress  string   `json:"biz_store_address"`
	StoreEntrancePic []string `json:"store_entrance_pic"`
	IndoorPic        []string `json:"indoor_pic"`
	BizSubAppid      string   `json:"biz_sub_appid,omitempty"`
}
type MpInfo struct {
	MpAppid    string    `json:"mp_appid,omitempty"`
	MpSubAppid string    `json:"mp_sub_appid,omitempty"`
	MpPics     []string `json:"mp_pics,omitempty"`
}
type MiniProgramInfo struct {
	MiniProgramAppid    string    `json:"mini_program_appid,omitempty"`
	MiniProgramSubAppid string    `json:"mini_program_sub_appid,omitempty"`
	MiniProgramPics     []string `json:"mini_program_pics,omitempty"`
}
type AppInfo struct {
	AppAppid    string   `json:"app_appid,omitempty"`
	AppSubAppid string   `json:"app_sub_appid,omitempty"`
	AppPics     []string `json:"app_pics,omitempty"`
}
type WebInfo struct {
	Domain           string `json:"domain"`
	WebAuthorisation string `json:"web_authorisation,omitempty"`
	WebAppid         string `json:"web_appid,omitempty"`
}
type WeworkInfo struct {
	//CorpID     string   `json:"corp_id,omitempty"`
	SubCorpID  string   `json:"sub_corp_id,omitempty"`
	WeworkPics []string `json:"wework_pics,omitempty"`
}

// e3 ------------------------------------------------------

type RespIncomingSubmitApplication struct {
	ApplymentID int64 `json:"applyment_id"`
}

type ReqModifySettlement struct {
	AccountType     string `json:"account_type"`
	AccountBank     string `json:"account_bank"`
	BankAddressCode string `json:"bank_address_code"`
	BankName        string `json:"bank_name"`
	BankBranchID    string `json:"bank_branch_id"`
	AccountNumber   string `json:"account_number"`
}
