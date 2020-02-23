package wxpayv3

import (
	"encoding/json"
	"errors"
	"strings"
)

//CreatePayCredential 申请扣款
func (this *Client) CreatePayCredential(param ReqCreatePayCredential) (interface{}, error) {
	result := RespCreatePayCredential{}
	if param.Pay_credential == "" {
		return result, errors.New("支付凭证不能为空！")
	}
	if param.Mchid == "" {
		return result, errors.New("商户号不能为空！")
	}
	if param.Sub_mchid == "" {
		return result, errors.New("子商户号不能为空！")
	}
	if param.Amount < 1 {
		return result, errors.New("支付金额不能为空！")
	}
	if param.Device_ip == "" {
		return result, errors.New("设备IP不能为空！")
	}
	if param.Mac == "" {
		return result, errors.New("设备MAC地址不能为空！")
	}
	if param.Description == "" {
		return result, errors.New("商品信息无效！")
	}
	if param.Attach == "" {
		return result, errors.New("商户备注信息无效")
	}
	if param.Out_trade_no == "" {
		return result, errors.New("商户订单号无效")
	}
	if param.Business_scene_id != Business_scene_type_Mess && param.Business_scene_id != Business_scene_type_Supermarket && param.Business_scene_id != Business_scene_type_Infirmary && param.Business_scene_id != Business_scene_type_Dev {
		return result, errors.New("支付场景无效")
	}

	reqdata := CreatePayCredential{}
	reqdata.Pay_credential = param.Pay_credential

	reqdata.Merchant_info.Mchid = param.Mchid
	reqdata.Merchant_info.Sub_mchid = param.Sub_mchid
	if param.Appid != "" {
		reqdata.Merchant_info.Appid = param.Appid
	}
	if param.Sub_appid != "" {
		reqdata.Merchant_info.Sub_appid = param.Sub_appid
	}
	reqdata.Trade_amount_info.Amount = param.Amount
	reqdata.Trade_amount_info.Currency = "CNY"
	reqdata.Scene_info.Device_ip = param.Device_ip
	reqdata.Device_info.Mac = param.Mac
	reqdata.Goods_tag = ""
	reqdata.Description = param.Description
	reqdata.Attach = param.Attach
	reqdata.Out_trade_no = param.Out_trade_no
	reqdata.Business_info.Business_product_id = Business_scene_type_K12
	reqdata.Business_info.Business_scene_id = param.Business_scene_id

	//b, _ := json.Marshal(reqdata)
	//fmt.Println(string(b))

	rqs, err := this.doRequest(reqdata, &result)
	if err != nil {
		return result, err
	}

	if strings.Contains(rqs, "code") {
		errmsg := SysError{}

		err = json.Unmarshal([]byte(rqs), &errmsg)
		if err != nil {
			return nil, err
		}
		//errb, _ := json.Marshal(errmsg)
		//fmt.Println(string(errb))
		if errmsg.Code == "ORDER_CLOSED" {
			return errmsg, errors.New("发起扣款失败，请换单号重试!")
		} else if errmsg.Code == "ORDER_NOT_EXIST" {
			return errmsg, errors.New("扣款超过限额，订单已失效!")
		} else {
			return errmsg, errors.New("系统错误!")
		}
	}
	return result, nil
}
