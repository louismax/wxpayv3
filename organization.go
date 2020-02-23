package wxpayv3

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// CreateOrganization 创建机构
func (this *Client) CreateOrganization(param CreateOrganization) (RespCreateOrganization, error) {
	result := RespCreateOrganization{}
	if param.Organization_name == "" {
		return result, errors.New("机构名称不能为空！")
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
			return result, errors.New("新建机构失败，请检查是否重名!")
		} else if errmsg.Code == "SIGN_ERROR" {
			return result, errors.New("系统错误，签名失败!")
		} else {
			return result, nil
		}
	}
	return result, nil
}

// QueryOrganization 查询机构
func (this *Client) QueryOrganization(param QueryOrganization) (RespQueryOrganization, error) {
	result := RespQueryOrganization{}
	if param.Organization_id == "" {
		return result, errors.New("机构ID不能为空！")
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
			return result, errors.New("新建机构失败，请检查是否重名!")
		} else if errmsg.Code == "SIGN_ERROR" {
			return result, errors.New("系统错误，签名失败!")
		} else {
			return result, nil
		}
	}

	return result, nil
}

// UpdateOrganization 修改机构
func (this *Client) UpdateOrganization(param UpdateOrganization) error {
	//result := RespQueryOrganization{}
	if param.Organization_id == "" {
		return errors.New("机构ID不能为空！")
	}
	if param.Organization_name == "" {
		return errors.New("机构名称不能为空！")
	}
	rqs, err := this.doRequest(param, nil)
	fmt.Println(rqs)
	if err != nil {
		return errors.New("修改机构信息失败")
	} else {
		return nil
	}
}
