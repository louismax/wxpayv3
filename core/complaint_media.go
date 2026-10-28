package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/louismax/wxpayv3/constant"
	"github.com/louismax/wxpayv3/custom"
	"github.com/louismax/wxpayv3/utils"
)

// UploadComplaintImage 使用投诉专用接口上传图片，不能使用普通商户资料上传的media_id替代。
func (c *PayClient) UploadComplaintImage(filePath string) (*custom.RespUploadImage, error) {
	contentType := ""
	suffix := strings.ToLower(filepath.Ext(filePath))
	switch suffix {
	case ".jpg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	case ".bmp":
		contentType = "image/bmp"
	default:
		return nil, fmt.Errorf("投诉图片仅支持JPG、PNG、BMP格式")
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	image, err := io.ReadAll(io.LimitReader(file, 2*1024*1024+1))
	if err != nil {
		return nil, err
	}
	if len(image) == 0 || len(image) > 2*1024*1024 {
		return nil, fmt.Errorf("投诉图片不能为空且大小不能超过2M")
	}
	nonce, err := utils.GenerateNonce()
	if err != nil {
		return nil, err
	}
	digest := sha256.Sum256(image)
	meta := custom.ReqUploadImage{Filename: nonce + suffix, Sha256: hex.EncodeToString(digest[:])}
	metaJSON, err := json.Marshal(meta)
	if err != nil {
		return nil, err
	}
	apiURL := utils.BuildUrl(nil, nil, constant.APIComplaintImageUpload)
	// 图片上传的签名主体为meta JSON，而不是整个multipart报文。
	authorization, err := c.Authorization(http.MethodPost, apiURL, metaJSON)
	if err != nil {
		return nil, err
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	metaPart, err := CreateForm("meta", "application/json", writer)
	if err != nil {
		return nil, err
	}
	if _, err = metaPart.Write(metaJSON); err != nil {
		return nil, err
	}
	filePart, err := CreateFormFile("file", meta.Filename, contentType, writer)
	if err != nil {
		return nil, err
	}
	if _, err = filePart.Write(image); err != nil {
		return nil, err
	}
	if err = writer.Close(); err != nil {
		return nil, err
	}
	req, err := NewRequest(authorization, apiURL, http.MethodPost, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if c.WechatPayPublicKeyID != "" {
		req.Header.Set("Wechatpay-Serial", c.WechatPayPublicKeyID)
	} else if c.DefaultPlatformSerialNo != "" {
		req.Header.Set("Wechatpay-Serial", c.DefaultPlatformSerialNo)
	}
	client := c.HttpClient
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err = c.VerifyResponse(resp.StatusCode, &resp.Header, responseBody); err != nil {
		return nil, err
	}
	result := custom.RespUploadImage{}
	if err = json.Unmarshal(responseBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DownloadComplaintImage 使用渠道返回的原始URL，保留媒体ID中的转义字符。
func (c *PayClient) DownloadComplaintImage(imageURL string) ([]byte, error) {
	parsed, err := url.Parse(imageURL)
	if err != nil {
		return nil, err
	}
	if parsed.Scheme != "https" || parsed.User != nil || parsed.Fragment != "" ||
		(parsed.Host != "api.mch.weixin.qq.com" && parsed.Host != "api2.mch.weixin.qq.com") ||
		!strings.HasPrefix(parsed.Path, "/v3/merchant-service/images/") ||
		strings.TrimPrefix(parsed.Path, "/v3/merchant-service/images/") == "" {
		return nil, fmt.Errorf("请使用投诉详情或协商历史返回的微信支付图片URL")
	}
	authorization, err := c.Authorization(http.MethodGet, imageURL, nil)
	if err != nil {
		return nil, err
	}
	req, err := NewRequest(authorization, imageURL, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "*/*")
	client := c.HttpClient
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err = c.VerifyResponse(resp.StatusCode, &resp.Header, body); err != nil {
		return nil, err
	}
	return body, nil
}
