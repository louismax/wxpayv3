package core

import (
	"encoding/json"
	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/custom"
	"github.com/louismax/wxpayv3/utils"
	"net/http"
	"net/url"
)

func (c *PayClient) SmartGuideRegister(data custom.ReqSmartGuideRegister) (*custom.RespSmartGuideRegister, error) {
	body, err := c.doRequest(data, utils.BuildUrl(nil, nil, constant.APISmartGuideRegister), http.MethodPost)
	if err != nil {
		return nil, err
	}
	resp := custom.RespSmartGuideRegister{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *PayClient) SmartGuideAssign(guideId string, data custom.ReqSmartGuideAssign) error {
	params := map[string]string{"guide_id": guideId}
	_, err := c.doRequest(data, utils.BuildUrl(params, nil, constant.APISmartGuideAssign), http.MethodPost)
	if err != nil {
		return err
	}
	return nil
}

func (c *PayClient) SmartGuideQuery(storeId, subMchid, userId, mobile, workId, limit, offset string) (*custom.RespSmartGuideQuery, error) {
	qy := url.Values{}
	qy.Set("store_id", storeId)
	if subMchid != "" {
		qy.Set("sub_mchid", subMchid)
	}
	if userId != "" {
		qy.Set("userid", userId)
	}
	if mobile != "" {
		qy.Set("mobile", mobile)
	}
	if workId != "" {
		qy.Set("work_id", workId)
	}
	if limit != "" {
		qy.Set("limit", limit)
	}
	if offset != "" {
		qy.Set("offset", offset)
	}

	body, err := c.doRequest(nil, utils.BuildUrl(nil, qy, constant.APISmartGuideQuery), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespSmartGuideQuery{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *PayClient) SmartGuideUpdate(guideId string, data custom.ReqSmartGuideUpdate) error {
	params := map[string]string{"guide_id": guideId}
	_, err := c.doRequest(data, utils.BuildUrl(params, nil, constant.APISmartGuideUpdate), http.MethodPatch)
	if err != nil {
		return err
	}
	return nil
}
