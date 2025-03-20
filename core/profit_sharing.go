package core

import (
	"encoding/json"
	"fmt"
	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/custom"
	"github.com/louismax/wxpayv3/utils"
	"net/http"
	"net/url"
)

// InitiateProfitSharing 请求分账(注意,默认会做敏感数据加密)
func (c *PayClient) InitiateProfitSharing(data custom.ReqInitiateProfitSharing) (*custom.RespInitiateProfitSharing, error) {
	var err error
	for i, v := range data.Receivers {
		if v.Name != "" {
			if c.WechatPayPublicKeyID != "" && c.WechatPayPublicKey != nil { //优先使用微信平台公钥加密敏感数据
				data.Receivers[i].Name, err = c.RsaEncryptByWxPayPubKey(v.Name)
				if err != nil {
					return nil, err
				}
			} else if c.DefaultPlatformSerialNo != "" && len(c.PlatformCertMap) > 0 { //使用微信平台证书加密敏感数据
				data.Receivers[i].Name, err = c.RsaEncryptByWxPayPubCertKey(v.Name)
				if err != nil {
					return nil, err
				}
			} else {
				return nil, fmt.Errorf("发起分账需要做敏感数据加密,当前实例的微信平台公钥或证书不允许为空")
			}
		}
	}

	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIInitiateProfitSharing), http.MethodPost)
	if err != nil {
		return nil, err
	}
	resp := custom.RespInitiateProfitSharing{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryProfitSharingResult 查询分账结果
func (c *PayClient) QueryProfitSharingResult(subMchid, transactionId, outOrderNo string) (*custom.RespQueryProfitSharingResult, error) {
	params := map[string]string{"out_order_no": outOrderNo, "sub_mchid": subMchid, "transaction_id": transactionId}
	body, err := c.doRequest(nil, utils.BuildUrl(params, nil, constant.APIQueryProfitSharingResult), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespQueryProfitSharingResult{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// InitiateProfitSharingReturnOrders is InitiateProfitSharingReturnOrders
func (c *PayClient) InitiateProfitSharingReturnOrders(data custom.ReqInitiateProfitSharingReturnOrders) (*custom.RespInitiateProfitSharingReturnOrders, error) {
	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIInitiateProfitSharingReturnOrders), http.MethodPost)
	if err != nil {
		return nil, err
	}
	resp := custom.RespInitiateProfitSharingReturnOrders{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryProfitSharingReturnOrders IS QueryProfitSharingReturnOrders
func (c *PayClient) QueryProfitSharingReturnOrders(subMchid, outReturnNo, outOrderNo string) (*custom.RespQueryProfitSharingReturnOrders, error) {
	params := map[string]string{"out_order_no": outOrderNo, "sub_mchid": subMchid, "out_return_no": outReturnNo}
	body, err := c.doRequest(nil, utils.BuildUrl(params, nil, constant.APIQueryProfitSharingReturnOrders), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespQueryProfitSharingReturnOrders{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnfreezeRemainingFunds  解冻剩余资金
func (c *PayClient) UnfreezeRemainingFunds(data custom.ReqUnfreezeRemainingFunds) (*custom.RespUnfreezeRemainingFunds, error) {
	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIUnfreezeRemainingFunds), http.MethodPost)
	if err != nil {
		return nil, err
	}
	resp := custom.RespUnfreezeRemainingFunds{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryRemainingFrozenAmount 查询剩余冻结金额
func (c *PayClient) QueryRemainingFrozenAmount(transactionId string) (*custom.RespQueryRemainingFrozenAmount, error) {
	params := map[string]string{"transaction_id": transactionId}
	body, err := c.doRequest(nil, utils.BuildUrl(params, nil, constant.APIQueryRemainingFrozenAmount), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespQueryRemainingFrozenAmount{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryMaximumSplitRatio 查询最大分账比例
func (c *PayClient) QueryMaximumSplitRatio(subMchid string) (*custom.RespQueryMaximumSplitRatio, error) {
	params := map[string]string{"sub_mchid": subMchid}
	body, err := c.doRequest(nil, utils.BuildUrl(params, nil, constant.APIQueryMaximumSplitRatio), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespQueryMaximumSplitRatio{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddProfitSharingReceiver  添加分账接收方(注意,默认会做敏感数据加密)
func (c *PayClient) AddProfitSharingReceiver(data custom.ReqAddProfitSharingReceiver) (*custom.RespAddProfitSharingReceiver, error) {
	if c.WechatPayPublicKeyID != "" && c.WechatPayPublicKey != nil { //优先使用微信平台公钥加密敏感数据
		if data.Name != "" {
			var err error
			data.Name, err = c.RsaEncryptByWxPayPubKey(data.Name)
			if err != nil {
				return nil, err
			}
		}
	} else if c.DefaultPlatformSerialNo != "" && len(c.PlatformCertMap) > 0 { //使用微信平台证书加密敏感数据
		if data.Name != "" {
			var err error
			data.Name, err = c.RsaEncryptByWxPayPubCertKey(data.Name)
			if err != nil {
				return nil, err
			}
		}
	} else {
		return nil, fmt.Errorf("添加分账接收方需要做敏感数据加密,当前实例的微信平台公钥或证书不允许为空")
	}

	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIAddProfitSharingReceiver), http.MethodPost)
	if err != nil {
		return nil, err
	}
	resp := custom.RespAddProfitSharingReceiver{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteProfitSharingReceiver  删除分账接受方
func (c *PayClient) DeleteProfitSharingReceiver(data custom.ReqDeleteProfitSharingReceiver) (*custom.RespDeleteProfitSharingReceiver, error) {
	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIDeleteProfitSharingReceiver), http.MethodPost)
	if err != nil {
		return nil, err
	}
	resp := custom.RespDeleteProfitSharingReceiver{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ApplyProfitSharingBill 申请分账账单
func (c *PayClient) ApplyProfitSharingBill(billDate, subMchid, tarType string) (*custom.RespApplyTransactionBill, error) {
	qy := url.Values{}
	qy.Set("bill_date", billDate)
	if subMchid != "" {
		qy.Set("sub_mchid", subMchid)
	}
	if tarType != "" {
		qy.Set("tar_type", tarType)
	}
	body, err := c.doRequest(nil, utils.BuildUrl(nil, qy, constant.APIApplyProfitSharingBill), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespApplyTransactionBill{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
