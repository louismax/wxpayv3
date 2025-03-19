package core

import (
	"encoding/json"
	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/custom"
	"github.com/louismax/wxpayv3/utils"
	"net/http"
	"net/url"
	"strconv"
)

func (c *PayClient) QueryComplaintsList(beginDate, endDate string, limit, offset int, mchId ...string) (*custom.RespComplaintsList, error) {
	query := url.Values{}
	query.Set("begin_date", beginDate)
	query.Set("end_date", endDate)
	if limit > 0 {
		query.Set("limit", strconv.Itoa(limit))
	}
	if offset > 0 {
		query.Set("offset", strconv.Itoa(offset))
	}
	if len(mchId) > 0 {
		query.Set("complainted_mchid", mchId[0])
	}
	body, err := c.doRequest(nil, utils.BuildUrl(nil, query, constant.APIComplaintsList), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespComplaintsList{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
