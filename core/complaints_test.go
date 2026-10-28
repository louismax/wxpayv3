package core

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/louismax/wxpayv3/custom"
)

type complaintTransport func(*http.Request) (*http.Response, error)

func (f complaintTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func complaintTestClient(t *testing.T) (*PayClient, *rsa.PrivateKey) {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	return &PayClient{
		MchId: "service-provider", ApiSerialNo: "merchant-serial", ApiPrivateKey: key,
		WechatPayPublicKeyID: "PUB_KEY_ID_TEST", WechatPayPublicKey: &key.PublicKey,
		ApiV3Key: "0123456789abcdef0123456789abcdef",
	}, key
}

func complaintSignedHeaders(t *testing.T, key *rsa.PrivateKey, body []byte) http.Header {
	t.Helper()
	digest := sha256.Sum256(append([]byte("1700000000\nresponse-nonce\n"), append(body, '\n')...))
	signature, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, digest[:])
	if err != nil {
		t.Fatal(err)
	}
	header := make(http.Header)
	header.Set("Wechatpay-Serial", "PUB_KEY_ID_TEST")
	header.Set("Wechatpay-Timestamp", "1700000000")
	header.Set("Wechatpay-Nonce", "response-nonce")
	header.Set("Wechatpay-Signature", base64.StdEncoding.EncodeToString(signature))
	return header
}

func complaintVerifyRequest(t *testing.T, key *rsa.PrivateKey, req *http.Request, body []byte) {
	t.Helper()
	values := map[string]string{}
	for _, match := range regexp.MustCompile(`(\w+)="([^"]*)"`).FindAllStringSubmatch(req.Header.Get("Authorization"), -1) {
		values[match[1]] = match[2]
	}
	if values["mchid"] != "service-provider" {
		t.Fatalf("请求未使用服务商身份: %s", values["mchid"])
	}
	message := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", req.Method, req.URL.RequestURI(), values["timestamp"], values["nonce_str"], body)
	digest := sha256.Sum256([]byte(message))
	signature, err := base64.StdEncoding.DecodeString(values["signature"])
	if err != nil {
		t.Fatal(err)
	}
	if err = rsa.VerifyPKCS1v15(&key.PublicKey, crypto.SHA256, digest[:], signature); err != nil {
		t.Fatalf("请求签名与实际URI或请求体不一致: %v", err)
	}
}

