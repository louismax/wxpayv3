package core

import (
	"encoding/json"
	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/custom"
	"github.com/louismax/wxpayv3/utils"
	"net/http"
)

//InitiateProfitSharing is InitiateProfitSharing
func (c *PayClient) InitiateProfitSharing(data custom.ReqInitiateProfitSharing) (*custom.RespInitiateProfitSharing, error) {
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

//QueryProfitSharingResult IS QueryProfitSharingResult
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

//InitiateProfitSharingReturnOrders is InitiateProfitSharingReturnOrders
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

//QueryProfitSharingReturnOrders IS QueryProfitSharingReturnOrders
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

//UnfreezeRemainingFunds  is UnfreezeRemainingFunds
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

//QueryRemainingFrozenAmount IS QueryRemainingFrozenAmount
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

//QueryMaximumSplitRatio IS QueryMaximumSplitRatio
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

//AddProfitSharingReceiver  is AddProfitSharingReceiver
func (c *PayClient) AddProfitSharingReceiver(data custom.ReqAddProfitSharingReceiver) (*custom.RespAddProfitSharingReceiver, error) {
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

//DeleteProfitSharingReceiver  is DeleteProfitSharingReceiver
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
