package wxpayv3

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

type OldAccount struct {
	appID    string
	subappid string
	mchID    string
	submchid string
	apiKey   string
	//isSandbox bool //沙箱环境
}

// [旧版]创建微信支付账号
func NewOldAccount(appID, mchID, submchid, subappid, apiKey string) *OldAccount {
	return &OldAccount{
		appID:    appID,
		mchID:    mchID,
		submchid: submchid,
		subappid: subappid,
		apiKey:   apiKey,
		//isSandbox: isSandbox,
	}
}

const (
	bodyType   = "application/xml; charset=utf-8"
	MD5        = "MD5"
	HMACSHA256 = "HMAC-SHA256"
	Fail       = "FAIL"
	Success    = "SUCCESS"
	Sign       = "sign"
)

type OldClient struct {
	account              *OldAccount // 商户请求信息
	signType             string      // 签名类型
	httpConnectTimeoutMs int         // 连接超时时间
	httpReadTimeoutMs    int         // 读取超时时间
}

func NewOldClient(account *OldAccount) *OldClient {
	return &OldClient{
		account:              account,
		signType:             MD5,
		httpConnectTimeoutMs: 2000,
		httpReadTimeoutMs:    1000,
	}
}

// 获取SDK调用凭证
func (c *OldClient) GetAuthInfo(params Params) (Params, error) {
	url := "https://payapp.weixin.qq.com/face/get_wxpayface_authinfo"

	params.SetString("version", "1").
		SetString("store_id", "K12").
		SetString("store_name", "K12").
		SetString("now", fmt.Sprintf("%d", time.Now().Unix()))

	xmlStr, err := c.postWithoutCert(url, params)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Err:%+v,body:%+v",err,xmlStr))
	}
	return c.processResponseXml(xmlStr)
}

//向 params 中添加 appid、mch_id、nonce_str、sign_type、sign
func (c *OldClient) fillRequestData(params Params) Params {
	params["appid"] = c.account.appID
	if c.account.subappid != "" {
		params["sub_appid"] = c.account.subappid
	}
	params["mch_id"] = c.account.mchID
	if c.account.submchid != "" {
		params["sub_mch_id"] = c.account.submchid
	}
	params["nonce_str"] = GetGUID()
	params["sign_type"] = c.signType
	params["sign"] = c.Sign(params)
	return params
}

// 签名方法
func (c *OldClient) Sign(params Params) string {
	// 创建切片
	var keys = make([]string, 0, len(params))
	// 遍历签名参数
	for k := range params {
		if k != "sign" { // 排除sign字段
			keys = append(keys, k)
		}
	}
	// 由于切片的元素顺序是不固定，所以这里强制给切片元素加个顺序
	sort.Strings(keys)

	//创建字符缓冲
	var buf bytes.Buffer
	for _, k := range keys {
		if len(params.GetString(k)) > 0 {
			buf.WriteString(k)
			buf.WriteString(`=`)
			buf.WriteString(params.GetString(k))
			buf.WriteString(`&`)
		}
	}
	// 加入apiKey作加密密钥
	buf.WriteString(`key=`)
	buf.WriteString(c.account.apiKey)

	var (
		dataMd5    [16]byte
		dataSha256 []byte
		str        string
	)

	switch c.signType {
	case MD5:
		dataMd5 = md5.Sum(buf.Bytes())
		str = hex.EncodeToString(dataMd5[:]) //需转换成切片
	case HMACSHA256:
		h := hmac.New(sha256.New, []byte(c.account.apiKey))
		h.Write(buf.Bytes())
		dataSha256 = h.Sum(nil)
		str = hex.EncodeToString(dataSha256[:])
	}
	return strings.ToUpper(str)
}

// https无证书请求
func (c *OldClient) postWithoutCert(url string, params Params) (string, error) {
	h := &http.Client{}
	p := c.fillRequestData(params)
	fmt.Println(fmt.Sprintf("WXApi参数组装：%+v", p))
	response, err := h.Post(url, bodyType, strings.NewReader(MapToXml(p)))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// 处理 HTTPS API返回数据，转换成Map对象。return_code为SUCCESS时，验证签名。
func (c *OldClient) processResponseXml(xmlStr string) (Params, error) {
	//qlong.DBG("下单返回信息String：%s", xmlStr)
	xmlStr = strings.Replace(xmlStr, "\n\t", "", -1)
	xmlStr = strings.Replace(xmlStr, " ", "", -1)
	var returnCode string
	params := XmlToMap(xmlStr)
	//fmt.Println(fmt.Sprintf("API返回信息：%+v", params))
	if params.ContainsKey("return_code") {
		returnCode = params.GetString("return_code")
	} else {
		return nil, errors.New(fmt.Sprintf("no return_code in XML,body:%+v",xmlStr))
	}
	if returnCode == Fail {
		return params, nil
	} else if returnCode == Success {
		if c.ValidSign(params) {
			return params, nil
		} else {
			return nil, errors.New(fmt.Sprintf("invalid sign value in XML,body:%+v",xmlStr))
		}
	} else {
		return nil, errors.New(fmt.Sprintf("return_code value is invalid in XML,body:%+v",xmlStr))
	}
}

// 验证签名
func (c *OldClient) ValidSign(params Params) bool {
	if !params.ContainsKey(Sign) {
		return false
	}
	return params.GetString(Sign) == c.Sign(params)
}
