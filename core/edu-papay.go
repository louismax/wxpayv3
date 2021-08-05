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

//EduPaPayPresign EduPaPayPresign
func (c *PayClient) EduPaPayPresign(data custom.ReqEduPaPayPresign) (*custom.RespEduPaPayPresign, error) {
	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIEduPaPayPresign), http.MethodPost)
	if err != nil {
		return nil, err
	}
	resp := custom.RespEduPaPayPresign{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//EduPaPayContractQueryById EduPaPayContractQueryById
func (c *PayClient) EduPaPayContractQueryById(contractId string, query url.Values) (*custom.RespEduPaPayContractQuery, error) {
	params := map[string]string{"contract_id": contractId}
	if query.Get("appid") == "" {
		return nil, fmt.Errorf("参数不合法,query中appid为必填参数")
	}
	body, err := c.doRequest(nil, utils.BuildUrl(params, query, constant.APIEduPaPayContractQueryById), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespEduPaPayContractQuery{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//EduPaPayContractQueryByOpenId EduPaPayContractQueryByOpenId
func (c *PayClient) EduPaPayContractQueryByOpenId(openid string, query url.Values) (*custom.RespEduPaPayContractQueryList, error) {
	params := map[string]string{"openid": openid}
	if query.Get("appid") == "" {
		return nil, fmt.Errorf("参数不合法,query中appid为必填参数")
	}
	if query.Get("plan_id") == "" {
		return nil, fmt.Errorf("参数不合法,query中plan_id为必填参数")
	}
	body, err := c.doRequest(nil, utils.BuildUrl(params, query, constant.APIEduPaPayContractQueryByOpenId), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespEduPaPayContractQueryList{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *PayClient) DissolveEduPaPayContract(contractId string) error {
	params := map[string]string{"contract_id": contractId}
	_, err := c.doRequest(nil, utils.BuildUrl(params, nil, constant.APIDissolveEduPaPayContract), http.MethodDelete)
	if err != nil {
		return err
	}
	return nil
}

func (c *PayClient) SendEduPaPayNotifications(contractId string, data custom.ReqSendEduPaPayNotifications) error {
	params := map[string]string{"contract_id": contractId}
	_, err := c.doRequest(data, utils.BuildUrl(params, nil, constant.APISendEduPaPayNotifications), http.MethodPost)
	if err != nil {
		return err
	}
	return nil
}

func (c *PayClient) EduPaPayTransactions(data custom.ReqEduPaPayTransactions) error {
	_, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIEduPaPayTransactions), http.MethodPost)
	if err != nil {
		return err
	}
	return nil
}

func (c *PayClient) EduPaPayQueryOrderByTransactionId(transactionId string, query url.Values) (*custom.RespEduPaPayQueryOrder, error) {
	params := map[string]string{"transaction_id": transactionId}
	body, err := c.doRequest(nil, utils.BuildUrl(params, query, constant.APIEduPaPayQueryOrderByTransactionId), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespEduPaPayQueryOrder{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *PayClient) EduPaPayQueryOrderByOutTradeNo(outTradeNo string, query url.Values) (*custom.RespEduPaPayQueryOrder, error) {
	params := map[string]string{"out_trade_no": outTradeNo}
	body, err := c.doRequest(nil, utils.BuildUrl(params, query, constant.APIEduPaPayQueryOrderByOutTradeNo), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespEduPaPayQueryOrder{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