func TestComplaintRequests(t *testing.T) {
	client, key := complaintTestClient(t)
	const base = "/v3/merchant-service/complaints-v2"
	for _, name := range []string{"list-all", "list-merchant", "detail", "history", "reply", "complete", "immediate", "api-error", "invalid-json"} {
		t.Run(name, func(t *testing.T) {
			response := `{"complaint_id":"complaint-1","complainted_mchid":"sub-mchid","payer_phone":"13800000000","complaint_order_info":[{"amount":123}]}`
			path, method := base+"/complaint-1", http.MethodGet
			status, calls := http.StatusOK, 0
			switch name {
			case "list-all", "list-merchant":
				path = base
				response = `{"data":[{"complaint_id":"complaint-1","complainted_mchid":"sub-mchid"}],"total_count":1}`
			case "history":
				path += "/negotiation-historys"
				response = `{"data":[{"log_id":"log-1","user_appy_platform_service_reason":"响应不及时","complaint_media_list":{"media_type":"OPERATION_IMAGE","media_url":["https://api.mch.weixin.qq.com/v3/merchant-service/images/a"]},"normal_message":{"blocks":[{"type":"TEXT","text":{"text":"已联系"}}]},"click_message":{"action_id":"a","custom_data":"b"}}],"total_count":1}`
			case "reply":
				path, method, status, response = path+"/response", http.MethodPost, http.StatusNoContent, ""
			case "complete":
				path, method, status, response = path+"/complete", http.MethodPost, http.StatusNoContent, ""
			case "immediate":
				path, method, response = path+"/response-immediate-service", http.MethodPost, `{"log_id":"log-new"}`
			case "api-error":
				status, response = http.StatusForbidden, `{"code":"NO_AUTH","message":"无权访问"}`
			case "invalid-json":
				response = "invalid json"
			}
			client.HttpClient = &http.Client{Transport: complaintTransport(func(req *http.Request) (*http.Response, error) {
				calls++
				if req.Method != method || req.URL.Path != path {
					t.Fatalf("请求错误: %s %s", req.Method, req.URL.Path)
				}
				var requestBody []byte
				if req.Body != nil {
					requestBody, _ = io.ReadAll(req.Body)
				}
				complaintVerifyRequest(t, key, req, requestBody)
				query := req.URL.Query()
				if name == "list-all" || name == "list-merchant" {
					if query.Get("begin_date") != "2026-10-01" || query.Get("end_date") != "2026-10-27" || query.Get("limit") != "50" || query.Get("offset") != "10" {
						t.Fatalf("列表分页或日期参数错误: %v", query)
					}
					_, present := query["complainted_mchid"]
					if present != (name == "list-merchant") || (present && query.Get("complainted_mchid") != "sub-mchid") {
						t.Fatalf("全部子商户与单商户查询参数错误: %v", query)
					}
				}
				if name == "history" && (query.Get("limit") != "300" || query.Get("offset") != "10") {
					t.Fatalf("协商历史分页错误: %v", query)
				}
				if method == http.MethodPost {
					var body map[string]interface{}
					if err := json.Unmarshal(requestBody, &body); err != nil {
						t.Fatal(err)
					}
					if body["complainted_mchid"] != "sub-mchid" {
						t.Fatalf("缺少被诉子商户: %v", body)
					}
					if name == "complete" && len(body) != 1 {
						t.Fatalf("完成投诉不应发送其他字段: %v", body)
					}
					if name == "reply" && (body["response_content"] != "已联系用户" || body["response_images"].([]interface{})[0] != "media-1") {
						t.Fatalf("回复正文或凭证错误: %v", body)
					}
					if name == "immediate" && body["message"].(map[string]interface{})["blocks"].([]interface{})[0].(map[string]interface{})["type"] != "TEXT" {
						t.Fatalf("即时服务消息格式错误: %v", body)
					}
				}
				data := []byte(response)
				return &http.Response{StatusCode: status, Header: complaintSignedHeaders(t, key, data), Body: io.NopCloser(bytes.NewReader(data))}, nil
			})}
			var err error
			switch name {
			case "list-all", "list-merchant":
				var result *custom.RespComplaintsList
				if name == "list-all" {
					result, err = client.QueryComplaintsList("2026-10-01", "2026-10-27", 50, 10)
				} else {
					result, err = client.QueryComplaintsList("2026-10-01", "2026-10-27", 50, 10, "sub-mchid")
				}
				if err == nil && (result.TotalCount != 1 || result.Data[0].ComplaintedMchid != "sub-mchid") {
					t.Fatalf("列表遗漏子商户信息: %+v", result)
				}
			case "history":
				var result *custom.RespComplaintNegotiationHistory
				result, err = client.QueryComplaintNegotiationHistory("complaint-1", 300, 10)
				if err == nil && (result.Data[0].LogID != "log-1" || result.Data[0].NormalMessage.Blocks[0].Text.Text != "已联系" || len(result.Data[0].ClickMessage) == 0) {
					t.Fatalf("协商记录解析不完整: %+v", result)
				}
			case "reply":
				err = client.ResponseComplaint("complaint-1", custom.ReqResponseComplaint{ComplaintedMchid: "sub-mchid", ResponseContent: "已联系用户", ResponseImages: []string{"media-1"}})
			case "complete":
				err = client.CompleteComplaint("complaint-1", custom.ReqCompleteComplaint{ComplaintedMchid: "sub-mchid"})
			case "immediate":
				var result *custom.RespResponseComplaintImmediateService
				result, err = client.ResponseComplaintImmediateService("complaint-1", custom.ReqResponseComplaintImmediateService{
					ComplaintedMchid: "sub-mchid", Message: custom.ComplaintMessage{Blocks: []custom.ComplaintMessageBlock{{Type: "TEXT", Text: &custom.ComplaintMessageText{Text: "已联系用户"}}}},
				})
				if err == nil && result.LogID != "log-new" {
					t.Fatalf("未返回即时服务流水号: %+v", result)
				}
			default:
				var result *custom.RespComplaintDetail
				result, err = client.QueryComplaintDetail("complaint-1")
				if name == "detail" && err == nil && (result.ComplaintedMchid != "sub-mchid" || result.ComplaintOrderInfo[0].Amount != 123) {
					t.Fatalf("详情解析错误: %+v", result)
				}
			}
			if (err != nil) != (name == "api-error" || name == "invalid-json") {
				t.Fatalf("错误处理不符: %v", err)
			}
			if calls != 1 {
				t.Fatalf("请求次数错误: %d", calls)
			}
		})
	}
}

