package wxpayv3

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
)

//CancelRequest 主动解约申请
type CancelRequest struct {
	Organization_id string `json:"organization_id"`
	User_id         string `json:"user_id"`
}

// APIUrl CancelRequest APIURL
func (this CancelRequest) APIUrl() string {
	return fmt.Sprintf("/v3/offlinefacemch/organizations/%s/users/user-id/%s/terminate-contract", this.Organization_id, this.User_id)
}

// Method CancelRequest Method
func (this CancelRequest) Method() string {
	return "POST"
}

// Params CancelRequest Params
func (this CancelRequest) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// RawJsonStr CancelRequest RawJsonStr
func (this CancelRequest) RawJsonStr() string {
	return ""
}

//CancelCiphertext 解约报文数据
type CancelCiphertext struct {
	Id            string `json:"id"`            //通知唯一ID
	Create_time   string `json:"create_time"`   //通知创建的时间
	Event_type    string `json:"event_type"`    //通知的类型
	Resource_type string `json:"resource_type"` //通知的资源数据类型
	Resource      struct {
		Algorithm       string `json:"algorithm"`       //加密算法类型
		Ciphertext      string `json:"ciphertext"`      //数据密文
		Original_type   string `json:"original_type"`   //原始回调类型
		Associated_data string `json:"associated_data"` //附加数据
		Nonce           string `json:"nonce"`           //随机串
	} `json:"resource"` //通知资源数据
	Summary string `json:"summary"` //回调摘要
}

// CancelPlaintext 解约解密明文
type CancelPlaintext struct {
	User_id            string `json:"user_id"`            //微信刷脸用户唯一标识
	Out_user_id        string `json:"out_user_id"`        //商户刷脸用户唯一标识
	Organization_id    string `json:"organization_id"`    //机构编号
	Mch_id             string `json:"mch_id"`             //微信支付分配的商户号
	Notify_create_time string `json:"notify_create_time"` //通知创建时间
	Appid              string `json:"appid"`              //微信APPID
	Openid             string `json:"openid"`             //微信openid
}

// 解约报文解密
func (this CancelCiphertext) CancelDecryption(apiv3key string) (interface{}, error) {
	// 对编码密文进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(this.Resource.Ciphertext)
	if err != nil {
		return nil, err
	}

	c, err := aes.NewCipher([]byte(apiv3key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(decodeBytes) < nonceSize {
		return nil, errors.New("密文证书长度不够")
	}
	res := CancelPlaintext{}
	//res.Serial_no = cpinfo.Serial_no
	if this.Resource.Associated_data != "" {
		plaintext, err := gcm.Open(nil, []byte(this.Resource.Nonce), decodeBytes, []byte(this.Resource.Associated_data))
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(plaintext, &res)
		if err != nil {
			return nil, err
		}
		return res, nil
	} else {
		plaintext, err := gcm.Open(nil, []byte(this.Resource.Nonce), decodeBytes, nil)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(plaintext, &res)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
