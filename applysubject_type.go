package wxpayv3

import (
	"encoding/json"
	"fmt"
)

const (
	SUBJECT_TYPE_ENTERPRISE          = "SUBJECT_TYPE_ENTERPRISE"          //企业
	SUBJECT_TYPE_INSTITUTIONS_CLONED = "SUBJECT_TYPE_INSTITUTIONS_CLONED" //事业单位
	SUBJECT_TYPE_INDIVIDUAL          = "SUBJECT_TYPE_INDIVIDUAL"          //个体工商户
	SUBJECT_TYPE_OTHERS              = "SUBJECT_TYPE_OTHERS"              //其他组织
	SUBJECT_TYPE_MICRO               = "SUBJECT_TYPE_MICRO"               //小微商户

	CERTIFICATE_TYPE_2388 = "CERTIFICATE_TYPE_2388" //事业单位法人证书
	CERTIFICATE_TYPE_2389 = "CERTIFICATE_TYPE_2389" //统一社会信用代码证书
	CERTIFICATE_TYPE_2390 = "CERTIFICATE_TYPE_2390" //有偿服务许可证（军队医院适用)
	CERTIFICATE_TYPE_2391 = "CERTIFICATE_TYPE_2391" //医疗机构执业许可证（军队医院适用)
	CERTIFICATE_TYPE_2392 = "CERTIFICATE_TYPE_2392" //企业营业执照（挂靠企业的党组织适用)
	CERTIFICATE_TYPE_2393 = "CERTIFICATE_TYPE_2393" //组织机构代码证（政府机关适用)
	CERTIFICATE_TYPE_2394 = "CERTIFICATE_TYPE_2394" //社会团体法人登记证书
	CERTIFICATE_TYPE_2395 = "CERTIFICATE_TYPE_2395" //民办非企业单位登记证书
	CERTIFICATE_TYPE_2396 = "CERTIFICATE_TYPE_2396" //基金会法人登记证书
	CERTIFICATE_TYPE_2397 = "CERTIFICATE_TYPE_2397" //慈善组织公开募捐资格证书
	CERTIFICATE_TYPE_2398 = "CERTIFICATE_TYPE_2398" //农民专业合作社法人营业执照
	CERTIFICATE_TYPE_2399 = "CERTIFICATE_TYPE_2399" //宗教活动场所登记证
	CERTIFICATE_TYPE_2400 = "CERTIFICATE_TYPE_2400" //其他证书/批文/证明

	MICRO_TYPE_STORE  = "MICRO_TYPE_STORE"  //门店场所
	MICRO_TYPE_MOBILE = "MICRO_TYPE_MOBILE" //流动经营/便民服务
	MICRO_TYPE_ONLINE = "MICRO_TYPE_ONLINE" //线上商品/服务交易

	IDENTIFICATION_TYPE_IDCARD            = "IDENTIFICATION_TYPE_IDCARD"            //身份证（限中国大陆居民)
	IDENTIFICATION_TYPE_OVERSEA_PASSPORT  = "IDENTIFICATION_TYPE_OVERSEA_PASSPORT"  //护照（限境外人士)
	IDENTIFICATION_TYPE_HONGKONG_PASSPORT = "IDENTIFICATION_TYPE_HONGKONG_PASSPORT" //中国香港居民-来往内地通行证
	IDENTIFICATION_TYPE_MACAO_PASSPORT    = "IDENTIFICATION_TYPE_MACAO_PASSPORT"    //中国澳门居民-来往内地通行证
	IDENTIFICATION_TYPE_TAIWAN_PASSPORT   = "IDENTIFICATION_TYPE_TAIWAN_PASSPORT"   //中国台湾居民-来往大陆通行证
)

type ReqSubmitApplication struct {
	ChannelId          string             `json:"channel_id,omitempty"`
	BusinessCode       string             `json:"business_code"`
	ContactInfo        ContactInfo        `json:"contact_info"`        //联系人信息
	SubjectInfo        SubjectInfo        `json:"subject_info"`        //主体信息
	IdentificationInfo IdentificationInfo `json:"identification_info"` //法人身份信息
}

type ContactInfo struct {
	Name         string `json:"name"`
	Mobile       string `json:"mobile"`
	IdCardNumber string `json:"id_card_number"`
}

