package wxpayv3

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/louismax/wxpayv3/marketing"
	"strings"
)

//CreatePayCredential 申请扣款(待废弃)
func (this *Client) CreatePayCredential(param ReqCreatePayCredential) (interface{}, error) {
	result := RespCreatePayCredential{}
	if param.PayCredential == "" {
		return result, errors.New("支付凭证不能为空！")
	}
	if param.Mchid == "" {
		return result, errors.New("商户号不能为空！")
	}
	if param.SubMchid == "" {
		return result, errors.New("子商户号不能为空！")
	}
	if param.Amount < 1 {
		return result, errors.New("支付金额不能为空！")
	}
	if param.DeviceIp == "" {
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
	if param.OutTradeNo == "" {
		return result, errors.New("商户订单号无效")
	}
	if param.BusinessSceneId != Business_scene_type_Mess && param.BusinessSceneId != Business_scene_type_Supermarket && param.BusinessSceneId != Business_scene_type_Infirmary && param.BusinessSceneId != Business_scene_type_Dev {
		return result, errors.New("支付场景无效")
	}

	reqdata := CreatePayCredential{}
	reqdata.Pay_credential = param.PayCredential

	reqdata.Merchant_info.Mchid = param.Mchid
	reqdata.Merchant_info.Sub_mchid = param.SubMchid
	if param.Appid != "" {
		reqdata.Merchant_info.Appid = param.Appid
	}
	if param.SubAppid != "" {
		reqdata.Merchant_info.Sub_appid = param.SubAppid
	}
	reqdata.Trade_amount_info.Amount = param.Amount
	reqdata.Trade_amount_info.Currency = "CNY"
	reqdata.Scene_info.Device_ip = param.DeviceIp
	reqdata.Device_info.Mac = param.Mac
	if param.GoodsTag != "" {
		reqdata.Goods_tag = param.GoodsTag
	}
	reqdata.Description = param.Description
	reqdata.Attach = param.Attach
	reqdata.Out_trade_no = param.OutTradeNo
	reqdata.Business_info.Business_product_id = Business_scene_type_K12
	reqdata.Business_info.Business_scene_id = param.BusinessSceneId

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
			return errmsg, errors.New(fmt.Sprintf("当前订单已经关闭，请换单号继续扣款!response:%+v", rqs))
		} else if errmsg.Code == "RESOURCE_ALREADY_EXISTS" {
			return errmsg, errors.New(fmt.Sprintf("订单已经支付，请勿重复扣款!response:%+v", rqs))
		} else if errmsg.Code == "PARAM_ERROR" {
			return errmsg, errors.New(fmt.Sprintf("参数错误，请检查参数!response:%+v", rqs))
		} else if errmsg.Code == "SYSTEM_ERROR" {
			return errmsg, errors.New(fmt.Sprintf("系统错误，请稍后重试!response:%+v", rqs))
		} else {
			return errmsg, errors.New(fmt.Sprintf("系统异常!response:%+v", rqs))
		}
	}
	return result, nil
}

//SendCoupon 发放代金券
func (this *Client) SendCoupon(param marketing.LssueCoupons) (interface{}, error) {
	result := marketing.RespLssueCoupons{}
	if param.Appid == "" {
		return result, errors.New("AppID不能为空！")
	}
	if param.Openid == "" {
		return result, errors.New("OpenID不能为空！")
	}
	if param.Stock_id == "" {
		return result, errors.New("券批次号不能为空！")
	}
	if param.Stock_creator_mchid == "" {
		return result, errors.New("发券商户号不能为空！")
	}
	if param.Out_request_no == "" {
		return result, errors.New("商户发券单号不能为空！")
	}

	//fmt.Println(fmt.Sprintf("Request:%+v", param))

	rqs, err := this.doRequest(param, &result)
	if err != nil {
		return result, err
	}

	if strings.Contains(rqs, "code") {
		errmsg := SysError{}

		//fmt.Println(fmt.Sprintf("RESP:%+v", rqs))

		err = json.Unmarshal([]byte(rqs), &errmsg)
		if err != nil {
			return nil, err
		}

		return errmsg, errors.New("系统错误!")

	}
	return result, nil
}

