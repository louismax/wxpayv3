package core

import (
	"encoding/json"
	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/custom"
	"github.com/louismax/wxpayv3/utils"
	"net/http"
	"net/url"
)

//PaymentRefund 申请退款
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

//ApplyTransactionBill 申请交易账单
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

//ApplyFundBill 申请资金账单
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

//DownloadBill 下载账单
func (c *PayClient) DownloadBill(downloadUrl string) ([]byte, error) {
	body, err := c.doRequest(nil, utils.BuildUrl(nil, nil, downloadUrl), http.MethodGet, true)
	if err != nil {
		return nil, err
	}
	return body, nil
}
