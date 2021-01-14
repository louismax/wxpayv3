package wxpayv3

import (
	"encoding/json"
	"fmt"
)

const (
	SUBJECT_TYPE_MICRO = "SUBJECT_TYPE_MICRO"

	MICRO_TYPE_STORE  = "MICRO_TYPE_STORE"  //门店场所
	MICRO_TYPE_MOBILE = "MICRO_TYPE_MOBILE" //流动经营/便民服务
	MICRO_TYPE_ONLINE = "MICRO_TYPE_ONLINE" //线上商品/服务交易

	IDENTIFICATION_TYPE_IDCARD = "IDENTIFICATION_TYPE_IDCARD" //中国大陆居民-身份证

	BANK_ACCOUNT_TYPE_PERSONAL = "BANK_ACCOUNT_TYPE_PERSONAL" //经营者个人银行卡
)

//-----普通服务商子商户进件
// 目前仅对接小微商户
type ReqMerchantsIntoPieces struct {
	BusinessCode    string           `json:"business_code,omitempty"`     //业务申请编号
	ContactInfo     *ContactInfo     `json:"contact_info,omitempty"`      //超级管理员信息
	SubjectInfo     *SubjectInfo     `json:"subject_info,omitempty"`      //主体信息
	BusinessInfo    *BusinessInfo_MI `json:"business_info,omitempty"`     //经营资料
	SettlementInfo  *SettlementInfo  `json:"settlement_info,omitempty"`   //结算规则
	BankAccountInfo *BankAccountInfo `json:"bank_account_info,omitempty"` //收款银行卡
	AdditionInfo    *AdditionInfo    `json:"addition_info,omitempty"`     //-补充材料
}

type ContactInfo struct {
	ContactName     string `json:"contact_name,omitempty"`            //超级管理员姓名
	ContactIdNumber string `json:"contact_id_number,omitempty"`       //超级管理员身份证件号码
	Openid          string `json:"openid,omitempty,omitempty"`        //超级管理员微信openid
	MobilePhone     string `json:"mobile_phone,omitempty"`            //联系手机
	ContactEmail    string `json:"contact_email,omitempty,omitempty"` //联系邮箱
}

type SubjectInfo struct {
	SubjectType  string        `json:"subject_type,omitempty"`   //主体类型
	MicroBizInfo *MicroBizInfo `json:"micro_biz_info,omitempty"` //小微辅助证明材料
	IdentityInfo *IdentityInfo `json:"identity_info,omitempty"`  //经营者身份证件
}

type BusinessInfo_MI struct {
	MerchantShortname string `json:"merchant_shortname,omitempty"` //商户简称
	ServicePhone      string `json:"service_phone,omitempty"`      //客服电话
}

type SettlementInfo struct {
	SettlementId        string   `json:"settlement_id"`                  //入驻结算规则ID
	QualificationType   string   `json:"qualification_type"`             //所属行业
	Qualifications      []string `json:"qualifications,omitempty"`       //特殊资质图片
	ActivitiesId        string   `json:"activities_id,omitempty"`        //优惠费率活动ID
	ActivitiesRate      string   `json:"activities_rate,omitempty"`      //优惠费率活动值
	ActivitiesAdditions []string `json:"activities_additions,omitempty"` //优惠费率活动补充材料
}

type BankAccountInfo struct {
	BankAccountType string `json:"bank_account_type"`        //账户类型
	AccountName     string `json:"account_name"`             //开户名称
	AccountBank     string `json:"account_bank"`             //开户银行
	BankAddressCode string `json:"bank_address_code"`        //开户银行省市编码
	BankBranchId    string `json:"bank_branch_id,omitempty"` //开户银行联行号
	BankName        string `json:"bank_name,omitempty"`      //开户银行全称（含支行]
	AccountNumber   string `json:"account_number"`           //银行账号
}

type AdditionInfo struct {
	LegalPersonCommitment string   `json:"legal_person_commitment,omitempty"` //法人开户承诺函
	LegalPersonVideo      string   `json:"legal_person_video,omitempty"`      //法人开户意愿视频
	BusinessAdditionPics  []string `json:"business_addition_pics,omitempty"`  //补充材料
	BusinessAdditionMsg   string   `json:"business_addition_msg,omitempty"`   //补充说明
}

type MicroBizInfo struct {
	MicroBizYype    string           `json:"micro_biz_type"`              //小微经营类型
	MicroStoreInfo  *MicroStoreInfo  `json:"micro_store_info,omitempty"`  //门店场所 三选一
	MicroMobileInfo *MicroMobileInfo `json:"micro_mobile_info,omitempty"` //流动经营/便民服务 三选一
	MicroOnlineInfo *MicroOnlineInfo `json:"micro_online_info,omitempty"` //线上商品/服务交易 三选一
}

