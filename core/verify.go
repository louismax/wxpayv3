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
)

//ErrResponseBody ErrResponseBody
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

// VerifyResponse VerifyResponse
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
		response.ReqId = header.Get("Request-Id")
		return &response
	}
	headerSerial := c.getWechatPaySerial(header)
	headerSignature := c.getWechatPaySignature(header)
	headerTimestamp := c.getWechatPayTimestamp(header)
	headerNonce := c.getWechatPayNonce(header)
	return c.verify(headerSerial, headerSignature, headerTimestamp, headerNonce, body)
}

func (c *PayClient) getWechatPaySerial(header *http.Header) string {
	return header.Get("Wechatpay-Serial")
}

func (c *PayClient) getWechatPaySignature(header *http.Header) string {
	return header.Get("Wechatpay-Signature")
}

func (c *PayClient) getWechatPayTimestamp(header *http.Header) string {
	return header.Get("Wechatpay-Timestamp")
}

func (c *PayClient) getWechatPayNonce(header *http.Header) string {
	return header.Get("Wechatpay-Nonce")
}

func (c *PayClient) verify(headerSerial string, headerSignature string, headerTimestamp string, headerNonce string, body []byte) error {
	if c.PlatformSerialNo != "" && headerSerial != c.PlatformSerialNo {
		return fmt.Errorf("平台证书序列号不匹配,headerSerial:%s,Client_PlatformSerialNo:%s", headerSerial, c.PlatformSerialNo)
	}
	switch {
	case headerSignature == "":
		return fmt.Errorf("微信支付签名参数无效")
	case headerTimestamp == "":
		return fmt.Errorf("微信支付时间戳参数无效")
	case headerNonce == "":
		return fmt.Errorf("微信支付随机字符串参数无效")
	}
	verificationStr, err := c.buildVerificationString(headerTimestamp, headerNonce, body)
	if err != nil {
		return err
	}
	// 应答header中的signature是base64加密的，所以要先解密
	decodedSignature, err := base64.StdEncoding.DecodeString(headerSignature)
	if err != nil {
		return err
	}
	if c.PlatformCertificate != nil {
		err = c.verifySignature(string(decodedSignature), verificationStr)
		if err != nil {
			return err
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

func (c *PayClient) verifySignature(signature string, verificationStr []byte) error {
	h := sha256.New()
	h.Write(verificationStr)
	return rsa.VerifyPKCS1v15(c.PlatformCertificate.PublicKey.(*rsa.PublicKey), crypto.SHA256, h.Sum(nil), []byte(signature))
}
