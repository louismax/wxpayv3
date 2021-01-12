package wxpayv3

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// MerchantsIntoPieces 普通服务商商户进件
func (this *Client) MerchantsIntoPieces(param ReqMerchantsIntoPieces) (interface{}, error) {
	if param.BusinessCode == "" {
		return nil, errors.New("业务申请编号不能为空")
	}
	if param.ContactInfo.ContactName == "" {
		return nil, errors.New("超级管理员姓名不能为空")
	}
	if param.ContactInfo.ContactIdNumber == "" && param.ContactInfo.Openid == "" {
		return nil, errors.New("超级管理员身份证件号码或Openid不能为空，二选一")
	}
	if param.ContactInfo.MobilePhone == "" {
		return nil, errors.New("超级管理员手机号不能为空")
	}

	if param.SubjectInfo.SubjectType != SUBJECT_TYPE_MICRO {
		return nil, errors.New("主体类型无效，目前仅支持小微商户")
	}
	//仅验证小微辅助证明材料
	if param.SubjectInfo.MicroBizInfo.MicroBizYype != MICRO_TYPE_STORE && param.SubjectInfo.MicroBizInfo.MicroBizYype != MICRO_TYPE_MOBILE && param.SubjectInfo.MicroBizInfo.MicroBizYype != MICRO_TYPE_ONLINE {
		return nil, errors.New("小微经营类型无效")
	}
	if param.SubjectInfo.MicroBizInfo.MicroBizYype == MICRO_TYPE_STORE {
		if param.SubjectInfo.MicroBizInfo.MicroStoreInfo.MicroName == "" {
			return nil, errors.New("门店名称无效")
		}
		if param.SubjectInfo.MicroBizInfo.MicroStoreInfo.MicroAddressCode == "" {
			return nil, errors.New("门店省市编码无效")
		}
		if param.SubjectInfo.MicroBizInfo.MicroStoreInfo.MicroAddress == "" {
			return nil, errors.New("门店地址无效")
		}
		if param.SubjectInfo.MicroBizInfo.MicroStoreInfo.StoreEntrancePic == "" {
			return nil, errors.New("门店门头照片无效")
		}
		if param.SubjectInfo.MicroBizInfo.MicroStoreInfo.MicroIndoorCopy == "" {
			return nil, errors.New("店内环境照片无效")
		}
	}
	if param.SubjectInfo.MicroBizInfo.MicroBizYype == MICRO_TYPE_MOBILE {
		if param.SubjectInfo.MicroBizInfo.MicroMobileInfo.MicroMobileName == "" {
			return nil, errors.New("经营/服务名称无效")
		}
		if param.SubjectInfo.MicroBizInfo.MicroMobileInfo.MicroMobileCity == "" {
			return nil, errors.New("经营/服务所在地省市无效")
		}
		if param.SubjectInfo.MicroBizInfo.MicroMobileInfo.MicroMobileAddress == "" {
			return nil, errors.New("经营/服务所在地无效")
		}
		if param.SubjectInfo.MicroBizInfo.MicroMobileInfo.MicroMobilePics == "" {
			return nil, errors.New("经营/服务现场照片无效")
		}
	}
	if param.SubjectInfo.MicroBizInfo.MicroBizYype == MICRO_TYPE_ONLINE {
		if param.SubjectInfo.MicroBizInfo.MicroOnlineInfo.MicroOnlineStore == "" {
			return nil, errors.New("线上店铺名称无效")
		}
		if param.SubjectInfo.MicroBizInfo.MicroOnlineInfo.MicroEcName == "" {
			return nil, errors.New("电商平台名称无效")
		}
		if param.SubjectInfo.MicroBizInfo.MicroOnlineInfo.MicroQrcode == "" && param.SubjectInfo.MicroBizInfo.MicroOnlineInfo.MicroLink == "" {
			return nil, errors.New("店铺二维码/店铺链接二选一必填")
		}
	}

	if param.SubjectInfo.IdentityInfo.IdDocType != IDENTIFICATION_TYPE_IDCARD {
		return nil, errors.New("经营者证件类型无效，仅支持身份证")
	}
	if param.SubjectInfo.IdentityInfo.IdCardInfo.IdCardCopy == "" {
		return nil, errors.New("身份证人像面照片无效")
	}
	if param.SubjectInfo.IdentityInfo.IdCardInfo.IdCardNational == "" {
		return nil, errors.New("身份证国徽面照片无效")
	}
	if param.SubjectInfo.IdentityInfo.IdCardInfo.IdCardName == "" {
		return nil, errors.New("身份证姓名无效")
	}
	if param.SubjectInfo.IdentityInfo.IdCardInfo.IdCardNumber == "" {
		return nil, errors.New("身份证号码无效")
	}
	if param.SubjectInfo.IdentityInfo.IdCardInfo.CardPeriodBegin == "" {
		return nil, errors.New("身份证有效期开始时间无效")
	}
	if param.SubjectInfo.IdentityInfo.IdCardInfo.CardPeriodEnd == "" {
		return nil, errors.New("身份证有效期结束时间无效")
	}

	if param.BusinessInfo.MerchantShortname == "" {
		return nil, errors.New("商户简称不能为空")
	}
	if param.BusinessInfo.ServicePhone == "" {
		return nil, errors.New("客服电话不能为空")
	}

	if param.SettlementInfo.SettlementId == "" {
		return nil, errors.New("入驻结算规则ID不能为空")
	}
	if param.SettlementInfo.QualificationType == "" {
		return nil, errors.New("所属行业不能为空")
	}

	if param.BankAccountInfo.BankAccountType != BANK_ACCOUNT_TYPE_PERSONAL {
		return nil, errors.New("收款银行账户类型无效")
	}
	if param.BankAccountInfo.AccountName == "" {
		return nil, errors.New("收款银行开户名称无效")
	}
	if param.BankAccountInfo.AccountBank == "" {
		return nil, errors.New("收款银行开户银行无效")
	}
	if param.BankAccountInfo.BankAddressCode == "" {
		return nil, errors.New("收款银行开户银行省市编码无效")
	}
	if param.BankAccountInfo.AccountBank == "其他银行" {
		if param.BankAccountInfo.BankBranchId == "" && param.BankAccountInfo.BankName == "" {
			return nil, errors.New("收款银行开户银行支行信息无效无效")
		}
	}
	if param.BankAccountInfo.AccountNumber == "" {
		return nil, errors.New("银行账号无效")
	}

	var err error
	param.ContactInfo.ContactName, err = this.RsaOAEPEncrypt(param.ContactInfo.ContactName)
	if err != nil {
		return nil, errors.New("用户姓名加密错误！")
	}
	if param.ContactInfo.ContactIdNumber != "" {
		param.ContactInfo.ContactIdNumber, err = this.RsaOAEPEncrypt(param.ContactInfo.ContactIdNumber)
		if err != nil {
			return nil, errors.New("身份证信息加密错误！")
		}
	}
	if param.ContactInfo.Openid != "" {
		param.ContactInfo.Openid, err = this.RsaOAEPEncrypt(param.ContactInfo.Openid)
		if err != nil {
			return nil, errors.New("Openid加密错误！")
		}
	}
	param.ContactInfo.MobilePhone, err = this.RsaOAEPEncrypt(param.ContactInfo.MobilePhone)
	if err != nil {
		return nil, errors.New("联系手机加密错误！")
	}
	if param.ContactInfo.ContactEmail != "" {
		param.ContactInfo.ContactEmail, err = this.RsaOAEPEncrypt(param.ContactInfo.ContactEmail)
		if err != nil {
			return nil, errors.New("联系邮箱加密错误！")
		}
	}
	param.SubjectInfo.IdentityInfo.IdCardInfo.IdCardName, err = this.RsaOAEPEncrypt(param.SubjectInfo.IdentityInfo.IdCardInfo.IdCardName)
	if err != nil {
		return nil, errors.New("经营者身份证姓名加密错误！")
	}
	param.SubjectInfo.IdentityInfo.IdCardInfo.IdCardNumber, err = this.RsaOAEPEncrypt(param.SubjectInfo.IdentityInfo.IdCardInfo.IdCardNumber)
	if err != nil {
		return nil, errors.New("经营者身份证号码加密错误！")
	}
	param.BankAccountInfo.AccountName, err = this.RsaOAEPEncrypt(param.BankAccountInfo.AccountName)
	if err != nil {
		return nil, errors.New("收款银行卡开户名称加密错误！")
	}
	param.BankAccountInfo.AccountNumber, err = this.RsaOAEPEncrypt(param.BankAccountInfo.AccountNumber)
	if err != nil {
		return nil, errors.New("收款银行卡号码加密错误！")
	}

	result := RespMerchantsIntoPieces{}
	rqs, err := this.doRequest(param, &result)
	if err != nil {
		return result, err
	}
	fmt.Println(fmt.Sprintf("原始返回参数:%+v", rqs))

	if strings.Contains(rqs, "code") {
		errmsg := SysError{}
		err = json.Unmarshal([]byte(rqs), &errmsg)
		if err != nil {
			return nil, err
		}
		return errmsg, errors.New("Fail")
	} else {
		return result, nil
	}
}

//GetStatusRepairOrderForBusCode 商户申请单号获取申请单信息
func (this *Client) GetStatusRepairOrderForBusCode(param ReqGetStatusRepairOrderForBusCode) (interface{}, error) {
	if param.BusinessCode == "" {
		return nil, errors.New("商户申请单号不能为空！")
	}

	result := RespGetStatusRepairOrder{}
	rqs, err := this.doRequest(param, &result)
	if err != nil {
		return nil, err
	}
	fmt.Println(fmt.Sprintf("原始返回参数:%+v", rqs))

	return result, nil
}

//GetStatusRepairOrderForApplyCode 微信申请单号获取申请单信息
func (this *Client) GetStatusRepairOrderForApplyCode(param ReqGetStatusRepairOrderForApplyCode) (interface{}, error) {
	if param.ApplymentId == "" {
		return nil, errors.New("微信申请单号不能为空！")
	}

	result := RespGetStatusRepairOrder{}
	rqs, err := this.doRequest(param, &result)
	if err != nil {
		return nil, err
	}
	fmt.Println(fmt.Sprintf("原始返回参数:%+v", rqs))

	return result, nil
}