type SubjectInfo struct {
	SubjectType          string                 `json:"subject_type"`
	BusinessLicenceInfo  BusinessLicenceInfo    `json:"business_licence_info,omitempty,omitempty"` // 营业执照信息
	CertificateInfo      CertificateInfoSA      `json:"certificate_info,omitempty,omitempty"`      //登记证书信息
	CompanyProveCopy     string                 `json:"company_prove_copy,omitempty,omitempty"`    //单位证明函照片
	AssistProveInfo      AssistProveInfo        `json:"assist_prove_info,omitempty,omitempty"`     //辅助证明材料信息
	SpecialOperationList []SpecialOperationList `json:"special_operation_list,omitempty,omitempty"`
}

type BusinessLicenceInfo struct {
	LicenceNumber    string `json:"licence_number,omitempty"`
	LicenceCopy      string `json:"licence_copy,omitempty"`
	MerchantName     string `json:"merchant_name,omitempty"`
	LegalPerson      string `json:"legal_person,omitempty"`
	CompanyAddress   string `json:"company_address,omitempty"`
	LicenceValidDate string `json:"licence_valid_date,omitempty"`
}

type CertificateInfoSA struct {
	CertType       string `json:"cert_type,omitempty"`
	CertNumber     string `json:"cert_number,omitempty"`
	CertCopy       string `json:"cert_copy,omitempty"`
	MerchantName   string `json:"merchant_name,omitempty"`
	LegalPerson    string `json:"legal_person,omitempty"`
	CompanyAddress string `json:"company_address,omitempty"`
	CertValidDate  string `json:"cert_valid_date,omitempty"`
}

type AssistProveInfo struct {
	MicroBizType     string `json:"micro_biz_type,omitempty"`
	StoreName        string `json:"store_name,omitempty"`
	StoreAddressCode string `json:"store_address_code,omitempty"`
	StoreAddress     string `json:"store_address,omitempty"`
	StoreHeaderCopy  string `json:"store_header_copy,omitempty"`
	StoreIndoorCopy  string `json:"store_indoor_copy,omitempty"`
}

type SpecialOperationList struct {
	CategoryId        string   `json:"category_id,omitempty"`
	OperationCopyList []string `json:"operation_copy_list,omitempty"`
}

type IdentificationInfo struct {
	IdentificationType      string `json:"identification_type"`
	IdentificationName      string `json:"identification_name"`
	IdentificationNumber    string `json:"identification_number"`
	IdentificationValidDate string `json:"identification_valid_date"`
	IdentificationFrontCopy string `json:"identification_front_copy"`
	IdentificationBackCopy  string `json:"identification_back_copy"`
}

// APIUrl ReqSubmitApplication APIURL
func (this ReqSubmitApplication) APIUrl() string {
	return "/v3/apply4subject/applyment"
}

// Method ReqSubmitApplication Method
func (this ReqSubmitApplication) Method() string {
	return "POST"
}

// Params ReqSubmitApplication Params
func (this ReqSubmitApplication) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// RawJsonStr ReqSubmitApplication RawJsonStr
func (this ReqSubmitApplication) RawJsonStr() string {
	jsons, errs := json.Marshal(this)
	if errs != nil {
		fmt.Println(errs.Error())
		return "Error"
	}
	fmt.Println(string(jsons))
	return string(jsons)
}
func (this ReqSubmitApplication) CheckCertificateInfoCertType(mk string) bool {
	switch mk {
	case CERTIFICATE_TYPE_2388:
		return true
	case CERTIFICATE_TYPE_2389:
		return true
	case CERTIFICATE_TYPE_2390:
		return true
	case CERTIFICATE_TYPE_2391:
		return true
	case CERTIFICATE_TYPE_2392:
		return true
	case CERTIFICATE_TYPE_2393:
		return true
	case CERTIFICATE_TYPE_2394:
		return true
	case CERTIFICATE_TYPE_2395:
		return true
	case CERTIFICATE_TYPE_2396:
		return true
	case CERTIFICATE_TYPE_2397:
		return true
	case CERTIFICATE_TYPE_2398:
		return true
	case CERTIFICATE_TYPE_2399:
		return true
	case CERTIFICATE_TYPE_2400:
		return true
	default:
		return false
	}
}

type RespSubmitApplication struct {
	ApplymentId string `json:"applyment_id"`
}
