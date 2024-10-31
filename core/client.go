package core

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/custom"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client is Client
type Client interface {
	// Authorization 获取签名Authorization，由认证类型和签名信息组成
	Authorization(httpMethod string, urlString string, body []byte) (string, error)
	// Certificate 获取平台证书
	Certificate() (*custom.CertificateResp, error)
	// SetClientPlatformCert 设置平台证书
	SetClientPlatformCert(certificateStr string) error
	// RsaEncryptByPrivateKey 使用商户私钥加密敏感数据
	RsaEncryptByPrivateKey(origData []byte) (string, error)
	// RsaDecryptByPrivateKey 使用商户私钥解密敏感数据
	RsaDecryptByPrivateKey(ciphertext string) (string, error)
	// RsaEncryptByPublicKey 使用平台公钥加密敏感数据
	RsaEncryptByPublicKey(plaintext string) (string, error)
	// Decrypt 通知密文数据使用V3Key解密 （AES_256_GCM）
	Decrypt(algorithm string, cipherText string, associatedData string, nonce string) ([]byte, error)

	// UploadImage 上传图片（获取MediaId）
	UploadImage(filePath string) (*custom.RespUploadImage, error)

	DownloadBill(downloadUrl string) ([]byte, error)

	//IncomingSubmitApplication 提交进件申请单
	IncomingSubmitApplication(data custom.ReqIncomingSubmitApplication) (*custom.RespIncomingSubmitApplication, error)
	//ModifySettlement 修改结算账号
	ModifySettlement(subMchid string, data custom.ReqModifySettlement) error
	// QuerySettlementAccount 查询结算账户
	QuerySettlementAccount(subMchid string) (*custom.SettlementAccount, error)
	// GetStatusRepairOrderByBusinessCode 通过业务申请编号查询申请状态
	GetStatusRepairOrderByBusinessCode(businessCode string) (*custom.RespGetStatusRepairOrder, error)
	// GetStatusRepairOrderByApplymentId 通过申请单号查询申请状态
	GetStatusRepairOrderByApplymentId(applymentId string) (*custom.RespGetStatusRepairOrder, error)

	//InitiateProfitSharing 发起分账请求
	InitiateProfitSharing(data custom.ReqInitiateProfitSharing) (*custom.RespInitiateProfitSharing, error)
	//QueryProfitSharingResult 查询分账结果
	QueryProfitSharingResult(subMchid, transactionId, outOrderNo string) (*custom.RespQueryProfitSharingResult, error)
	//InitiateProfitSharingReturnOrders 请求分账回退
	InitiateProfitSharingReturnOrders(data custom.ReqInitiateProfitSharingReturnOrders) (*custom.RespInitiateProfitSharingReturnOrders, error)
	//QueryProfitSharingReturnOrders 查询分账回退结果
	QueryProfitSharingReturnOrders(subMchid, outReturnNo, outOrderNo string) (*custom.RespQueryProfitSharingReturnOrders, error)
	//UnfreezeRemainingFunds 解冻剩余资金
	UnfreezeRemainingFunds(data custom.ReqUnfreezeRemainingFunds) (*custom.RespUnfreezeRemainingFunds, error)
	//QueryRemainingFrozenAmount 查询订单待分金额
	QueryRemainingFrozenAmount(transactionId string) (*custom.RespQueryRemainingFrozenAmount, error)
	//QueryMaximumSplitRatio 查询子商户最大分账比例
	QueryMaximumSplitRatio(subMchid string) (*custom.RespQueryMaximumSplitRatio, error)
	//AddProfitSharingReceiver 添加分账接收方
	AddProfitSharingReceiver(data custom.ReqAddProfitSharingReceiver) (*custom.RespAddProfitSharingReceiver, error)
	//DeleteProfitSharingReceiver 删除分账接收方
	DeleteProfitSharingReceiver(data custom.ReqDeleteProfitSharingReceiver) (*custom.RespDeleteProfitSharingReceiver, error)
	//ApplyProfitSharingBill 申请分账账单
	ApplyProfitSharingBill(billDate, subMchid, tarType string) (*custom.RespApplyTransactionBill, error)

	//PaymentQueryOrderByTransactionId 查询订单-通过微信订单号(兼容服务商模式、直连商户模式)
	PaymentQueryOrderByTransactionId(transactionId, mchID string, subMchId ...string) (*custom.ReqPaymentQueryOrder, error)
	//PaymentQueryOrderByOutTradeNo 查询订单-通过商户订单号(兼容服务商模式、直连商户模式)
	PaymentQueryOrderByOutTradeNo(outTradeNo, mchID string, subMchId ...string) (*custom.ReqPaymentQueryOrder, error)
	//PaymentRefund 直连商户退款
	PaymentRefund(data custom.ReqPaymentRefund) (*custom.RespPaymentRefund, error)
	//PaymentRefundForPartner 服务商退款
	PaymentRefundForPartner(data custom.ReqPaymentRefundForPartner) (*custom.RespPaymentRefund, error)
	//ApplyTransactionBill //申请交易账单
	ApplyTransactionBill(billDate, subMchid, billType, tarType string) (*custom.RespApplyTransactionBill, error)
	//ApplyFundBill //申请资金账单
	ApplyFundBill(billDate, accountType, tarType string) (*custom.RespApplyTransactionBill, error)
	//JSAPIOrders 直连商户JSAPI下单
	JSAPIOrders(data custom.ReqJSAPIOrders) (*custom.RespJSAPIOrders, error)
	//JSAPIOrdersForPartner 服务商JSAPI下单
	JSAPIOrdersForPartner(data custom.ReqJSAPIOrdersForPartner) (*custom.RespJSAPIOrders, error)

	// EduPaPayPresign 教培续费通预签约
	EduPaPayPresign(data custom.ReqEduPaPayPresign) (*custom.RespEduPaPayPresign, error)
	// EduPaPayContractQueryById 通过协议号查询教培续费通签约
	EduPaPayContractQueryById(contractId string, query url.Values) (*custom.RespEduPaPayContractQuery, error)
	// EduPaPayContractQueryByOpenId 通用户标识查询教培续费通签约
	EduPaPayContractQueryByOpenId(openid string, query url.Values) (*custom.RespEduPaPayContractQueryList, error)
	// DissolveEduPaPayContract 教培续费通解约
	DissolveEduPaPayContract(contractId string) error
	// SendEduPaPayNotifications 发送预扣款通知
	SendEduPaPayNotifications(contractId string, data custom.ReqSendEduPaPayNotifications) error
	// EduPaPayTransactions 教培通扣款受理
	EduPaPayTransactions(data custom.ReqEduPaPayTransactions) error
	// EduPaPayQueryOrderByTransactionId 教培通微信订单号查单
	EduPaPayQueryOrderByTransactionId(transactionId string, query url.Values) (*custom.RespEduPaPayQueryOrder, error)
	// EduPaPayQueryOrderByOutTradeNo 教培通商户订单号查单
	EduPaPayQueryOrderByOutTradeNo(outTradeNo string, query url.Values) (*custom.RespEduPaPayQueryOrder, error)

	// QueryOrganizationInfoById 获取机构信息(根据机构ID)
	QueryOrganizationInfoById(organizationId string) (*custom.RespOrganizationInfo, error)
	// QueryOrganizationInfoByName 获取机构信息(根据机构名称)
	QueryOrganizationInfoByName(organizationName string) (*custom.RespOrganizationInfo, error)
	// ObtainAuthToken 获取授权凭证
	ObtainAuthToken(data custom.ReqObtainAuthToken) (*custom.RespObtainAuthToken, error)
	// Deprecated: PayCredential 旧版扣款接口,已废弃
	PayCredential(data custom.ReqPayCredential) (*custom.RespPayCredential, error)
	// QueryFaceUserInfo 查询刷脸用户信息
	QueryFaceUserInfo(organizationId, outUserId string, isDecrypt ...bool) (*custom.RespQueryFaceUserInfo, error)
	// UpdateFaceUserInfo 修改刷脸用户信息
	UpdateFaceUserInfo(data custom.ReqUpdateUserInfo) error
	// DissolveFaceUserContract 解除刷脸用户签约关系
	DissolveFaceUserContract(organizationId, outUserId string) error
	// PreSignature 预签约
	PreSignature(data custom.ReqPresignToken) (*custom.RespPresignToken, error)
	// OfflineFaceTransactions 申请扣款
	OfflineFaceTransactions(data custom.ReqOfflinefaceTransactions) (*custom.RespOfflinefaceTransactions, error)
	// ContractQuery 签约查询
	ContractQuery(contractId, AppId string) (*custom.RespContractQuery, error)
	// FaceMessageDecryption 人脸报文(签约解约)消息解密
	FaceMessageDecryption(data custom.FaceMessageCiphertext) (*custom.FaceMessagePlaintext, error)
	// QueryRepurchaseUsersList 查询重采用户列表
	QueryRepurchaseUsersList(organizationId, offset, limit string) (*custom.RespQueryRepurchaseUsersList, error)
	// QueryRetake 查询重采
	QueryRetake(collectionId string) (*custom.FaceCollections, error)
	// QueryOfflineFaceOrders 离线人脸团餐专属查单
	QueryOfflineFaceOrders(outTradeNo, spMchid, subMchid, businessProductId string) (*custom.RespOfflinefaceTransactions, error)
	// GetAuthInfo 获取AuthInfo
	GetAuthInfo(data custom.ReqGetAuthInfo) (*custom.RespGetAuthInfo, error)
	// GetRepaymentUrl 获取还款链接
	GetRepaymentUrl(data custom.ReqGetRepaymentUrl) (*custom.RespGetRepaymentUrl, error)

	//SmartGuideRegister 服务人员注册
	SmartGuideRegister(data custom.ReqSmartGuideRegister) (*custom.RespSmartGuideRegister, error)
	//SmartGuideAssign 服务人员分配
	SmartGuideAssign(guideId string, data custom.ReqSmartGuideAssign) error
	//SmartGuideQuery 服务人员查询
	SmartGuideQuery(storeId, subMchid, userId, mobile, workId, limit, offset string) (*custom.RespSmartGuideQuery, error)
	//SmartGuideUpdate 服务人员信息更新
	SmartGuideUpdate(guideId string, data custom.ReqSmartGuideUpdate) error

	// EduSchoolPayPreSign 校园轻松付预签约
	EduSchoolPayPreSign(data custom.ReqEduSchoolPayPreSign) (*custom.RespEduSchoolPayPreSign, error)
	// EduSchoolPayContractQueryById 校园轻松付通过协议号查询签约
	EduSchoolPayContractQueryById(contractId string) (*custom.RespEduSchoolPayContractQuery, error)
	// DissolveEduSchoolPayContract 校园轻松付解约
	DissolveEduSchoolPayContract(contractId string) error
	// EduSchoolPayContractQueryByOpenId 校园轻松付查询用户签约列表
	EduSchoolPayContractQueryByOpenId(openId string, query url.Values) (*custom.RespEduSchoolPayContractQueryPage, error)
	// EduSchoolPayTransactions 校园轻松付扣款
	EduSchoolPayTransactions(data custom.ReqEduSchoolPayTransactions) (*custom.RespEduSchoolPayTransactions, error)
	// EduSchoolPayQueryOrderByTransactionId 校园轻松付微信支付订单号查单
	EduSchoolPayQueryOrderByTransactionId(transactionId string, query url.Values) (*custom.RespEduSchoolPayTransactions, error)
	// EduSchoolPayQueryOrderByOutTradeNo 校园轻松付商户订单号查单
	EduSchoolPayQueryOrderByOutTradeNo(outTradeNo string, query url.Values) (*custom.RespEduSchoolPayTransactions, error)

	//QueryViolationNotifications 查询商户违规通知回调地址
	QueryViolationNotifications() (*custom.GeneralViolationNotifications, error)
	//CreateViolationNotifications 创建商户违规通知回调地址
	CreateViolationNotifications(data custom.GeneralViolationNotifications) (*custom.GeneralViolationNotifications, error)
	//UpdateViolationNotifications 修改商户违规通知回调地址
	UpdateViolationNotifications(data custom.GeneralViolationNotifications) (*custom.GeneralViolationNotifications, error)
	//DeleteViolationNotifications 删除商户违规通知回调地址
	DeleteViolationNotifications() error
}

