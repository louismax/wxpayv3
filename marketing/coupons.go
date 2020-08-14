package marketing

import (
	"encoding/json"
	"fmt"
	"strconv"
)

/**
LssueCoupons 发放代金券API
https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/marketing/convention/chapter3_2.shtml
*/
type LssueCoupons struct {
	Stock_id            string `json:"stock_id"`            //批次号
	Openid              string `json:"-"`                   //openid
	Out_request_no      string `json:"out_request_no"`      //商户单号
	Appid               string `json:"appid"`               //appid
	Stock_creator_mchid string `json:"stock_creator_mchid"` //批次创建商户号
	Coupon_value        uint64 `json:"coupon_value"`        //指定面额发券，面额
	Coupon_minimum      uint64 `json:"coupon_minimum"`      //指定面额发券，券门槛
}

// APIUrl LssueCoupons APIURL
func (this LssueCoupons) APIUrl() string {
	return "/v3/marketing/favor/users/" + this.Openid + "/coupons"
}

// Method LssueCoupons Method
func (this LssueCoupons) Method() string {
	return "POST"
}

// Params CreateOrganization Params
func (this LssueCoupons) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

func (this LssueCoupons) RawJsonStr() string {
	jsonBytes, _ := json.Marshal(this)
	var mapResult map[string]interface{}
	json.Unmarshal(jsonBytes, &mapResult)

	if InterfaceToint(mapResult["coupon_value"]) == 0 {
		delete(mapResult, "coupon_value")
	}

	if InterfaceToint(mapResult["coupon_minimum"]) == 0 {
		delete(mapResult, "coupon_minimum")
	}

	//jsons, errs := json.Marshal(this) //转换成JSON返回的是byte[]
	jsons, errs := json.Marshal(mapResult) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
		return "Error"
	}
	//fmt.Println(string(jsons))
	return string(jsons)
}

func InterfaceToint(inter interface{}) int {
	res := int(0)
	switch inter.(type) {
	case string:
		//fmt.Println("string", inter.(string))
		i, _ := strconv.Atoi(inter.(string))
		res = i
	case int:
		//fmt.Println("int", inter.(int))
		res = inter.(int)
	case float64:
		res = int(inter.(float64))
	}
	return res
}

type RespLssueCoupons struct {
	Coupon_id string `json:"coupon_id"` //代金券id
}
