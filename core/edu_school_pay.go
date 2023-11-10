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

func (c *PayClient) EduSchoolPayPreSign(data custom.ReqEduSchoolPayPreSign) (*custom.RespEduSchoolPayPreSign, error) {
	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIEduSchoolPayPreSign), http.MethodPost)
	if err != nil {
		return nil, err
	}
	resp := custom.RespEduSchoolPayPreSign{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *PayClient) EduSchoolPayContractQueryById(contractId string) (*custom.RespEduSchoolPayContractQuery, error) {
	params := map[string]string{"contract_id": contractId}
	body, err := c.doRequest(nil, utils.BuildUrl(params, nil, constant.APIEduSchoolPayContractQueryById), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespEduSchoolPayContractQuery{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *PayClient) DissolveEduSchoolPayContract(contractId string) error {
	params := map[string]string{"contract_id": contractId}
	_, err := c.doRequest(nil, utils.BuildUrl(params, nil, constant.APIDissolveEduSchoolPayContract), http.MethodPost)
	if err != nil {
		return err
	}
	return nil
}

func (c *PayClient) EduSchoolPayContractQueryByOpenId(openId string, query url.Values) (*custom.RespEduSchoolPayContractQueryPage, error) {
	params := map[string]string{"openid": openId}

	if query.Get("plan_id") == "" {
		return nil, fmt.Errorf("参数不合法,query中plan_id为必填参数")
	}
	if query.Get("contract_status") == "" {
		return nil, fmt.Errorf("参数不合法,query中contract_status为必填参数")
	}

	body, err := c.doRequest(nil, utils.BuildUrl(params, query, constant.APIEduSchoolPayContractQueryByOpenId), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespEduSchoolPayContractQueryPage{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *PayClient) EduSchoolPayTransactions(data custom.ReqEduSchoolPayTransactions) (*custom.RespEduSchoolPayTransactions, error) {
	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIEduSchoolPayTransactions), http.MethodPost)
	if err != nil {
		return nil, err
	}
	resp := custom.RespEduSchoolPayTransactions{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *PayClient) EduSchoolPayQueryOrderByTransactionId(transactionId string, query url.Values) (*custom.RespEduSchoolPayTransactions, error) {
	params := map[string]string{"transaction_id": transactionId}
	if query.Get("sub_mchid") == "" {
		return nil, fmt.Errorf("参数不合法,query中sub_mchid为必填参数")
	}

	body, err := c.doRequest(nil, utils.BuildUrl(params, query, constant.APIEduSchoolPayQueryOrderByTransactionId), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespEduSchoolPayTransactions{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *PayClient) EduSchoolPayQueryOrderByOutTradeNo(outTradeNo string, query url.Values) (*custom.RespEduSchoolPayTransactions, error) {
	params := map[string]string{"out_trade_no": outTradeNo}
	if query.Get("sub_mchid") == "" {
		return nil, fmt.Errorf("参数不合法,query中sub_mchid为必填参数")
	}
	body, err := c.doRequest(nil, utils.BuildUrl(params, query, constant.APIEduSchoolPayQueryOrderByOutTradeNo), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespEduSchoolPayTransactions{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
