package wxpayv3

import (
	"errors"
	"fmt"
)

//SubmitApplication 提交申请单
func (this *Client) SubmitApplication(param ReqSubmitApplication) (interface{}, error) {
	if param.BusinessCode == "" {
		return nil, errors.New("业务申请编号不能为空")
	}
	if param.ContactInfo.Name == "" {
		return nil, errors.New("联系人姓名不能为空")
	}
	if param.ContactInfo.Mobile == "" {
		return nil, errors.New("联系人手机号不能为空")
	}
	if param.ContactInfo.IdCardNumber == "" {
		return nil, errors.New("联系人身份证号号码不能为空")
	}
	if param.SubjectInfo.SubjectType != SUBJECT_TYPE_ENTERPRISE && param.SubjectInfo.SubjectType != SUBJECT_TYPE_INSTITUTIONS_CLONED && param.SubjectInfo.SubjectType != SUBJECT_TYPE_INDIVIDUAL && param.SubjectInfo.SubjectType != SUBJECT_TYPE_OTHERS && param.SubjectInfo.SubjectType != SUBJECT_TYPE_MICRO {
		return nil, errors.New("主体类型无效")
	}

	if param.IdentificationInfo.IdentificationType != IDENTIFICATION_TYPE_IDCARD && param.IdentificationInfo.IdentificationType != IDENTIFICATION_TYPE_OVERSEA_PASSPORT && param.IdentificationInfo.IdentificationType != IDENTIFICATION_TYPE_HONGKONG_PASSPORT && param.IdentificationInfo.IdentificationType != IDENTIFICATION_TYPE_MACAO_PASSPORT && param.IdentificationInfo.IdentificationType != IDENTIFICATION_TYPE_TAIWAN_PASSPORT {
		return nil, errors.New("法人证件类型无效")
	}
	if param.IdentificationInfo.IdentificationName == "" {
		return nil, errors.New("法人证件姓名不能为空")
	}
	if param.IdentificationInfo.IdentificationNumber == "" {
		return nil, errors.New("法人证件号码不能为空")
	}
	if param.IdentificationInfo.IdentificationValidDate == "" {
		return nil, errors.New("法人证件有效日期不能为空")
	}
	if param.IdentificationInfo.IdentificationFrontCopy == "" {
		return nil, errors.New("法人证件正面照片不能为空")
	}
	if param.IdentificationInfo.IdentificationBackCopy == "" {
		return nil, errors.New("法人证件反面照片不能为空")
	}

	//主体类型为企业或个体户字段验证
	if param.SubjectInfo.SubjectType == SUBJECT_TYPE_ENTERPRISE || param.SubjectInfo.SubjectType == SUBJECT_TYPE_INDIVIDUAL {
		if param.SubjectInfo.BusinessLicenceInfo.LicenceNumber == "" {
			return nil, errors.New("营业执照注册号不能为空")
		}
		if param.SubjectInfo.BusinessLicenceInfo.LicenceCopy == "" {
			return nil, errors.New("营业执照照片不能为空")
		}
		if param.SubjectInfo.BusinessLicenceInfo.MerchantName == "" {
			return nil, errors.New("营业执照商户名称不能为空")
		}
		if param.SubjectInfo.BusinessLicenceInfo.LegalPerson == "" {
			return nil, errors.New("营业执照法人姓名不能为空")
		}
		if param.SubjectInfo.BusinessLicenceInfo.CompanyAddress == "" {
			return nil, errors.New("营业执注册地址不能为空")
		}
		if param.SubjectInfo.BusinessLicenceInfo.LicenceValidDate == "" {
			return nil, errors.New("营业执有效期不能为空")
		}
	}
	//主体类型为事业单位或其他组织 字段验证
	if param.SubjectInfo.SubjectType == SUBJECT_TYPE_INSTITUTIONS_CLONED || param.SubjectInfo.SubjectType == SUBJECT_TYPE_OTHERS {
		if !param.CheckCertificateInfoCertType(param.SubjectInfo.CertificateInfo.CertType) {
			return nil, errors.New("登记证书类型无效")
		}
		if param.SubjectInfo.CertificateInfo.CertNumber == "" {
			return nil, errors.New("登记证书编号不能为空")
		}
		if param.SubjectInfo.CertificateInfo.CertCopy == "" {
			return nil, errors.New("登记证书照片不能为空")
		}
		if param.SubjectInfo.CertificateInfo.MerchantName == "" {
			return nil, errors.New("登记证书商户名称不能为空")
		}
		if param.SubjectInfo.CertificateInfo.LegalPerson == "" {
			return nil, errors.New("登记证书法人姓名不能为空")
		}
		if param.SubjectInfo.CertificateInfo.CompanyAddress == "" {
			return nil, errors.New("登记证书注册地址不能为空")
		}
		if param.SubjectInfo.CertificateInfo.CertValidDate == "" {
			return nil, errors.New("登记证书有效日期不能为空")
		}
		if param.SubjectInfo.SubjectType == SUBJECT_TYPE_INSTITUTIONS_CLONED {
			if param.SubjectInfo.CompanyProveCopy == "" {
				return nil, errors.New("事业单位证明函照片不能为空")
			}
		}
	}

	//主体为小微商户字段验证
	if param.SubjectInfo.SubjectType == SUBJECT_TYPE_MICRO {
		if param.SubjectInfo.AssistProveInfo.MicroBizType != MICRO_TYPE_STORE && param.SubjectInfo.AssistProveInfo.MicroBizType != MICRO_TYPE_MOBILE && param.SubjectInfo.AssistProveInfo.MicroBizType != MICRO_TYPE_ONLINE {
			return nil, errors.New("小微经营类型无效")
		}
		if param.SubjectInfo.AssistProveInfo.StoreName == "" {
			return nil, errors.New("小微门店名称不能为空")
		}
		if param.SubjectInfo.AssistProveInfo.StoreAddressCode == "" {
			return nil, errors.New("小微门店省市编码不能为空")
		}
		if param.SubjectInfo.AssistProveInfo.StoreAddress == "" {
			return nil, errors.New("小微门店地址不能为空")
		}
		if param.SubjectInfo.AssistProveInfo.StoreHeaderCopy == "" {
			return nil, errors.New("小微门店门头照片不能为空")
		}
		if param.SubjectInfo.AssistProveInfo.StoreIndoorCopy == "" {
			return nil, errors.New("小微店内环境照片不能为空")
		}
	}

	result := RespSubmitApplication{}
	rqs, err := this.doRequest(param, &result)
	if err != nil {
		return result, err
	}

	fmt.Println(fmt.Sprintf("原始返回参数:%+v", rqs))

	return result, nil
}
