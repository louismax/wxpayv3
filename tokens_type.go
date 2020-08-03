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

type Presign_Token struct {
	Business_name   string        `json:"business_name"`   //业务类型
	Facepay_user    Facepay_user  `json:"facepay_user"`    //刷脸用户信息
	Limit_bank_card LimitBankCard `json:"limit_bank_card"` //签约银行卡信息
}

type Facepay_user struct {
	Out_user_id         string         `json:"out_user_id"`
	Identification_name string         `json:"identification_name"`
	Organization_id     string         `json:"organization_id"`
	Identification      Identification `json:"identification"`
	Phone               string         `json:"phone"`
}

type Identification struct {
	Identification_type   string `json:"identification_type"`
	Identification_number string `json:"identification_number"`
}

type LimitBankCard struct {
	Bank_card_number    string         `json:"bank_card_number"`
	Identification_name string         `json:"identification_name"`
	Identification      Identification `json:"identification"`
	Valid_thru          string         `json:"valid_thru"`
}

// APIUrl GetTokens APIURL
func (this Presign_Token) APIUrl() string {
	return "/v3/offlineface/contracts/presign"
}

// Method GetTokens Method
func (this Presign_Token) Method() string {
	return "POST"
}

// Params GetTokens Params
func (this Presign_Token) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

func (this Presign_Token) RawJsonStr() string {
	jsons, errs := json.Marshal(this) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
		return "Error"
	}
	return string(jsons)
}

type RespPresign_Token struct {
	Presign_token string `json:"presign_token"`
}
