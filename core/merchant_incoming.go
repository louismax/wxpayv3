package core

import (
	"encoding/json"
	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/custom"
	"github.com/louismax/wxpayv3/utils"
	"net/http"
)


//IncomingSubmitApplication 提交进件申请单
func (c *PayClient) IncomingSubmitApplication(data custom.ReqIncomingSubmitApplication) (*custom.RespIncomingSubmitApplication, error) {
	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIIncomingSubmitApplication), http.MethodPost)
	if err != nil {
		return nil, err
	}

	resp := custom.RespIncomingSubmitApplication{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//ModifySettlement 修改结算账号
func (c *PayClient) ModifySettlement(subMchid string,data custom.ReqModifySettlement) error {
	params := map[string]string{"sub_mchid": subMchid}
	_, err := c.doRequest(data, utils.BuildUrl(params, nil, constant.APIModifySettlement), http.MethodPost)
	if err != nil {
		return err
	}
	return nil
}

//QuerySettlementAccount 查询结算账户
func (c *PayClient) QuerySettlementAccount(subMchid string) (*custom.SettlementAccount, error) {
	params := map[string]string{"sub_mchid": subMchid}
	body, err := c.doRequest(nil, utils.BuildUrl(params, nil, constant.APIQuerySettlementAccount), http.MethodGet)
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

//GetStatusRepairOrderByBusinessCode 通过业务申请编号查询申请状态
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

//GetStatusRepairOrderByApplymentId 通过申请单号查询申请状态
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
