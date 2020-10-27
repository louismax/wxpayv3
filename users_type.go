package wxpayv3

import (
	"encoding/json"
	"fmt"
)

const (
	Usertype_STUDENT = "STUDENT" //学生
	Usertype_STAFF   = "STAFF"   //教职工
)

const (
	Status_NORMAL   = "NORMAL"   //正常状态
	Status_DISABLED = "DISABLED" //禁用，后续刷脸用户将无法消费
)

type QueryUserInfo struct {
	Organization_id string `json:"organization_id"`
	Out_user_id     string `json:"out_user_id"`
}

// APIUrl QueryUserInfo APIURL
func (this QueryUserInfo) APIUrl() string {
	return fmt.Sprintf("/v3/offlinefacemch/organizations/%s/users/out-user-id/%s", this.Organization_id, this.Out_user_id)
}

// Method QueryUserInfo Method
func (this QueryUserInfo) Method() string {
	return "GET"
}

// Params QueryUserInfo Params
func (this QueryUserInfo) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// RawJsonStr QueryUserInfo RawJsonStr
func (this QueryUserInfo) RawJsonStr() string {
	return ""
}

type RespQueryUserInfo struct {
	User_id         string       `json:"user_id"`         //微信侧刷脸用户唯一ID
	Out_user_id     string       `json:"out_user_id"`     //商户刷脸用户ID
	Organization_id string       `json:"organization_id"` //机构ID
	User_name       string       `json:"user_name"`       //姓名
	User_type       string       `json:"user_type"`       //用户类型 学生：STUDENT教职工：STAFF
	Student_info    Student_Info `json:"student_info"`    //学生信息
	Staff_info      Staff_info   `json:"staff_info"`      //教职工信息
	Status          string       `json:"status"`          //用户状态 NOMAL：正常状态 DISABLED：禁用状态，此时支付被限制
	Contract_state  string       `json:"contract_state"`  //签约状态
	Face_image_ok   bool         `josn:"face_image_ok"`   //人脸图片上传状态
	Contract_id     string       `json:"contract_id"`     //签约ID
}

type Student_Info struct {
	Class_name string `json:"class_name"`
}
type Staff_info struct {
	Occupation string `json:"occupation"`
}

type UpdateUserInfo struct {
	Organization_id string            `json:"organization_id"`
	Out_user_id     string            `json:"out_user_id"`
	RequestData     UpdateRequestData `json:"requestdata"`
}

type UpdateRequestData struct {
	User_name    string       `json:"user_name"`
	User_type    string       `json:"user_type"`
	Student_info Student_Info `json:"student_info"`
	Staff_info   Staff_info   `json:"staff_info"`
	Status       string       `json:"status"`
	Phone        string       `json:"phone"`
}

// APIUrl UpdateUserInfo APIURL
func (this UpdateUserInfo) APIUrl() string {
	return fmt.Sprintf("/v3/offlinefacemch/organizations/%s/users/out-user-id/%s", this.Organization_id, this.Out_user_id)
}

// Method UpdateUserInfo Method
func (this UpdateUserInfo) Method() string {
	return "PATCH"
}

// Params UpdateUserInfo Params
func (this UpdateUserInfo) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// RawJsonStr UpdateUserInfo RawJsonStr
func (this UpdateUserInfo) RawJsonStr() string {
	jsons, errs := json.Marshal(this.RequestData) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
		return "Error"
	}
	return string(jsons)
}

type QueryContracts struct {
	Contract_id string `json:"contract_id"`
	Appid       string `json:"appid"`
}

// APIUrl QueryUserInfo APIURL
func (this QueryContracts) APIUrl() string {
	return fmt.Sprintf("/v3/offlineface/contracts/%s?appid=%s", this.Contract_id, this.Appid)
}

// Method QueryUserInfo Method
func (this QueryContracts) Method() string {
	return "GET"
}

// Params QueryUserInfo Params
func (this QueryContracts) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// RawJsonStr QueryUserInfo RawJsonStr
func (this QueryContracts) RawJsonStr() string {
	return ""
}

type RespQueryContracts struct {
	Contract_id              string `json:"contract_id"`              //签约ID
	Mchid                    string `json:"mchid"`                    //商户号
	Organization_id          string `json:"organization_id"`          //机构ID
	User_id                  string `json:"user_id"`                  //用户ID
	Openid                   string `json:"openid"`                   //签约用户openid
	Contract_state           string `json:"contract_state"`           //签约状态
	Contract_signed_time     string `json:"contract_signed_time"`     //签约时间
	Contract_terminated_time string `json:"contract_terminated_time"` //解约时间
	Contract_mode            string `json:"contract_mode"`            //签约模式 LIMIT_BANK_CARD：指定卡签约；PRIORITY_BANK_CARD：优先卡签约；LIMIT_NONE：任意卡签约
	Contract_bank_card_from  string `josn:"contract_bank_card_from"`  //签约卡来源 MERCHANT_LIMITED_BANK_CARD：商户指定的签约卡；USER_SELECT_FREE：用户选择的签约卡

}
