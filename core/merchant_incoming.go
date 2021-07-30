package core

import (
	"encoding/json"
	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/custom"
	"github.com/louismax/wxpayv3/utils"
	"net/http"
)

//QuerySettlementAccount QuerySettlementAccount
func (c *PayClient) QuerySettlementAccount(subMchid string) (*custom.SettlementAccount, error) {
	params := map[string]string{"sub_mchid": subMchid}
	body, err := c.doRequest(nil, utils.BuildUrl(params, nil, constant.ApiQuerySettlementAccount), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.SettlementAccount{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetStatusRepairOrderByBusinessCode GetStatusRepairOrderByBusinessCode
func (c *PayClient) GetStatusRepairOrderByBusinessCode(businessCode string) (*custom.RespGetStatusRepairOrder, error) {
	params := map[string]string{"business_code": businessCode}
	body, err := c.doRequest(nil, utils.BuildUrl(params, nil, constant.APIGetStatusRepairOrderByBusinessCode), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespGetStatusRepairOrder{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetStatusRepairOrderByApplymentId GetStatusRepairOrderByApplymentId
func (c *PayClient) GetStatusRepairOrderByApplymentId(applymentId string) (*custom.RespGetStatusRepairOrder, error) {
	params := map[string]string{"applyment_id": applymentId}
	body, err := c.doRequest(nil, utils.BuildUrl(params, nil, constant.APIGetStatusRepairOrderByApplymentId), http.MethodGet)
	if err != nil {
		return nil, err
	}

	resp := custom.RespGetStatusRepairOrder{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
