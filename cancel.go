package wxpayv3

import (
	"errors"
	"fmt"
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
	fmt.Println(rqs)

	return true, nil
}
