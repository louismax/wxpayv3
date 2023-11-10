package core

import (
	"encoding/json"
	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/custom"
	"github.com/louismax/wxpayv3/utils"
	"net/http"
)

func (c *PayClient) EduSchoolPayPreSign(data custom.ReqEduSchoolPayPreSign) (*custom.RespEduPaPayPresign, error) {
	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APIEduSchoolPayPreSign), http.MethodPost)
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
