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

// QueryComplaintDetail 查询投诉单详情。
func (c *PayClient) QueryComplaintDetail(complaintId string) (*custom.RespComplaintDetail, error) {
	body, err := c.doRequest(nil, utils.BuildUrl(map[string]string{
		"complaint_id": url.PathEscape(complaintId),
	}, nil, constant.APIComplaintDetail), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespComplaintDetail{}
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryComplaintNegotiationHistory 查询协商历史，limit最大为300。
func (c *PayClient) QueryComplaintNegotiationHistory(complaintId string, limit, offset int) (*custom.RespComplaintNegotiationHistory, error) {
	query := url.Values{}
	if limit > 0 {
		query.Set("limit", strconv.Itoa(limit))
	}
	if offset > 0 {
		query.Set("offset", strconv.Itoa(offset))
	}
	body, err := c.doRequest(nil, utils.BuildUrl(map[string]string{
		"complaint_id": url.PathEscape(complaintId),
	}, query, constant.APIComplaintNegotiationHistory), http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp := custom.RespComplaintNegotiationHistory{}
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ResponseComplaint 回复用户，服务商使用complainted_mchid指定被诉子商户。
func (c *PayClient) ResponseComplaint(complaintId string, data custom.ReqResponseComplaint) error {
	_, err := c.doRequest(data, utils.BuildUrl(map[string]string{
		"complaint_id": url.PathEscape(complaintId),
	}, nil, constant.APIComplaintResponse), http.MethodPost, 1)
	return err
}

// ResponseComplaintImmediateService 回复即时服务投诉，并返回可关联协商历史的log_id。
func (c *PayClient) ResponseComplaintImmediateService(complaintId string, data custom.ReqResponseComplaintImmediateService) (*custom.RespResponseComplaintImmediateService, error) {
	body, err := c.doRequest(data, utils.BuildUrl(map[string]string{
		"complaint_id": url.PathEscape(complaintId),
	}, nil, constant.APIComplaintImmediateResponse), http.MethodPost, 1)
	if err != nil {
		return nil, err
	}
	resp := custom.RespResponseComplaintImmediateService{}
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CompleteComplaint 反馈处理完成，不执行退款。
func (c *PayClient) CompleteComplaint(complaintId string, data custom.ReqCompleteComplaint) error {
	_, err := c.doRequest(data, utils.BuildUrl(map[string]string{
		"complaint_id": url.PathEscape(complaintId),
	}, nil, constant.APIComplaintComplete), http.MethodPost, 1)
	return err
}
