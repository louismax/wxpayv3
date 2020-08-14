package wxpayv3

import (
	"fmt"
	"testing"
)

func TestClient_GetPresign_Token(t *testing.T) {
	client, err := New("", "", "apiclient_key.pem", "apiclient_cert.pem")
	if err != nil {
		fmt.Println(fmt.Sprintf("V3客户端初始化失败,err:%+v", err))
	}
	param1 := CertifiCates{}
	res, err := client.CertifiCates(param1)
	if err != nil {
		fmt.Println(fmt.Sprintf("获取平台证书失败,err:%+v", err))
	} else {
		Cipher := res.(SystemOauthTokenRsp)
		cretinfoint, err := Cipher.CertificateDecryption("")
		if err != nil {
			fmt.Println(fmt.Sprintf("平台证书解密失败,err:%+v", err))
		} else {
			cretinfo := cretinfoint.(CertificateInfo)

			//平台证书初始化
			client.InitCertificate(cretinfo)

			param := Presign_Token{}
			param.Business_name = "K12"                                         //固定值
			param.Facepay_user.Out_user_id = ""                                 //商户侧userid
			param.Facepay_user.Identification_name = ""                         //姓名
			param.Facepay_user.Organization_id = ""                             //所属学校机构id
			param.Facepay_user.Identification.Identification_type = "IDCARD"    //固定值，必须身份证
			param.Facepay_user.Identification.Identification_number = ""        //身份证号
			param.Facepay_user.Phone = ""                                       //手机号
			param.Limit_bank_card.Bank_card_number = ""                         //银行卡号
			param.Limit_bank_card.Identification_name = ""                      //开卡人姓名
			param.Limit_bank_card.Identification.Identification_type = "IDCARD" //固定值，必须身份证
			param.Limit_bank_card.Identification.Identification_number = ""     //开卡人身份证号
			param.Limit_bank_card.Valid_thru = ""                               //银行卡有效期,部分银行需要

			res, err := client.GetPresign_Token(param)
			if err != nil {
				fmt.Println(fmt.Sprintf("调用失败,err:%+v", err))
			}
			fmt.Println(fmt.Sprintf("调用结果,res:%+v", res))
			t.Log(res)
		}

	}
}
