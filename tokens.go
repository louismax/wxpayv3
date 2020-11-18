package wxpayv3

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// GetTokens 获取小程序授权凭证
func (this *Client) GetTokens(userid, wxogid string) (RespGetTokens, error) {
	result := RespGetTokens{}
	if userid == "" {
		return result, errors.New("用户id不能为空！")
	}
	if wxogid == "" {
		return result, errors.New("机构ID不能为空！")
	}

	param := GetTokens{}
	param.Scene = "WEBSESSION" //小程序默认
	param.Web_init_data.Out_user_id = userid
	param.Web_init_data.Organization_id = wxogid

	rqs, err := this.doRequest(param, &result)
	if err != nil {
		return result, err
	}

	if strings.Contains(rqs, "code") {
		errmsg := SysError{}

		err = json.Unmarshal([]byte(rqs), &errmsg)
		if err != nil {
			return result, err
		}
		fmt.Println(fmt.Sprintf("%+v", errmsg))
		if errmsg.Code == "SYSTEM_ERROR" {
			return result, errors.New("获取授权凭证失败!")
		} else if errmsg.Code == "SIGN_ERROR" {
			return result, errors.New("系统错误，签名失败!")
		} else {
			return result, nil
		}
	}
	return result, nil
}

func (this *Client) GetPresign_Token(param Presign_Token) (RespPresign_Token, error) {
	result := RespPresign_Token{}
	if param.Business_name == "" {
		param.Business_name = "K12"
	}
	if param.Facepay_user.Out_user_id == "" {
		return result, errors.New("刷脸用户ID不能为空！")
	}
	if param.Facepay_user.Identification_name == "" {
		return result, errors.New("刷脸用户姓名不能为空！")
	}
	if param.Facepay_user.Organization_id == "" {
		return result, errors.New("微信机构ID不能为空！")
	}
	if param.Facepay_user.Phone == "" {
		return result, errors.New("手机号码不能为空！")
	}
	//Louis 2020年11月17日17:04:08 用户证件非必传
	//if param.Facepay_user.Identification.Identification_type == "" {
	//	return result, errors.New("刷脸用户证件类型不能为空！")
	//}
	//if param.Facepay_user.Identification.Identification_number == "" {
	//	return result, errors.New("刷脸用户证件号码不能为空！")
	//}
	if param.Limit_bank_card.Bank_card_number == "" {
		return result, errors.New("银行卡号不能为空！")
	}
	if param.Limit_bank_card.Identification_name == "" {
		return result, errors.New("开卡人姓名不能为空！")
	}
	if param.Limit_bank_card.Identification.Identification_type == "" {
		return result, errors.New("开卡人证件类型不能为空！")
	}
	if param.Limit_bank_card.Identification.Identification_number == "" {
		return result, errors.New("开卡人证件号码不能为空！")
	}
	if this.PFSerialno == "" {
		return result, errors.New("请先初始化平台证书")
	}

	var err error
	param.Facepay_user.Identification_name, err = this.RsaOAEPEncrypt(param.Facepay_user.Identification_name)
	if err != nil {
		return result, errors.New("刷脸用户姓名加密错误！")
	}
	param.Facepay_user.Phone, err = this.RsaOAEPEncrypt(param.Facepay_user.Phone)
	if err != nil {
		return result, errors.New("手机号码加密错误！")
	}
	//param.Facepay_user.Identification.Identification_number, err = this.RsaOAEPEncrypt(param.Facepay_user.Identification.Identification_number)
	//if err != nil {
	//	return result, errors.New("刷脸用户证件号码加密错误！")
	//}
	param.Limit_bank_card.Bank_card_number, err = this.RsaOAEPEncrypt(param.Limit_bank_card.Bank_card_number)
	if err != nil {
		return result, errors.New("银行卡号加密错误！")
	}
	param.Limit_bank_card.Identification_name, err = this.RsaOAEPEncrypt(param.Limit_bank_card.Identification_name)
	if err != nil {
		return result, errors.New("开户人姓名加密错误！")
	}
	param.Limit_bank_card.Identification.Identification_number, err = this.RsaOAEPEncrypt(param.Limit_bank_card.Identification.Identification_number)
	if err != nil {
		return result, errors.New("开户人证件号码加密错误！")
	}

	rqs, err := this.doRequest(param, &result)
	if err != nil {
		return result, err
	}

	if strings.Contains(rqs, "code") {
		errmsg := SysError{}

		err = json.Unmarshal([]byte(rqs), &errmsg)
		if err != nil {
			return result, err
		}
		fmt.Println(fmt.Sprintf("%+v", errmsg))
		if errmsg.Code == "SYSTEM_ERROR" {
			return result, errors.New("获取预签约会话失败!")
		} else if errmsg.Code == "SIGN_ERROR" {
			return result, errors.New("系统错误，签名失败!")
		} else if errmsg.Code == "PARAM_ERROR" {
			return result, errors.New("系统错误，请求的参数错误!")
		} else {
			return result, errors.New(fmt.Sprintf("系统错误:resp：%+v!", errmsg))
		}
	}
	return result, nil
}
