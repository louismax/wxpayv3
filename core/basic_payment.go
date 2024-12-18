package core

import (
	"encoding/json"
	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/custom"
	"github.com/louismax/wxpayv3/utils"
	"net/http"
	"net/url"
)

// PaymentQueryOrderByTransactionId 查询订单-通过微信订单号
func (c *PayClient) PaymentQueryOrderByTransactionId(transactionId, mchID string, subMchId ...string) (*custom.ReqPaymentQueryOrder, error) {
	params := map[string]string{"transaction_id": transactionId}
	qy := url.Values{}

	resp := custom.ReqPaymentQueryOrder{}
	if len(subMchId) > 0 {
		//服务商模式
		qy.Set("sp_mchid", mchID)
		qy.Set("sub_mchid", subMchId[0])
		body, err := c.doRequest(nil, utils.BuildUrl(params, qy, constant.APIPaymentPartnerQueryOrderByTransactionId), http.MethodGet)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(body, &resp); err != nil {
			return nil, err
		}
	} else {
		qy.Set("mchid", mchID)
		//直连商户模式
		body, err := c.doRequest(nil, utils.BuildUrl(params, qy, constant.APIPaymentQueryOrderByTransactionId), http.MethodGet)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(body, &resp); err != nil {
			return nil, err
		}
	}
	return &resp, nil
}

// PaymentQueryOrderByOutTradeNo 查询订单-通过商户订单号
func (c *PayClient) PaymentQueryOrderByOutTradeNo(outTradeNo, mchID string, subMchId ...string) (*custom.ReqPaymentQueryOrder, error) {
	params := map[string]string{"out-trade-no": outTradeNo}
	qy := url.Values{}

	resp := custom.ReqPaymentQueryOrder{}
	if len(subMchId) > 0 {
		//服务商模式
		qy.Set("sp_mchid", mchID)
		qy.Set("sub_mchid", subMchId[0])
		body, err := c.doRequest(nil, utils.BuildUrl(params, qy, constant.APIPaymentPartnerQueryOrderByOutTradeNo), http.MethodGet)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(body, &resp); err != nil {
			return nil, err
		}
	} else {
		qy.Set("mchid", mchID)
		//直连商户模式
		body, err := c.doRequest(nil, utils.BuildUrl(params, qy, constant.APIPaymentQueryOrderByOutTradeNo), http.MethodGet)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(body, &resp); err != nil {
			return nil, err
		}
	}
	return &resp, nil
}

// PaymentCloseOrder 关闭订单
func (c *PayClient) PaymentCloseOrder(outTradeNo, mchID string, subMchId ...string) error {
	params := map[string]string{"out_trade_no": outTradeNo}
	if len(subMchId) > 0 {
		//服务商模式
		reqData := map[string]interface{}{
			"sp_mchid":  mchID,
			"sub_mchid": subMchId[0],
		}
		_, err := c.doRequest(reqData, utils.BuildUrl(params, nil, constant.APIPaymentPartnerCloseOrder), http.MethodPost)
		if err != nil {
			return err
		}
	} else {
		//直连商户模式
		reqData := map[string]interface{}{
			"mchid": mchID,
		}
		_, err := c.doRequest(reqData, utils.BuildUrl(params, nil, constant.APIPaymentCloseOrder), http.MethodPost)
		if err != nil {
			return err
		}
	}
	return nil
}

// PaymentRefund 直连商户申请退款
func (c *PayClient) PaymentRefund(data custom.ReqPaymentRefund) (*custom.RespPaymentRefund, error) {
	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIPaymentRefund), http.MethodPost)
	if err != nil {
		return nil, err
	}
	resp := custom.RespPaymentRefund{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PaymentRefundForPartner 服务商申请退款
func (c *PayClient) PaymentRefundForPartner(data custom.ReqPaymentRefundForPartner) (*custom.RespPaymentRefund, error) {
	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIPaymentRefund), http.MethodPost)
	if err != nil {
		return nil, err
	}
	resp := custom.RespPaymentRefund{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ApplyTransactionBill 申请交易账单
func (c *PayClient) ApplyTransactionBill(billDate, subMchid, billType, tarType string) (*custom.RespApplyTransactionBill, error) {
	qy := url.Values{}
	qy.Set("bill_date", billDate)
	if subMchid != "" {
		qy.Set("sub_mchid", subMchid)
	}
	if billType != "" {
		qy.Set("bill_type", billType)
	}
	if tarType != "" {
		qy.Set("tar_type", tarType)
	}
	body, err := c.doRequest(nil, utils.BuildUrl(nil, qy, constant.APIApplyTransactionBill), http.MethodGet)
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

// ApplyFundBill 申请资金账单
func (c *PayClient) ApplyFundBill(billDate, accountType, tarType string) (*custom.RespApplyTransactionBill, error) {
	qy := url.Values{}
	qy.Set("bill_date", billDate)
	if accountType != "" {
		qy.Set("account_type", accountType)
	}
	if tarType != "" {
		qy.Set("tar_type", tarType)
	}
	body, err := c.doRequest(nil, utils.BuildUrl(nil, qy, constant.APIApplyFundBill), http.MethodGet)
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

// DownloadBill 下载账单
func (c *PayClient) DownloadBill(downloadUrl string) ([]byte, error) {
	body, err := c.doRequest(nil, utils.BuildUrl(nil, nil, downloadUrl), http.MethodGet, true)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// JSAPIOrders 直连商户JSAPI下单
func (c *PayClient) JSAPIOrders(data custom.ReqJSAPIOrders) (*custom.RespJSAPIOrders, error) {
	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIJSAPIOrders), http.MethodPost)
	if err != nil {
		return nil, err
	}
	resp := custom.RespJSAPIOrders{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// JSAPIOrdersForPartner 服务商JSAPI下单
func (c *PayClient) JSAPIOrdersForPartner(data custom.ReqJSAPIOrdersForPartner) (*custom.RespJSAPIOrders, error) {
	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIJSAPIOrdersForPartner), http.MethodPost)
	if err != nil {
		return nil, err
	}
	resp := custom.RespJSAPIOrders{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
