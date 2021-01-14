package wxpayv3

import (
	"fmt"
	"testing"
)

func TestClient_MerchantsIntoPieces(t *testing.T) {
	client, err := New("", "", "apiclient_key.pem", "apiclient_cert.pem")
	if err != nil {
		t.Log(fmt.Sprintf("V3客户端初始化失败,err:%+v", err))
		return
	}
	param := CertifiCates{}
	res, err := client.CertifiCates(param)
	if err != nil {
		t.Log(fmt.Sprintf("获取平台证书失败,err:%+v", err))
		return
	} else {
		//平台证书解密
		Cipher := res.(SystemOauthTokenRsp)

		cretinfoint, err := Cipher.CertificateDecryption("")
		if err != nil {
			t.Log(fmt.Sprintf("平台证书解密失败,err:%+v", err))
			return
		} else {
			cretinfo := cretinfoint.(CertificateInfo)
			//平台证书初始化
			client.InitCertificate(cretinfo)

			param2 := ReqMerchantsIntoPieces{
				BusinessCode: GetGUID(),
				ContactInfo: &ContactInfo{
					ContactName:     "",
					ContactIdNumber: "",
					MobilePhone:     "",
					ContactEmail:    "",
				},
				SubjectInfo: &SubjectInfo{
					SubjectType: SUBJECT_TYPE_MICRO,
					MicroBizInfo: &MicroBizInfo{
						MicroBizYype: MICRO_TYPE_STORE,
						MicroStoreInfo: &MicroStoreInfo{
							MicroName:        "",
							MicroAddressCode: "",
							MicroAddress:     "",
							StoreEntrancePic: "",
							MicroIndoorCopy:  "",
						},
					},
					IdentityInfo: &IdentityInfo{
						IdDocType: IDENTIFICATION_TYPE_IDCARD,
						IdCardInfo: &IdCardInfo{
							IdCardCopy:      "",
							IdCardNational:  "",
							IdCardName:      "",
							IdCardNumber:    "",
							CardPeriodBegin: "2020-01-01",
							CardPeriodEnd:   "长期",
						},
					},
				},
				BusinessInfo: &BusinessInfo_MI{
					MerchantShortname: "",
					ServicePhone:      "",
				},
				SettlementInfo: &SettlementInfo{
					SettlementId:      "",
					QualificationType: "",
				},
				BankAccountInfo: &BankAccountInfo{
					BankAccountType: BANK_ACCOUNT_TYPE_PERSONAL,
					AccountName:     "",
					AccountBank:     "",
					BankAddressCode: "",
					AccountNumber:   "",
				},
			}

			res2, err := client.MerchantsIntoPieces(param2)
			if err != nil {
				t.Log(fmt.Sprintf("RESP:%+v", res2))
				t.Log(fmt.Sprintf("失败,err:%+v", err))
				return
			}
			t.Log(fmt.Sprintf("result:%+v", res2))
		}

	}
}

func TestClient_GetStatusRepairOrderForApplyCode(t *testing.T) {
	client, err := New("", "", "apiclient_key.pem", "apiclient_cert.pem")
	if err != nil {
		t.Log(fmt.Sprintf("V3客户端初始化失败,err:%+v", err))
		return
	}
	param := ReqGetStatusRepairOrderForApplyCode{
		ApplymentId: "",
	}
	res, err := client.GetStatusRepairOrderForApplyCode(param)
	if err != nil {
		t.Log(fmt.Sprintf("RESP:%+v", res))
		t.Log(fmt.Sprintf("失败,err:%+v", err))
		return
	}
	t.Log(fmt.Sprintf("result:%+v", res))
}

func TestClient_GetStatusRepairOrderForBusCode(t *testing.T) {
	client, err := New("1501889641", "1433BE83CE5A9D6022972F4B144A714C598CFADF", "apiclient_key.pem", "apiclient_cert.pem")
	if err != nil {
		t.Log(fmt.Sprintf("V3客户端初始化失败,err:%+v", err))
		return
	}
	param := ReqGetStatusRepairOrderForBusCode{
		BusinessCode: "3a846fdca4e59c3e36bbde799bd04526",
	}
	res, err := client.GetStatusRepairOrderForBusCode(param)
	if err != nil {
		t.Log(fmt.Sprintf("RESP:%+v", res))
		t.Log(fmt.Sprintf("失败,err:%+v", err))
		return
	}
	t.Log(fmt.Sprintf("result:%+v", res))
}
