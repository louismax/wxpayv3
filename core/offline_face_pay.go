package core

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/custom"
	"github.com/louismax/wxpayv3/utils"
	"net/http"
	"net/url"
)

func (c *PayClient) QueryOrganizationInfoById(organizationId string) (*custom.RespOrganizationInfo, error) {
	params := map[string]string{"organization_id": organizationId}
	body, err := c.doRequest(nil, utils.BuildUrl(params, nil, constant.APIQueryOrganizationInfoById), http.MethodGet)
	if err != nil {
		return nil, err
	}

	resp := custom.RespOrganizationInfo{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *PayClient) QueryOrganizationInfoByName(organizationName string) (*custom.RespOrganizationInfo, error) {
	params := map[string]string{"organization_name": url.QueryEscape(organizationName)}
	body, err := c.doRequest(nil, utils.BuildUrl(params, nil, constant.APIQueryOrganizationInfoByName), http.MethodGet)
	if err != nil {
		return nil, err
	}

	resp := custom.RespOrganizationInfo{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *PayClient) ObtainAuthToken(data custom.ReqObtainAuthToken) (*custom.RespObtainAuthToken, error) {
	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIObtainAuthToken), http.MethodPost)
	if err != nil {
		return nil, err
	}

	resp := custom.RespObtainAuthToken{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: PayCredential 旧版扣款接口,已废弃
func (c *PayClient) PayCredential(data custom.ReqPayCredential) (*custom.RespPayCredential, error) {
	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIPayCredential), http.MethodPost)
	if err != nil {
		return nil, err
	}

	resp := custom.RespPayCredential{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *PayClient) QueryFaceUserInfo(organizationId, outUserId string, isDecrypt ...bool) (*custom.RespQueryFaceUserInfo, error) {
	params := map[string]string{"organization_id": organizationId, "out_user_id": outUserId}
	body, err := c.doRequest(nil, utils.BuildUrl(params, nil, constant.APIQueryFaceUserInfo), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespQueryFaceUserInfo{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	if len(isDecrypt) > 0 {
		if isDecrypt[0] {
			resp.UserName, err = c.RsaDecryptByPrivateKey(resp.UserName)
			if err != nil {
				return nil, err
			}
		}
	}
	return &resp, nil
}

func (c *PayClient) UpdateFaceUserInfo(data custom.ReqUpdateUserInfo) error {
	params := map[string]string{"organization_id": data.OrganizationId, "out_user_id": data.OutUserId}
	reqData := custom.ReqUpdateRequestData{
		UserType: data.RequestData.UserType,
		Status:   data.RequestData.Status,
	}
	if data.RequestData.StudentInfo != nil {
		reqData.StudentInfo = data.RequestData.StudentInfo
	}
	if data.RequestData.StaffInfo != nil {
		reqData.StaffInfo = data.RequestData.StaffInfo
	}
	var err error
	reqData.UserName, err = c.RsaEncryptByPublicKey(data.RequestData.UserName)
	if err != nil {
		return err
	}
	reqData.Phone, err = c.RsaEncryptByPublicKey(data.RequestData.Phone)
	if err != nil {
		return err
	}
	_, err = c.doRequest(reqData, utils.BuildUrl(params, nil, constant.APIUpdateFaceUserInfo), http.MethodPatch)
	if err != nil {
		return err
	}
	return nil
}

func (c *PayClient) DissolveFaceUserContract(organizationId, outUserId string) error {
	params := map[string]string{"organization_id": organizationId, "user_id": outUserId}
	_, err := c.doRequest(nil, utils.BuildUrl(params, nil, constant.APIDissolveFaceUserContract), http.MethodPost)
	if err != nil {
		return err
	}
	return nil
}

func (c *PayClient) PreSignature(data custom.ReqPresignToken) (*custom.RespPresignToken, error) {
	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIPreSignature), http.MethodPost)
	if err != nil {
		return nil, err
	}
	resp := custom.RespPresignToken{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *PayClient) OfflinefaceTransactions(data custom.ReqOfflinefaceTransactions) (*custom.RespOfflinefaceTransactions, error) {
	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIOfflinefaceTransactions), http.MethodPost)
	if err != nil {
		return nil, err
	}
	resp := custom.RespOfflinefaceTransactions{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *PayClient) ContractQuery(contractId, AppId string) (*custom.RespContractQuery, error) {
	params := map[string]string{"contract_id": contractId, "appid": AppId}
	body, err := c.doRequest(nil, utils.BuildUrl(params, nil, constant.APIContractQuery), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespContractQuery{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *PayClient) FaceMessageDecryption(data custom.FaceMessageCiphertext) (*custom.FaceMessagePlaintext, error) {
	// 对编码密文进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(data.Resource.Ciphertext)
	if err != nil {
		return nil, err
	}
	cx, err := aes.NewCipher([]byte(c.ApiV3Key))
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(cx)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(decodeBytes) < nonceSize {
		return nil, fmt.Errorf("密文证书长度不够")
	}
	res := custom.FaceMessagePlaintext{}
	if data.Resource.AssociatedData != "" {
		plaintext, err := gcm.Open(nil, []byte(data.Resource.Nonce), decodeBytes, []byte(data.Resource.AssociatedData))
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(plaintext, &res)
		if err != nil {
			return nil, err
		}
		return &res, nil
	} else {
		plaintext, err := gcm.Open(nil, []byte(data.Resource.Nonce), decodeBytes, nil)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(plaintext, &res)
		if err != nil {
			return nil, err
		}
		return &res, nil
	}
}

func (c *PayClient) QueryRepurchaseUsersList(organizationId, offset, limit string) (*custom.RespQueryRepurchaseUsersList, error) {
	params := map[string]string{"organization_id": organizationId, "offset": offset, "limit": limit}
	body, err := c.doRequest(nil, utils.BuildUrl(params, nil, constant.APIQueryRepurchaseUsersList), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespQueryRepurchaseUsersList{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *PayClient) QueryRetake(collectionId string) (*custom.FaceCollections, error) {
	params := map[string]string{"collection_id": collectionId}
	body, err := c.doRequest(nil, utils.BuildUrl(params, nil, constant.APIQueryRetake), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.FaceCollections{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *PayClient) QueryOfflineFaceOrders(outTradeNo, spMchid, subMchid, businessProductId string) (*custom.RespOfflinefaceTransactions, error) {
	params := map[string]string{"out_trade_no": outTradeNo, "sp_mchid": spMchid, "sub_mchid": subMchid, "business_product_id": businessProductId}
	body, err := c.doRequest(nil, utils.BuildUrl(params, nil, constant.APIQueryOfflineFaceOrders), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespOfflinefaceTransactions{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *PayClient) GetAuthInfo(data custom.ReqGetAuthInfo) (*custom.RespGetAuthInfo, error) {
	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIGetAuthInfo), http.MethodPost)
	if err != nil {
		return nil, err
	}
	resp := custom.RespGetAuthInfo{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *PayClient) GetRepaymentUrl(data custom.ReqGetRepaymentUrl) (*custom.RespGetRepaymentUrl, error) {
	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIGetRepaymentUrl), http.MethodPost)
	if err != nil {
		return nil, err
	}
	resp := custom.RespGetRepaymentUrl{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
