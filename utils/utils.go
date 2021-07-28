package utils

import (
	"bufio"
	"bytes"
	"crypto"
	"crypto/rand"
	craned "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"github.com/louismax/wxpayv3/constant"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	NonceSymbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // 随机字符串可用字符集
	NonceLength  = 32                                                               // 随机字符串的长度
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
	if httpMethod == http.MethodPost || httpMethod == http.MethodPut {
		_, _ = buff.Write(body)
	}
	_ = buff.WriteByte('\n')
	_ = buff.Flush()
	return buffer.Bytes(), nil
}

func Sign(message []byte, privateKey *rsa.PrivateKey) (string, error) {
	h := sha256.New()
	h.Write(message)
	signature, err := rsa.SignPKCS1v15(craned.Reader, privateKey, crypto.SHA256, h.Sum(nil))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

func BuildUrl(params map[string]string, query url.Values, subRoutes ...string) string {
	url := constant.ApiDomain
	for _, route := range subRoutes {
		url += strings.TrimLeft(route, "/")
	}
	for key, param := range params {
		url = strings.ReplaceAll(url, "{"+key+"}", param)
	}
	if query != nil {
		url += "?"
		url += query.Encode()
	}
	return url
}
