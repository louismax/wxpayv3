package utils

import (
	"bufio"
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	craned "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/custom"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	// NonceSymbols 随机字符串可用字符集
	NonceSymbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// NonceLength 随机字符串的长度
	NonceLength = 32
)

//GenerateNonce 生成32位随机字符串
func GenerateNonce() (string, error) {
	bs := make([]byte, NonceLength)
	_, err := rand.Read(bs)
	if err != nil {
		return "", err
	}
	symbolsByteLength := byte(len(NonceSymbols))
	for i, b := range bs {
		bs[i] = NonceSymbols[b%symbolsByteLength]
	}
	return string(bs), nil
}

//BuildMessage BuildMessage
func BuildMessage(httpMethod string, urlString string, body []byte, nonceStr string, timestamp int64) ([]byte, error) {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	urlPart := parsedUrl.Path
	if len(parsedUrl.RawQuery) != 0 {
		urlPart = urlPart + "?" + parsedUrl.RawQuery
	}

	buffer := bytes.NewBuffer([]byte{})
	buff := bufio.NewWriter(buffer)

	_, _ = buff.WriteString(httpMethod)
	_ = buff.WriteByte('\n')
	_, _ = buff.WriteString(urlPart)
	_ = buff.WriteByte('\n')
	_, _ = buff.WriteString(strconv.FormatInt(timestamp, 10))
	_ = buff.WriteByte('\n')
	_, _ = buff.WriteString(nonceStr)
	_ = buff.WriteByte('\n')
	if httpMethod == http.MethodPost || httpMethod == http.MethodPut || httpMethod == http.MethodPatch {
		_, _ = buff.Write(body)
	}
	_ = buff.WriteByte('\n')
	_ = buff.Flush()
	return buffer.Bytes(), nil
}

//Sign Sign
func Sign(message []byte, privateKey *rsa.PrivateKey) (string, error) {
	h := sha256.New()
	h.Write(message)
	signature, err := rsa.SignPKCS1v15(craned.Reader, privateKey, crypto.SHA256, h.Sum(nil))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

//BuildUrl BuildUrl
func BuildUrl(params map[string]string, query url.Values, subRoutes ...string) string {
	urlX := constant.ApiDomain
	for _, route := range subRoutes {
		if strings.Contains(route, constant.ApiDomain){
			urlX = route
			break
		}else{
			urlX += strings.TrimLeft(route, "/")
		}
	}
	for key, param := range params {
		urlX = strings.ReplaceAll(urlX, "{"+key+"}", param)
	}
	if query != nil {
		urlX += "?"
		urlX += query.Encode()
	}
	return urlX
}

//FaceMessageDecryption 无需初始化客户端的进行离线团餐人脸报文解密
func FaceMessageDecryption(data custom.FaceMessageCiphertext, apiV3Key string) (*custom.FaceMessagePlaintext, error) {
	// 对编码密文进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(data.Resource.Ciphertext)
	if err != nil {
		return nil, err
	}
	cx, err := aes.NewCipher([]byte(apiV3Key))
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(cx)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(decodeBytes) < nonceSize {
		return nil, fmt.Errorf("密文证书长度不够")
	}
	res := custom.FaceMessagePlaintext{}
	if data.Resource.AssociatedData != "" {
		plaintext, err := gcm.Open(nil, []byte(data.Resource.Nonce), decodeBytes, []byte(data.Resource.AssociatedData))
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(plaintext, &res)
		if err != nil {
			return nil, err
		}
		return &res, nil
	}
	plaintext, err := gcm.Open(nil, []byte(data.Resource.Nonce), decodeBytes, nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(plaintext, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