// PayClient PayClient
type PayClient struct {
	MchId               string            // 商户号
	ApiV3Key            string            // apiV3密钥
	ApiSerialNo         string            // API证书序列号
	ApiPrivateKey       *rsa.PrivateKey   // API私钥
	ApiCertificate      *x509.Certificate // API证书
	PlatformSerialNo    string            // 平台证书序列号
	PlatformCertificate *x509.Certificate // 平台证书
	HttpClient          *http.Client
}

func (c *PayClient) doRequest(requestData interface{}, url string, httpMethod string, isCheck ...bool) ([]byte, error) {
	var data []byte
	if requestData != nil {
		var err error
		data, err = json.Marshal(requestData)
		if err != nil {
			return nil, err
		}
	}
	authorization, err := c.Authorization(httpMethod, url, data)
	if err != nil {
		return nil, err
	}
	// 重试3次，避免因网络原因导致失败
	retryTimes := 3
	var resp *http.Response
	for i := 0; i < retryTimes; i++ {
		resp, err = SimpleRequest(c.HttpClient, url, httpMethod, authorization, data, c.PlatformSerialNo)
		if err != nil {
			continue
		}
		break
	}
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if len(isCheck) > 0 {
		if !isCheck[0] {
			err = c.VerifyResponse(resp.StatusCode, &resp.Header, body)
			if err != nil {
				return nil, err
			}
		}
	} else {
		err = c.VerifyResponse(resp.StatusCode, &resp.Header, body)
		if err != nil {
			return nil, err
		}
	}

	return body, nil
}