type IdentityInfo struct {
	IdDocType  string      `json:"id_doc_type"`  //证件类型
	IdCardInfo *IdCardInfo `json:"id_card_info"` //身份证信息
}

type MicroStoreInfo struct {
	MicroName        string `json:"micro_name,omitempty"`         //门店名称
	MicroAddressCode string `json:"micro_address_code,omitempty"` //门店省市编码
	MicroAddress     string `json:"micro_address,omitempty"`      //门店地址
	StoreEntrancePic string `json:"store_entrance_pic,omitempty"` //门店门头照片
	MicroIndoorCopy  string `json:"micro_indoor_copy,omitempty"`  //店内环境照片
	StoreLongitude   string `json:"store_longitude,omitempty"`    //门店经度
	StoreLatitude    string `json:"store_latitude,omitempty"`     //门店纬度
}

type MicroMobileInfo struct {
	MicroMobileName    string `json:"micro_mobile_name,omitempty"`    //经营/服务名称
	MicroMobileCity    string `json:"micro_mobile_city,omitempty"`    //经营/服务所在地省市
	MicroMobileAddress string `json:"micro_mobile_address,omitempty"` //经营/服务所在地（不含省市]
	MicroMobilePics    string `json:"micro_mobile_pics,omitempty"`    //经营/服务现场照片
}

type MicroOnlineInfo struct {
	MicroOnlineStore string `json:"micro_online_store,omitempty"` //线上店铺名称
	MicroEcName      string `json:"micro_ec_name,omitempty"`      //电商平台名称
	MicroQrcode      string `json:"micro_qrcode,omitempty"`       //店铺二维码
	MicroLink        string `json:"micro_link,omitempty"`         //店铺链接
}

type IdCardInfo struct {
	IdCardCopy      string `json:"id_card_copy"`      //身份证人像面照片
	IdCardNational  string `json:"id_card_national"`  //身份证国徽面照片
	IdCardName      string `json:"id_card_name"`      //身份证姓名
	IdCardNumber    string `json:"id_card_number"`    //身份证号码
	CardPeriodBegin string `json:"card_period_begin"` //身份证有效期开始时间
	CardPeriodEnd   string `json:"card_period_end"`   //身份证有效期结束时间
}

// APIUrl ReqMerchantsIntoPieces APIURL
func (this ReqMerchantsIntoPieces) APIUrl() string {
	return "/v3/applyment4sub/applyment/"
}

// Method ReqMerchantsIntoPieces Method
func (this ReqMerchantsIntoPieces) Method() string {
	return "POST"
}

// Params ReqMerchantsIntoPieces Params
func (this ReqMerchantsIntoPieces) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// RawJsonStr ReqMerchantsIntoPieces RawJsonStr
func (this ReqMerchantsIntoPieces) RawJsonStr() string {
	jsons, errs := json.Marshal(this)
	if errs != nil {
		fmt.Println(errs.Error())
		return "Error"
	}
	//fmt.Println(string(jsons))
	return string(jsons)
}

type RespMerchantsIntoPieces struct {
	ApplymentId int64 `json:"applyment_id"`
}

type ReqGetStatusRepairOrderForBusCode struct {
	BusinessCode string `json:"business_code"`
}

// APIUrl ReqGetStatusRepairOrderForBusCode APIURL
func (this ReqGetStatusRepairOrderForBusCode) APIUrl() string {
	return fmt.Sprintf("/v3/applyment4sub/applyment/business_code/%s", this.BusinessCode)
}

// Method ReqGetStatusRepairOrderForBusCode Method
func (this ReqGetStatusRepairOrderForBusCode) Method() string {
	return "GET"
}

// Params ReqGetStatusRepairOrderForBusCode Params
func (this ReqGetStatusRepairOrderForBusCode) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// RawJsonStr ReqGetStatusRepairOrderForBusCode RawJsonStr
func (this ReqGetStatusRepairOrderForBusCode) RawJsonStr() string {
	return ""
}

type ReqGetStatusRepairOrderForApplyCode struct {
	ApplymentId int64 `json:"applyment_id"`
}

// APIUrl ReqGetStatusRepairOrderForApplyCode APIURL
func (this ReqGetStatusRepairOrderForApplyCode) APIUrl() string {
	return fmt.Sprintf("/v3/applyment4sub/applyment/applyment_id/%d", this.ApplymentId)
}

// Method ReqGetStatusRepairOrderForApplyCode Method
func (this ReqGetStatusRepairOrderForApplyCode) Method() string {
	return "GET"
}

// Params ReqGetStatusRepairOrderForApplyCode Params
func (this ReqGetStatusRepairOrderForApplyCode) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// RawJsonStr ReqGetStatusRepairOrderForApplyCode RawJsonStr
func (this ReqGetStatusRepairOrderForApplyCode) RawJsonStr() string {
	return ""
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
