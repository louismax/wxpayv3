package wxpayv3

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// GetTokens 获取小程序授权凭证
func (this *Client) GetTokens(userid, wxogid string) (RespGetTokens, error) {
	result := RespGetTokens{}
	if userid == "" {
		return result, errors.New("用户id不能为空！")
	}
	if wxogid == "" {
		return result, errors.New("机构ID不能为空！")
	}

	param := GetTokens{}
	param.Scene = "WEBSESSION" //小程序默认
	param.Web_init_data.Out_user_id = userid
	param.Web_init_data.Organization_id = wxogid

	rqs, err := this.doRequest(param, &result)
	if err != nil {
		return result, err
	}

	if strings.Contains(rqs, "code") {
		errmsg := SysError{}

		err = json.Unmarshal([]byte(rqs), &errmsg)
		if err != nil {
			return result, err
		}
		fmt.Println(fmt.Sprintf("%+v", errmsg))
		if errmsg.Code == "SYSTEM_ERROR" {
			return result, errors.New("获取授权凭证失败!")
		} else if errmsg.Code == "SIGN_ERROR" {
			return result, errors.New("系统错误，签名失败!")
		} else {
			return result, nil
		}
	}
	return result, nil
}