// Decrypt Decrypt解密
func (c *PayClient) Decrypt(algorithm string, cipherText string, associatedData string, nonce string) ([]byte, error) {
	// 默认使用AEAD_AES_256_GCM
	switch algorithm {
	default:
		fallthrough
	case constant.AlgorithmAEADAES256GCM:
		decodedCipherText, _ := base64.StdEncoding.DecodeString(cipherText)

		block, err := aes.NewCipher([]byte(c.ApiV3Key))
		if err != nil {
			return nil, err
		}

		aesGcm, err := cipher.NewGCM(block)
		if err != nil {
			return nil, err
		}

		plaintext, err := aesGcm.Open(nil, []byte(nonce), decodedCipherText, []byte(associatedData))
		if err != nil {
			return nil, err
		}
		return plaintext, nil
	}
}

// RsaEncryptByPrivateKey 使用商户私钥RSA加密
func (c *PayClient) RsaEncryptByPrivateKey(origData []byte) (string, error) {
	h := crypto.Hash.New(crypto.SHA256)
	h.Write(origData)
	hashed := h.Sum(nil)
	// 进行rsa加密签名
	signedData, err := rsa.SignPKCS1v15(rand.Reader, c.ApiPrivateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signedData), nil
}

// RsaDecryptByPrivateKey 使用商户私钥RSA解密
func (c *PayClient) RsaDecryptByPrivateKey(ciphertext string) (string, error) {
	cipherData, _ := base64.StdEncoding.DecodeString(ciphertext)
	rng := rand.Reader

	plaintext, err := rsa.DecryptOAEP(sha1.New(), rng, c.ApiPrivateKey, cipherData, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// RsaEncryptByPublicKey 使用平台公钥RSA加密
func (c *PayClient) RsaEncryptByPublicKey(plaintext string) (string, error) {
	if c.PlatformSerialNo == "" || c.PlatformCertificate == nil {
		return "", fmt.Errorf("请先初始化平台证书")
	}
	secretMessage := []byte(plaintext)
	rng := rand.Reader

	cipherData, err := rsa.EncryptOAEP(sha1.New(), rng, c.PlatformCertificate.PublicKey.(*rsa.PublicKey), secretMessage, nil)
	if err != nil {
		return "", err
	}

	ciphertext := base64.StdEncoding.EncodeToString(cipherData)
	return ciphertext, nil
}
