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
			param.Business_name = "K12"
			param.Facepay_user.Out_user_id = ""
			param.Facepay_user.Identification_name = ""
			param.Facepay_user.Organization_id = ""
			param.Facepay_user.Identification.Identification_type = "IDCARD"
			param.Facepay_user.Identification.Identification_number = ""
			param.Facepay_user.Phone = ""
			param.Limit_bank_card.Bank_card_number = ""
			param.Limit_bank_card.Identification_name = ""
			param.Limit_bank_card.Identification.Identification_type = "IDCARD"
			param.Limit_bank_card.Identification.Identification_number = ""
			param.Limit_bank_card.Valid_thru = ""

			res, err := client.GetPresign_Token(param)
			if err != nil {
				fmt.Println(fmt.Sprintf("调用失败,err:%+v", err))
			}
			fmt.Println(fmt.Sprintf("调用结果,res:%+v", res))
			t.Log(res)
		}

	}
}
