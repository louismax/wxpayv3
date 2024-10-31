package core

import (
	"encoding/json"
	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/custom"
	"github.com/louismax/wxpayv3/utils"
	"net/http"
)

//QueryViolationNotifications 查询商户违规通知回调地址
func (c *PayClient) QueryViolationNotifications() (*custom.GeneralViolationNotifications, error) {
	body, err := c.doRequest(nil, utils.BuildUrl(nil, nil, constant.APIViolationNotifications), http.MethodGet)
	if err != nil {
		return nil, err
	}

	resp := custom.GeneralViolationNotifications{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//CreateViolationNotifications 创建商户违规通知回调地址
func (c *PayClient) CreateViolationNotifications(data custom.GeneralViolationNotifications) (*custom.GeneralViolationNotifications, error) {
	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIViolationNotifications), http.MethodPost)
	if err != nil {
		return nil, err
	}

	resp := custom.GeneralViolationNotifications{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//UpdateViolationNotifications 修改商户违规通知回调地址
func (c *PayClient) UpdateViolationNotifications(data custom.GeneralViolationNotifications) (*custom.GeneralViolationNotifications, error) {
	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIViolationNotifications), http.MethodPut)
	if err != nil {
		return nil, err
	}

	resp := custom.GeneralViolationNotifications{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//DeleteViolationNotifications 删除商户违规通知回调地址
func (c *PayClient) DeleteViolationNotifications() error {
	_, err := c.doRequest(nil, utils.BuildUrl(nil, nil, constant.APIViolationNotifications), http.MethodDelete)
	if err != nil {
		return err
	}
	return nil
}
