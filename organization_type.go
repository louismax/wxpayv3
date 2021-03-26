package wxpayv3

import (
	"encoding/json"
	"fmt"
)

type CreateOrganization struct {
	Organization_name string `json:"organization_name"`
}

// APIUrl CreateOrganization APIURL
func (this CreateOrganization) APIUrl() string {
	return "/v3/offlinefacemch/organizations"
}

// Method CreateOrganization Method
func (this CreateOrganization) Method() string {
	return "POST"
}

// Params CreateOrganization Params
func (this CreateOrganization) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

func (this CreateOrganization) RawJsonStr() string {
	jsons, errs := json.Marshal(this) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
		return "Error"
	}
	return string(jsons)
}

type RespCreateOrganization struct {
	Organization_id string `json:"organization_id"`
}

type QueryOrganization struct {
	Organization_id string `json:"organization_id"`
}

// APIUrl QueryOrganization APIURL
func (this QueryOrganization) APIUrl() string {
	return fmt.Sprintf("/v3/offlinefacemch/organizations?organization_id=%s", this.Organization_id)
}

// Method QueryOrganization Method
func (this QueryOrganization) Method() string {
	return "GET"
}

// Params QueryOrganization Params
func (this QueryOrganization) Params() map[string]string {
	var m = make(map[string]string)
	//m["Organization_id"] = this.Organization_id
	return m
}

func (this QueryOrganization) RawJsonStr() string {
	return ""
}

type RespQueryOrganization struct {
	Organization_id   string `json:"organization_id"`
	Organization_name string `json:"organization_name"`
}

type UpdateOrganization struct {
	Organization_id   string `json:"organization_id"`
	Organization_name string `json:"organization_name"`
}

// APIUrl UpdateOrganization APIURL
func (this UpdateOrganization) APIUrl() string {
	return "/v3/offlinefacemch/organizations/" + this.Organization_id
}

// Method UpdateOrganization Method
func (this UpdateOrganization) Method() string {
	return "PATCH"
}

// Params CreateOrganization Params
func (this UpdateOrganization) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

func (this UpdateOrganization) RawJsonStr() string {
	req := CreateOrganization{}
	req.Organization_name = this.Organization_name

	jsons, errs := json.Marshal(req) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
		return "Error"
	}
	return string(jsons)
}
