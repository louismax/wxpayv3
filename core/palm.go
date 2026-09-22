package core

import (
	"encoding/json"
	"net/http"

	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/custom"
	"github.com/louismax/wxpayv3/utils"
)

func (c *PayClient) PalmServicePreBind(serviceId, ticketId string, data map[string]interface{}) (*custom.RespPalmServicePreBind, error) {
	body, err := c.doRequest(data, utils.BuildUrl(map[string]string{
		"service_id":         serviceId,
		"merchant_ticket_id": ticketId,
	}, nil, constant.APIPalmServicePreBind), http.MethodPost)
	if err != nil {
		return nil, err
	}
	resp := custom.RespPalmServicePreBind{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *PayClient) PalmServiceQueryBind(serviceId, ticketId string) (*custom.RespPalmServiceQueryBind, error) {
	body, err := c.doRequest(nil, utils.BuildUrl(map[string]string{
		"service_id":         serviceId,
		"merchant_ticket_id": ticketId,
	}, nil, constant.APIPalmServiceQueryBind), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespPalmServiceQueryBind{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *PayClient) PalmServiceDeleteBind(serviceId, ticketId string) error {
	_, err := c.doRequest(nil, utils.BuildUrl(map[string]string{
		"service_id":         serviceId,
		"merchant_ticket_id": ticketId,
	}, nil, constant.APIPalmServiceDeleteBind), http.MethodDelete)
	if err != nil {
		return err
	}
	return nil
}
