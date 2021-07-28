package core

import (
	"encoding/json"
	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/custom"
	"github.com/louismax/wxpayv3/utils"
	"net/http"
)

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
