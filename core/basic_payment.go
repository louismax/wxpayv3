package core

import (
	"encoding/json"
	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/custom"
	"github.com/louismax/wxpayv3/utils"
	"net/http"
)

//PaymentRefund PaymentRefund
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
