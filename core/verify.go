package core

import (
	"bufio"
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// ErrResponseBody ErrResponseBody
type ErrResponseBody struct {
	HttpStatus int             `json:"http_status"`
	Code       string          `json:"code"`
	Message    string          `json:"message"`
	ReqId      string          `json:"req_id"`
	Detail     json.RawMessage `json:"detail"`
}

func (r *ErrResponseBody) Error() string {
	if r.Detail == nil {
		return fmt.Sprintf("HttpStatus:%v Code:%s Message:%s RequestId:%s", r.HttpStatus, r.Code, r.Message, r.ReqId)
	}
	return fmt.Sprintf("HttpStatus:%v Code:%s Message:%s RequestId:%s Detail:%s", r.HttpStatus, r.Code, r.Message, r.ReqId, r.Detail)
}

// VerifyResponse 验签
func (c *PayClient) VerifyResponse(httpStatus int, header *http.Header, body []byte) error {
	if httpStatus != http.StatusOK && httpStatus != http.StatusNoContent {
		if body == nil {
			return fmt.Errorf("验证响应失败")
		}
		var response ErrResponseBody
		err := json.Unmarshal(body, &response)
		if err != nil {
			return err
		}
		// 先Unmarshal再赋值，防止被覆盖为空值
		response.HttpStatus = httpStatus
		response.ReqId = c.getRequestId(header)
		return &response
	}
	headerSerial := c.getWechatPaySerial(header)       //获取应答签名证书序列号
	headerSignature := c.getWechatPaySignature(header) //获取应答签名值
	headerTimestamp := c.getWechatPayTimestamp(header) //获取应答时间戳
	headerNonce := c.getWechatPayNonce(header)         //获取应答随机串
	return c.verify(headerSerial, headerSignature, headerTimestamp, headerNonce, body)
}

func (c *PayClient) getRequestId(header *http.Header) string {
	return header.Get("Request-Id")
}

// getWechatPaySerial 获取headers中的Wechatpay-Serial
func (c *PayClient) getWechatPaySerial(header *http.Header) string {
	return header.Get("Wechatpay-Serial")
}

// getWechatPaySignature 获取headers中的Wechatpay-Signature
func (c *PayClient) getWechatPaySignature(header *http.Header) string {
	return header.Get("Wechatpay-Signature")
}

// getWechatPaySignature 获取headers中的Wechatpay-Timestamp
func (c *PayClient) getWechatPayTimestamp(header *http.Header) string {
	return header.Get("Wechatpay-Timestamp")
}

// getWechatPaySignature 获取headers中的Wechatpay-Nonce
func (c *PayClient) getWechatPayNonce(header *http.Header) string {
	return header.Get("Wechatpay-Nonce")
}

func (c *PayClient) verify(headerSerial string, headerSignature string, headerTimestamp string, headerNonce string, body []byte) error {
	if headerSerial == "" {
		//没有证书序列号,直接不用验签？！
		fmt.Printf("\033[33m%s\n", "[Warning]--WxPayV3:当前请求应答结果中不存在Wechatpay-Serial,请注意应答来源是否合法！！！")
		return nil
	}
	switch {
	case headerSignature == "":
		return fmt.Errorf("微信支付签名参数无效")
	case headerTimestamp == "":
		return fmt.Errorf("微信支付时间戳参数无效")
	case headerNonce == "":
		return fmt.Errorf("微信支付随机字符串参数无效")
	}

	//构造验签名串
	verificationStr, err := c.buildVerificationString(headerTimestamp, headerNonce, body)
	if err != nil {
		return err
	}
	// 应答header中的signature是base64加密的，所以要先解密
	decodedSignature, err := base64.StdEncoding.DecodeString(headerSignature)
	if err != nil {
		return err
	}

	//判断Serial是平台证书还是平台公钥
	if strings.Contains(headerSerial, "PUB_KEY_ID_") {
		if c.WechatPayPublicKeyID != "" && c.WechatPayPublicKey != nil {
			fmt.Printf("\033[35m%s\n", "[Info]--WxPayV3:使用微信支付平台公钥验签,应答中的公钥ID:"+headerSerial)
			if headerSerial == c.WechatPayPublicKeyID {
				err = c.verifySignatureByPubKey(string(decodedSignature), verificationStr)
				if err != nil {
					fmt.Printf("\033[31m%s\n", "[Error]--WxPayV3:签名校验失败！"+err.Error())
					return fmt.Errorf("签名校验失败！err:%s", err.Error())
				}
				fmt.Printf("\033[32m%s\n", "[Info]--WxPayV3:请求应答签名校验通过!")
			} else {
				fmt.Printf("\033[33m%s\n", "[Warning]--WxPayV3:当前实例配置的平台公钥ID与应答结果中的公钥ID不匹配,请确认请求来源是否合法或平台公钥是否更新！！！")
			}
		} else {
			fmt.Printf("\033[33m%s\n", "[Warning]--WxPayV3:当前实例未配置平台公钥,无法进行应答结果签名验证！请通过配置平台平台公钥进行验签,确保请求应答来源为微信支付服务端！")
		}
	} else {
		//通过平台证书
		if _, ok := c.PlatformCertMap[headerSerial]; ok {
			fmt.Printf("\033[35m%s\n", "[Info]--WxPayV3:使用微信支付平台证书验签,应答中的证书编号:"+headerSerial)
			err = c.verifySignature(headerSerial, string(decodedSignature), verificationStr)
			if err != nil {
				fmt.Printf("\033[31m%s\n", "[Error]--WxPayV3:签名校验失败！"+err.Error())
				return fmt.Errorf("签名校验失败！err:%s", err.Error())
			}
			fmt.Printf("\033[32m%s\n", "[Info]--WxPayV3:请求应答签名校验通过!")
		} else {
			fmt.Printf("\033[33m%s\n", "[Warning]--WxPayV3:当前实例未配置平台证书,无法进行应答结果签名验证！请通过配置平台证书进行验签,确保请求应答来源为微信支付服务端！")
		}
	}
	return nil
}

func (c *PayClient) buildVerificationString(timestamp string, nonce string, body []byte) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	bluff := bufio.NewWriter(buffer)
	_, _ = bluff.WriteString(timestamp)
	_ = bluff.WriteByte('\n')
	_, _ = bluff.WriteString(nonce)
	_ = bluff.WriteByte('\n')
	if len(body) != 0 {
		_, _ = bluff.Write(body)
	}
	_ = bluff.WriteByte('\n')
	err := bluff.Flush()
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (c *PayClient) verifySignature(headerSerial, signature string, verificationStr []byte) error {
	h := sha256.New()
	h.Write(verificationStr)
	return rsa.VerifyPKCS1v15(c.PlatformCertMap[headerSerial].PublicKey.(*rsa.PublicKey), crypto.SHA256, h.Sum(nil), []byte(signature))
}

func (c *PayClient) verifySignatureByPubKey(signature string, verificationStr []byte) error {
	h := sha256.New()
	h.Write(verificationStr)
	return rsa.VerifyPKCS1v15(c.WechatPayPublicKey, crypto.SHA256, h.Sum(nil), []byte(signature))
}
