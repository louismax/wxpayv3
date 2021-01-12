package wxpayv3

import (
	"fmt"
	"testing"
)

func TestClient_SubmitApplication(t *testing.T) {
	client, err := New("", "", "apiclient_key.pem", "apiclient_cert.pem")
	if err != nil {
		t.Log(fmt.Sprintf("V3客户端初始化失败,err:%+v", err))
		return
	}

	param := ReqSubmitApplication{
		BusinessCode: GetGUID(),
		ContactInfo: ContactInfo{
			Name:         "",
			Mobile:       "",
			IdCardNumber: "",
		},
		SubjectInfo: SubjectInfo{
			SubjectType: SUBJECT_TYPE_MICRO,
			BusinessLicenceInfo: BusinessLicenceInfo{
				LicenceNumber:    "",
				LicenceCopy:      "",
				MerchantName:     "",
				LegalPerson:      "",
				CompanyAddress:   "",
				LicenceValidDate: "[\"1970-01-01\",\"forever\"]",
			},
			AssistProveInfo: AssistProveInfo{
				MicroBizType:     MICRO_TYPE_MOBILE,
				StoreName:        "",
				StoreAddressCode: "",
				StoreAddress:     "",
				StoreHeaderCopy:  "",
				StoreIndoorCopy:  "",
			},
		},
		IdentificationInfo: IdentificationInfo{
			IdentificationType:      IDENTIFICATION_TYPE_IDCARD,
			IdentificationName:      "",
			IdentificationNumber:    "",
			IdentificationValidDate: "[\"1970-01-01\",\"forever\"]",
			IdentificationFrontCopy: "",
			IdentificationBackCopy:  "",
		},
	}

	res, err := client.SubmitApplication(param)
	if err != nil {
		t.Log(fmt.Sprintf("RESP:%+v", res))
		t.Log(fmt.Sprintf("失败,err:%+v", err))
		return
	}
	t.Log(fmt.Sprintf("result:%+v", res))

}
