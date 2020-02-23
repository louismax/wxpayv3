package wxpayv3

import (
	"encoding/json"
	"errors"
	"strings"
)

// CertifiCates 获取平台证书列表
func (this *Client) CertifiCates(param CertifiCates) (interface{}, error) {

	result := SystemOauthTokenRsp{}

	rqs, err := this.doRequest(param, &result)
	if err != nil {
		return result, err
	}
	if strings.Contains(rqs, "code") {
		errmsg := SysError{}

		err = json.Unmarshal([]byte(rqs), &errmsg)
		if err != nil {
			return nil, err
		}
		return errmsg, errors.New("获取平台证书失败")
	}
	return result, nil
}
