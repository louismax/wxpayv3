package core

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/custom"
	"github.com/louismax/wxpayv3/utils"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path"
	"strings"
)

func (c *PayClient) UploadImage(filePath string) (*custom.RespUploadImage, error) {
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
	defer func() {
		_ = file.Close()
	}()
	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}
	var size = stats.Size()
	fBytes := make([]byte, size)
	bf := bufio.NewReader(file)

	if _, err = bf.Read(fBytes); err != nil {
		return nil, err
	}
	//二进制内容进行sha256计算得到的值
	h := sha256.New()
	h.Write(fBytes)

	//参数组装
	xName, _ := utils.GenerateNonce()
	req := custom.ReqUploadImage{
		Filename: fmt.Sprintf("%s%s", xName, fileSuffix),
		Sha256:   hex.EncodeToString(h.Sum(nil)),
	}
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	apiUrl := utils.BuildUrl(nil, nil, constant.ApiUploadImage)
	//获取签名
	authorization, err := c.Authorization(http.MethodPost, apiUrl, data)
	if err != nil {
		return nil, err
	}

	body := &bytes.Buffer{}             // 初始化body参数
	writer := multipart.NewWriter(body) // 实例化multipart

	part, err := CreateForm("meta", "application/json", writer)
	if err != nil {
		return nil, err
	}

	if _, err = part.Write(data); err != nil {
		return nil, err
	}

	parTF, err := CreateFormFile("file", req.Filename, ctp, writer) // 创建multipart 文件字段
	if err != nil {
		return nil, err
	}
	if _, err = parTF.Write(fBytes); err != nil {
		return nil, err
	}

	if err = writer.Close(); err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, apiUrl, body)
	if err != nil {
		return nil, err
	}
	//设置短连接
	request.Close = true

	request.Header.Set("Content-Type", "multipart/form-data")
	request.Header.Set("Accept", "*/*")
	request.Header.Set("Authorization", authorization)

	if c.PlatformSerialNo != "" {
		request.Header.Set("Wechatpay-Serial", c.PlatformSerialNo)
	}

	resp, err := c.HttpClient.Do(request)
	if resp != nil {
		defer func() {
			_ = resp.Body.Close()
		}()
	}
	if err != nil {
		return nil, err
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := custom.RespUploadImage{}
	err = json.Unmarshal(respData, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil

}

func CreateForm(key, contentType string, w *multipart.Writer) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s";`,
			escapeQuotes(key)))
	h.Set("Content-Type", contentType)
	return w.CreatePart(h)
}

var quoteEscape = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscape.Replace(s)
}
func CreateFormFile(fieldName, filename, contentType string, w *multipart.Writer) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(fieldName), escapeQuotes(filename)))
	h.Set("Content-Type", contentType)
	return w.CreatePart(h)
}
