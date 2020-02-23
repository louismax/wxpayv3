package wxpayv3

import (
	"encoding/json"
	"fmt"
)

type GetTokens struct {
	Scene         string        `json:"scene"`
	Web_init_data Web_Init_Data `json:"web_init_data"`
}

type Web_Init_Data struct {
	Out_user_id     string `json:"out_user_id"`     //商户侧用户ID
	Organization_id string `json:"organization_id"` //机构ID
}

// APIUrl GetTokens APIURL
func (this GetTokens) APIUrl() string {
	return "/v3/offlinefacemch/tokens"
}

// Method GetTokens Method
func (this GetTokens) Method() string {
	return "POST"
}

// Params GetTokens Params
func (this GetTokens) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

func (this GetTokens) RawJsonStr() string {
	jsons, errs := json.Marshal(this) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
		return "Error"
	}
	return string(jsons)
}

type RespGetTokens struct {
	Token string `json:"token"`
}
