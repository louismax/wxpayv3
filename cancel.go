package wxpayv3

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// CancelRequest 主动解约
func (this *Client) CancelRequest(param CancelRequest) (bool, error) {
	if param.Organization_id == "" {
		return false, errors.New("机构ID不能为空！")
	}
	if param.User_id == "" {
		return false, errors.New("刷脸用户ID不能为空！")
	}

	rqs, err := this.doRequest(param, nil)
	if err != nil {
		return false, err
	}
	if strings.Contains(rqs, "code") {
		errmsg := SysError{}
		err = json.Unmarshal([]byte(rqs), &errmsg)
		if err != nil {
			return false, err
		}
		return false, errors.New(fmt.Sprintf("申请代扣解约失败,ERR:%+v", errmsg))
	}
	fmt.Println(rqs)
	return true, nil
}