func TestComplaintImage(t *testing.T) {
	client, key := complaintTestClient(t)
	image := []byte{0x89, 'P', 'N', 'G', 0, 0xff, 0x10}
	filePath := filepath.Join(t.TempDir(), "proof.PNG")
	if err := os.WriteFile(filePath, image, 0600); err != nil {
		t.Fatal(err)
	}
	client.HttpClient = &http.Client{Transport: complaintTransport(func(req *http.Request) (*http.Response, error) {
		response := image
		if req.Method == http.MethodPost {
			if req.URL.Path != "/v3/merchant-service/images/upload" {
				t.Fatalf("使用了错误的图片上传接口: %s", req.URL)
			}
			if err := req.ParseMultipartForm(3 * 1024 * 1024); err != nil {
				t.Fatal(err)
			}
			defer req.MultipartForm.RemoveAll()
			meta := custom.ReqUploadImage{}
			metaJSON := []byte(req.FormValue("meta"))
			if err := json.Unmarshal(metaJSON, &meta); err != nil {
				t.Fatal(err)
			}
			digest := sha256.Sum256(image)
			if meta.Sha256 != hex.EncodeToString(digest[:]) {
				t.Fatal("图片摘要错误")
			}
			file, header, err := req.FormFile("file")
			if err != nil {
				t.Fatal(err)
			}
			defer file.Close()
			data, _ := io.ReadAll(file)
			if !bytes.Equal(data, image) || header.Filename != meta.Filename || header.Header.Get("Content-Type") != "image/png" {
				t.Fatal("上传图片内容或元信息不一致")
			}
			complaintVerifyRequest(t, key, req, metaJSON)
			response = []byte(`{"media_id":"complaint-media"}`)
		} else {
			if req.URL.RequestURI() != "/v3/merchant-service/images/media%3D%3D" || req.Header.Get("Accept") != "*/*" {
				t.Fatalf("图片URL未原样使用: %s", req.URL)
			}
			complaintVerifyRequest(t, key, req, nil)
		}
		return &http.Response{StatusCode: http.StatusOK, Header: complaintSignedHeaders(t, key, response), Body: io.NopCloser(bytes.NewReader(response))}, nil
	})}
	result, err := client.UploadComplaintImage(filePath)
	if err != nil || result.MediaId != "complaint-media" {
		t.Fatalf("上传失败: %+v %v", result, err)
	}
	data, err := client.DownloadComplaintImage("https://api.mch.weixin.qq.com/v3/merchant-service/images/media%3D%3D")
	if err != nil || !bytes.Equal(data, image) {
		t.Fatalf("图片二进制下载失败: %v", err)
	}
	if _, err = client.DownloadComplaintImage("https://example.com/v3/merchant-service/images/media"); err == nil {
		t.Fatal("不应向非微信域名发送带签名的请求")
	}
	if err = os.WriteFile(filePath, make([]byte, 2*1024*1024+1), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err = client.UploadComplaintImage(filePath); err == nil {
		t.Fatal("未拒绝超过2M的图片")
	}
	client.HttpClient = &http.Client{Transport: complaintTransport(func(req *http.Request) (*http.Response, error) {
		body := `{"code":"NO_AUTH","message":"无权访问图片"}`
		return &http.Response{StatusCode: http.StatusForbidden, Header: make(http.Header), Body: io.NopCloser(strings.NewReader(body))}, nil
	})}
	if _, err = client.DownloadComplaintImage("https://api.mch.weixin.qq.com/v3/merchant-service/images/media"); err == nil {
		t.Fatal("未返回图片下载的渠道错误")
	}
}

func TestComplaintNotification(t *testing.T) {
	client, key := complaintTestClient(t)
	block, err := aes.NewCipher([]byte(client.ApiV3Key))
	if err != nil {
		t.Fatal(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		t.Fatal(err)
	}
	notification := custom.ComplaintNotification{ID: "notify-1", EventType: "COMPLAINT.STATE_CHANGE", ResourceType: "encrypt-resource"}
	notification.Resource = custom.ComplaintEncryptedResource{
		Algorithm: "AEAD_AES_256_GCM", Nonce: "0123456789ab", AssociatedData: "complaint", OriginalType: "complaint",
	}
	notification.Resource.Ciphertext = base64.StdEncoding.EncodeToString(gcm.Seal(nil, []byte(notification.Resource.Nonce), []byte(`{"complaint_id":"complaint-1","action_type":"USER_RESPONSE"}`), []byte("complaint")))
	body, err := json.Marshal(notification)
	if err != nil {
		t.Fatal(err)
	}
	header := complaintSignedHeaders(t, key, body)
	envelope, resource, err := client.ParseComplaintNotification(header, body)
	if err != nil || envelope.ID != "notify-1" || resource.ComplaintID != "complaint-1" || resource.ActionType != "USER_RESPONSE" {
		t.Fatalf("投诉通知解密失败: %+v %+v %v", envelope, resource, err)
	}
	certClient := *client
	certClient.PlatformCertMap = map[string]*x509.Certificate{"test-cert-serial": {PublicKey: &key.PublicKey}}
	certHeader := header.Clone()
	certHeader.Set("Wechatpay-Serial", "test-cert-serial")
	if _, _, err = certClient.ParseComplaintNotification(certHeader, body); err != nil {
		t.Fatalf("平台证书通知验签失败: %v", err)
	}
	for _, name := range []string{"missing-serial", "unknown-key", "unknown-cert", "missing-signature", "invalid-signature", "invalid-nonce", "invalid-key", "invalid-ciphertext"} {
		t.Run(name, func(t *testing.T) {
			copyClient := *client
			data := body
			headers := header.Clone()
			switch name {
			case "missing-serial":
				headers.Del("Wechatpay-Serial")
			case "unknown-key":
				headers.Set("Wechatpay-Serial", "PUB_KEY_ID_OTHER")
			case "unknown-cert":
				headers.Set("Wechatpay-Serial", "unknown-cert")
			case "missing-signature":
				headers.Del("Wechatpay-Signature")
			case "invalid-signature":
				data = append(append([]byte{}, body...), ' ')
			case "invalid-nonce":
				changed := notification
				changed.Resource.Nonce = "bad"
				data, _ = json.Marshal(changed)
				headers = complaintSignedHeaders(t, key, data)
			case "invalid-key":
				copyClient.ApiV3Key = "bad"
			case "invalid-ciphertext":
				changed := notification
				changed.Resource.Ciphertext = "invalid base64"
				data, _ = json.Marshal(changed)
				headers = complaintSignedHeaders(t, key, data)
			}
			if _, _, err := copyClient.ParseComplaintNotification(headers, data); err == nil {
				t.Fatal("无效投诉通知未被拒绝")
			}
		})
	}
}

func TestComplaintWriteDoesNotRetry(t *testing.T) {
	client, _ := complaintTestClient(t)
	for _, name := range []string{"reply", "immediate", "complete"} {
		t.Run(name, func(t *testing.T) {
			calls := 0
			client.HttpClient = &http.Client{Transport: complaintTransport(func(req *http.Request) (*http.Response, error) {
				calls++
				return nil, fmt.Errorf("模拟请求发送后连接中断")
			})}
			var err error
			switch name {
			case "reply":
				err = client.ResponseComplaint("complaint-1", custom.ReqResponseComplaint{ComplaintedMchid: "sub-mchid", ResponseContent: "已联系"})
			case "immediate":
				_, err = client.ResponseComplaintImmediateService("complaint-1", custom.ReqResponseComplaintImmediateService{ComplaintedMchid: "sub-mchid"})
			case "complete":
				err = client.CompleteComplaint("complaint-1", custom.ReqCompleteComplaint{ComplaintedMchid: "sub-mchid"})
			}
			if err == nil || calls != 1 {
				t.Fatalf("结果不确定时不应自动重复写入: calls=%d err=%v", calls, err)
			}
		})
	}
}