//Transactions K12查单
func (this *Client) Transactions(param ReqTransactions) (interface{}, error) {
	if param.Out_Trade_No == "" {
		return nil, errors.New("商户单号不能为空！")
	}
	if param.Sp_Mchid == "" {
		return nil, errors.New("服务商号不能为空！")
	}
	if param.Sp_Mchid == "" {
		return nil, errors.New("子商户号不能为空！")
	}
	//if param.Business_Product_ID == "" {
	//	param.Business_Product_ID = "K12"
	//}
	param.Business_Product_ID = 2
	result := RespTransactionDeduction{}
	rqs, err := this.doRequest(param, &result)
	if err != nil {
		return nil, err
	}
	//fmt.Println(rqs)
	if strings.Contains(rqs, "code") {
		errmsg := SysError{}
		err = json.Unmarshal([]byte(rqs), &errmsg)
		if err != nil {
			return nil, err
		}
		//fmt.Println(fmt.Sprintf("%+v", errmsg))
		return errmsg, errors.New(rqs)
	}
	return result, nil
}

//TransactionDeduction 申请扣款(新版、支持分账)
func (this *Client) TransactionDeduction(param ReqTransactionDeduction) (interface{}, error) {
	result := RespTransactionDeduction{}
	if param.AuthCode == "" {
		return result, errors.New("支付凭证不能为空！")
	}
	if param.SpMchid == "" {
		return result, errors.New("商户号不能为空！")
	}
	if param.SubMchid == "" {
		return result, errors.New("子商户号不能为空！")
	}
	if param.SpAppid == "" {
		return result, errors.New("Appid不能为空！")
	}
	if param.Amount.Total < 1 {
		return result, errors.New("支付金额不能为空！")
	}
	if param.Amount.Currency == "" {
		return result, errors.New("货币类型不能为空！")
	}
	if param.SceneInfo.DeviceIp == "" {
		return result, errors.New("设备IP不能为空！")
	}
	if param.Description == "" {
		return result, errors.New("商品信息无效！")
	}
	if param.Attach == "" {
		return result, errors.New("商户备注信息无效")
	}
	if param.OutTradeNo == "" {
		return result, errors.New("商户订单号无效")
	}
	if param.Business.BusinessProductId != Business_scene_type_K12 && param.Business.BusinessProductId != Business_scene_type_QY {
		return result, errors.New("产品ID无效")
	}
	if param.Business.BusinessSceneId != Business_scene_type_Mess && param.Business.BusinessSceneId != Business_scene_type_Supermarket && param.Business.BusinessSceneId != Business_scene_type_Infirmary && param.Business.BusinessSceneId != Business_scene_type_Dev && param.Business.BusinessSceneId != Business_scene_type_K12Test && param.Business.BusinessSceneId != Business_scene_type_QYBus && param.Business.BusinessSceneId != Business_scene_type_TXBus {
		return result, errors.New("场景ID无效")
	}

	rqs, err := this.doRequest(param, &result)
	if err != nil {
		return result, err
	}

	fmt.Println(rqs)

	if strings.Contains(rqs, "code") {
		errmsg := SysError{}

		err = json.Unmarshal([]byte(rqs), &errmsg)
		if err != nil {
			return nil, err
		}
		if errmsg.Code == "ORDER_CLOSED" {
			return errmsg, errors.New(fmt.Sprintf("当前订单已经关闭,请换单号继续扣款,原始返回:%s", rqs))
		} else if errmsg.Code == "RESOURCE_ALREADY_EXISTS" {
			return errmsg, errors.New(fmt.Sprintf("订单已经支付,请勿重复扣款,原始返回:%s", rqs))
		} else if errmsg.Code == "PARAM_ERROR" {
			return errmsg, errors.New(fmt.Sprintf("参数错误,原始返回:%s", rqs))
		} else if errmsg.Code == "INVALID_REQUEST" {
			return errmsg, errors.New(fmt.Sprintf("凭证失效,原始返回:%s", rqs))
		} else if errmsg.Code == "SYSTEM_ERROR" {
			return errmsg, errors.New(fmt.Sprintf("系统错误,原始返回:%s", rqs))
		} else {
			return errmsg, errors.New(fmt.Sprintf("未知异常,原始返回:%s", rqs))
		}
	}
	return result, nil

}
