package wxpayv3

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// QueryUserInfo 查询刷脸用户信息
func (this *Client) QueryUserInfo(param QueryUserInfo) (RespQueryUserInfo, error) {
	result := RespQueryUserInfo{}
	if param.Organization_id == "" {
		return result, errors.New("机构ID不能为空！")
	}
	if param.Out_user_id == "" {
		return result, errors.New("商户用户ID不能为空！")
	}

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
			return result, errors.New("查询用户信息失败!")
		} else if errmsg.Code == "SIGN_ERROR" {
			return result, errors.New("系统错误，签名失败!")
		} else {
			return result, nil
		}
	}
	username, err := this.RsaDecrypt(result.User_name)
	if err != nil {
		return result, err
	}
	result.User_name = username

	return result, nil
}

// UpdateUserInfo 修改刷脸用户信息
func (this *Client) UpdateUserInfo(param UpdateUserInfo) error {
	if param.Organization_id == "" {
		return errors.New("机构ID不能为空！")
	}
	if param.Out_user_id == "" {
		return errors.New("用户ID不能为空！")
	}
	if param.RequestData.User_name == "" {
		return errors.New("用户姓名不能为空！")
	}
	if param.RequestData.User_type != Usertype_STUDENT && param.RequestData.User_type != Usertype_STAFF {
		return errors.New("用户类型无效！")
	}
	if param.RequestData.User_type == Usertype_STUDENT {
		if param.RequestData.Student_info.Class_name == "" {
			return errors.New("班级名称无效")
		}
	}
	if param.RequestData.User_type == Usertype_STAFF {
		if param.RequestData.Staff_info.Occupation == "" {
			return errors.New("职业无效")
		}
	}
	if param.RequestData.Status != Status_NORMAL && param.RequestData.Status != Status_DISABLED {
		return errors.New("用户状态无效！")
	}
	if param.RequestData.Phone == "" {
		return errors.New("用户手机号不能为空！")
	}
	if this.PFSerialno == "" {
		return errors.New("请先初始化平台证书")
	}

	var err error
	param.RequestData.User_name, err = this.RsaOAEPEncrypt(param.RequestData.User_name)
	if err != nil {
		return errors.New("用户姓名加密错误！")
	}

	param.RequestData.Phone, err = this.RsaOAEPEncrypt(param.RequestData.Phone)
	if err != nil {
		return errors.New("用户手机号加密错误！")
	}

	rqs, err := this.doRequest(param, nil)
	fmt.Println(rqs)
	if err != nil {
		return errors.New("修改用户信息失败")
	} else {
		return nil
	}
}

// QueryContracts 查询签约信息
func (this *Client) QueryContracts(param QueryContracts) (RespQueryContracts, error) {
	result := RespQueryContracts{}
	if param.Contract_id == "" {
		return result, errors.New("签约ID不能为空！")
	}
	if param.Appid == "" {
		return result, errors.New("APPID不能为空！")
	}

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
			return result, errors.New("查询签约信息失败!")
		} else if errmsg.Code == "SIGN_ERROR" {
			return result, errors.New("系统错误，签名失败!")
		} else {
			return result, nil
		}
	}
	return result, nil
}
