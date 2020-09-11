package wxpayv3

import (
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	mathrand "math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Client struct {
	Mchid      string            //商户号
	SerialNo   string            //商户证书编号
	Priv       interface{}       //商户证书私钥
	Pubc       *x509.Certificate //商户证书公钥
	PFSerialno string            //平台证书编号
	PFPubc     *x509.Certificate //平台证书公钥

	Client *http.Client
}

// New 初始化微信支付V3客户端
// mchid 商户号
// serial_no 证书编号
// CertPath 证书文件路径  apiclient_key.pem
// pubpath
func New(mchid, serial_no, CertPath, pubpath string) (client *Client, err error) {
	client = &Client{}
	if mchid == "" {
		return nil, errors.New("商户号无效！")
	}
	client.Mchid = mchid
	if serial_no == "" {
		return nil, errors.New("证书编号无效！")
	}
	client.SerialNo = serial_no

	pfxData, err := ioutil.ReadFile(CertPath)
	if err != nil {
		return nil, errors.New("私钥证书文件读取失败！")
	}
	block, _ := pem.Decode(pfxData)

	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, errors.New("解码包含私钥的PEM块失败！")
	}

	client.Priv, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	pfxData_pub, err := ioutil.ReadFile(pubpath)
	if err != nil {
		return nil, errors.New("公钥证书文件读取失败！")
	}

	block_pub, _ := pem.Decode(pfxData_pub)

	if block_pub == nil || block_pub.Type != "CERTIFICATE" {
		return nil, errors.New("解码包含公钥的PEM块失败！")
	}

	client.Pubc, err = x509.ParseCertificate(block_pub.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	client.Client = http.DefaultClient
	return client, nil
}

//InitCertificate 平台公钥初始化，请在获取平台证书后使用
func (this *Client) InitCertificate(info CertificateInfo) {
	this.PFSerialno = info.Serial_no
	block_pub, _ := pem.Decode([]byte(info.Publickey))

	if block_pub == nil || block_pub.Type != "CERTIFICATE" {
		log.Println("解码包含平台公钥的PEM块失败！")
		return
	}
	var err error
	this.PFPubc, err = x509.ParseCertificate(block_pub.Bytes)
	if err != nil {
		log.Fatal(err)
	}
}

func (this *Client) doRequest(param Param, result interface{}) (string, error) {

	timeUnix, sjstr, str := AbsSignedStr(param)

	signdata, err := this.RsaEncrypt([]byte(str))
	if err != nil {
		return "", errors.New(fmt.Sprintf("签名生成失败，ERR：%+v", err))
	}

	authstr := fmt.Sprintf("%s mchid=\"%s\",serial_no=\"%s\",nonce_str=\"%s\",signature=\"%s\",timestamp=\"%s\"", AuthType, this.Mchid, this.SerialNo, sjstr, signdata, strconv.FormatInt(timeUnix, 10))

	//fmt.Println(fmt.Sprintf("authstr:%+v", authstr))

	apiurl := ServerUrl + AssembleUrl(param)

	//fmt.Println(fmt.Sprintf("apiurl:%+v", apiurl))
	//fmt.Println(fmt.Sprintf("apiurl:%+v", apiurl))

	req, err := http.NewRequest(param.Method(), apiurl, bytes.NewBuffer([]byte(param.RawJsonStr())))
	//设置短连接
	req.Close = true

	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", ContentType)
	req.Header.Set("Accept", AcceptType)
	req.Header.Set("Authorization", authstr)
	if this.PFSerialno != "" {
		req.Header.Set("Wechatpay-Serial", this.PFSerialno)
	}

	resp, err := this.Client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", err
	}

	//无需解析的返回请求状态，如PATCH
	if result == nil {
		if resp.StatusCode == 204 {
			return resp.Status, nil
		} else {
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return "", err
			}
			err = json.Unmarshal(data, &result)
			if err != nil {
				return "", errors.New("Result解析失败")
			}
			fmt.Println(string(data))
			return string(data), nil
			//return resp.Status, errors.New(fmt.Sprintf("请求失败，错误码:%d", resp.StatusCode))
		}
	} else {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		err = json.Unmarshal(data, &result)
		if err != nil {
			return "", errors.New("Result解析失败")
		}
		return string(data), nil
	}
}

// RsaEncrypt SHA256 with RSA加密
func (this *Client) RsaEncrypt(origData []byte) (string, error) {
	h := crypto.Hash.New(crypto.SHA256)
	h.Write(origData)
	hashed := h.Sum(nil)
	// 进行rsa加密签名
	signedData, err := rsa.SignPKCS1v15(rand.Reader, this.Priv.(*rsa.PrivateKey), crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signedData), nil
}

// RsaDecrypt  OAEP RSA解密
func (this *Client) RsaDecrypt(ciphertext string) (string, error) {
	cipherdata, _ := base64.StdEncoding.DecodeString(ciphertext)
	rng := rand.Reader

	plaintext, err := rsa.DecryptOAEP(sha1.New(), rng, this.Priv.(*rsa.PrivateKey), cipherdata, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
		return "", err
	}
	return string(plaintext), nil
}

//必须在平台证书初始化后使用
//RsaOAEPEncrypt OAEP RSA 平台公钥加密(敏感字段)
func (this *Client) RsaOAEPEncrypt(plaintext string) (string, error) {
	if this.PFSerialno == "" {
		return "", errors.New("请先初始化平台证书")
	}
	secretMessage := []byte(plaintext)
	rng := rand.Reader

	cipherdata, err := rsa.EncryptOAEP(sha1.New(), rng, this.PFPubc.PublicKey.(*rsa.PublicKey), secretMessage, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
		return "", err
	}

	ciphertext := base64.StdEncoding.EncodeToString(cipherdata)
	return ciphertext, nil
}

// AssembleUrl 组装绝对url
func AssembleUrl(param Param) string {
	var suffixUrl = param.APIUrl()
	if len(param.Params()) > 0 {
		suffixUrl = suffixUrl + "?"
		inx := 1
		for k, v := range param.Params() {
			if inx < len(param.Params()) {
				suffixUrl = suffixUrl + k + "=" + v + "&"
			} else {
				suffixUrl = suffixUrl + k + "=" + v
			}
			inx++
		}
	}
	return suffixUrl
}

func AbsSignedStr(param Param) (int64, string, string) {
	timeUnix := time.Now().Unix()
	sjstr := GetRandomString(32)
	message := ""
	data := ""
	if param.Method() != "GET" {
		data = param.RawJsonStr()
		//fmt.Println(data)
	}
	message = param.Method() + "\n" + AssembleUrl(param) + "\n" + strconv.FormatInt(timeUnix, 10) + "\n" + sjstr + "\n" + data + "\n"

	return timeUnix, sjstr, message
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//GetGUID 产生GUID
func GetGUID() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return getMd5String(base64.URLEncoding.EncodeToString(b))
}

func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
