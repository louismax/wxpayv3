package wxpayv3

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

// UploadImage 上传图片
func (this *Client) UploadImage(filePath string) (interface{}, error) {
	//获取文件名带后缀
	filenameWithSuffix := path.Base(filePath)
	//获取文件后缀
	fileSuffix := path.Ext(filenameWithSuffix)

	ctp := ""
	if strings.ToLower(fileSuffix) == ".jpg" {
		ctp = "image/jpg"
	} else if strings.ToLower(fileSuffix) == ".png" {
		ctp = "image/png"
	} else if strings.ToLower(fileSuffix) == ".bmp" {
		ctp = "image/png"
	} else {
		return nil, errors.New("暂不支持的文件类型")
	}

	//获取文件的二进制内容
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}
	var size int64 = stats.Size()
	fbytes := make([]byte, size)
	bf := bufio.NewReader(file)
	_, err = bf.Read(fbytes)
	if err != nil {
		return nil, err
	}

	//二进制内容进行sha256计算得到的值
	h := sha256.New()
	h.Write(fbytes)

	//参数组装
	req := ReqUploadImage{
		Filename: fmt.Sprintf("%s%s", GetGUID(), fileSuffix),
		Sha256:   hex.EncodeToString(h.Sum(nil)),
	}
	timeUnix, sjstr, str := AbsSignedStr(req)

	//获取签名
	signdata, err := this.RsaEncrypt([]byte(str))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("签名生成失败，ERR：%+v", err))
	}
	authstr := fmt.Sprintf("%s mchid=\"%s\",serial_no=\"%s\",nonce_str=\"%s\",signature=\"%s\",timestamp=\"%s\"", AuthType, this.Mchid, this.SerialNo, sjstr, signdata, strconv.FormatInt(timeUnix, 10))

	apiurl := ServerUrl + AssembleUrl(req)

	body := &bytes.Buffer{}             // 初始化body参数
	writer := multipart.NewWriter(body) // 实例化multipart

	part, err := CreateForm("meta", "application/json", writer)
	_, err = part.Write([]byte(req.RawJsonStr()))
	if err != nil {
		return nil, err
	}

	partf, err := CreateFormFile("file", req.Filename, ctp, writer) // 创建multipart 文件字段
	if err != nil {
		return nil, err
	}
	_, err = partf.Write(fbytes)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(req.Method(), apiurl, body)
	//设置短连接
	request.Close = true

	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "multipart/form-data")
	request.Header.Set("Accept", AcceptType)
	request.Header.Set("Authorization", authstr)
	if this.PFSerialno != "" {
		request.Header.Set("Wechatpay-Serial", this.PFSerialno)
	}

	resp, err := this.Client.Do(request)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if strings.Contains(string(data), "code") {
		errmsg := SysError{}
		err = json.Unmarshal([]byte(data), &errmsg)
		if err != nil {
			return nil, err
		}
		return errmsg, errors.New("Fail")
	} else {
		result := RespUploadImage{}
		err = json.Unmarshal(data, &result)
		if err != nil {
			return "", errors.New(fmt.Sprintf("Result解析失败:%+v", err))
		}
		return result, nil
	}
}
