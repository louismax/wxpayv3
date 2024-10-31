package wxpayv3

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	//client, err := NewClient(
	//	InjectWechatPayParameterUseCertPath("1234", "1234", "apiclient_key.pem", "apiclient_cert.pem"),
	//)
	//if err != nil {
	//	t.Log(err)
	//	return
	//}
	//下载证书
	//resp, err := client.Certificate()
	//if err != nil {
	//	t.Log(err)
	//	return
	//}
	//
	//for _, v := range resp.Data {
	//	t.Logf("%+v", *v)
	//	t.Logf("%+v", *v.EncryptCertificate)
	//}

	//获取结算账号
	//resp, err := client.QuerySettlementAccount("1609337198")
	//if err != nil {
	//	t.Log(err)
	//	return
	//}
	//t.Log(resp)

	//上传图片
	//resp, err := client.UploadImage("./1.jpg")
	//if err != nil {
	//	t.Log(err)
	//	return
	//}
	//t.Logf("%+v", resp)

	//退款
	//uid, _ := utils.GenerateNonce()
	//resp, err := client.PaymentRefund(custom.ReqPaymentRefund{
	//	SubMchid:      "123",
	//	TransactionId: "4200001126202108129763281234",
	//	OutRefundNo:   uid,
	//	Reason:        "API调试",
	//	Amount: custom.PaymentRefundAmount{
	//		Refund:   1,
	//		Total:    1,
	//		Currency: "CNY",
	//	},
	//})
	//if err != nil {
	//	t.Log(err)
	//	return
	//}
	//t.Log(resp)
}
