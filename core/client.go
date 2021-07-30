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
)

type Client interface {
	// Authorization 获取签名Authorization，由认证类型和签名信息组成
	Authorization(httpMethod string, urlString string, body []byte) (string, error)
	// Certificate 获取平台证书
	Certificate() (*custom.CertificateResp, error)
	// SetClientPlatformCert 设置平台证书
	SetClientPlatformCert(certificateStr string) error
	// RsaEncryptByPrivateKey 使用商户私钥RAS加密敏感数据
	RsaEncryptByPrivateKey(origData []byte) (string, error)
	// RsaDecryptByPrivateKey 使用商户私钥RAS解密敏感数据
	RsaDecryptByPrivateKey(ciphertext string) (string, error)
	// RsaEncryptByPublicKey 使用平台公钥RAS加密敏感数据
	RsaEncryptByPublicKey(plaintext string) (string, error)

	// UploadImage 上传图片（获取MediaId）
	UploadImage(filePath string) (*custom.RespUploadImage, error)

	// QuerySettlementAccount 获取结算账户
	QuerySettlementAccount(subMchid string) (*custom.SettlementAccount, error)
	// GetStatusRepairOrderByBusinessCode 通过业务申请编号查询申请状态
	GetStatusRepairOrderByBusinessCode(businessCode string) (*custom.RespGetStatusRepairOrder, error)
	// GetStatusRepairOrderByApplymentId 通过申请单号查询申请状态
	GetStatusRepairOrderByApplymentId(applymentId string) (*custom.RespGetStatusRepairOrder, error)

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
	// OfflinefaceTransactions 申请扣款
	OfflinefaceTransactions(data custom.ReqOfflinefaceTransactions) (*custom.RespOfflinefaceTransactions, error)
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
}

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

func (c *PayClient) doRequest(requestData interface{}, url string, httpMethod string) ([]byte, error) {
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
	err = c.VerifyResponse(resp.StatusCode, &resp.Header, body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

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

//RsaEncryptByPublicKey 使用平台公钥RSA加密
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
