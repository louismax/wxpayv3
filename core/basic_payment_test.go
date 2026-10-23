package core

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/louismax/wxpayv3/custom"
)

type nativePaymentTransport func(*http.Request) (*http.Response, error)

func (f nativePaymentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestNativePaymentForPartner(t *testing.T) {
	testNativePayment(t, true)
}

func TestNativePayment(t *testing.T) {
	testNativePayment(t, false)
}

func testNativePayment(t *testing.T, partner bool) {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name   string
		close  bool
		status int
		body   string
	}{
		{"下单成功", false, http.StatusOK, `{"code_url":"weixin://wxpay/bizpayurl/up?pr=test"}`},
		{"下单失败", false, http.StatusBadRequest, `{"code":"PARAM_ERROR","message":"参数错误"}`},
		{"下单响应格式错误", false, http.StatusOK, `invalid json`},
		{"关单成功", true, http.StatusNoContent, ""},
		{"关单失败", true, http.StatusBadRequest, `{"code":"ORDERPAID","message":"订单已支付"}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			client := &PayClient{
				MchId: "sp-mchid", ApiSerialNo: "test-serial", ApiPrivateKey: key,
				HttpClient: &http.Client{Transport: nativePaymentTransport(func(req *http.Request) (*http.Response, error) {
					calls++
					path := "/v3/pay/partner/transactions/native"
					if tc.close {
						path = "/v3/pay/partner/transactions/out-trade-no/order-123/close"
					}
					if !partner {
						path = strings.Replace(path, "/partner", "", 1)
					}
					if req.Method != http.MethodPost || req.URL.Path != path {
						t.Fatalf("请求不符合文档: %s %s", req.Method, req.URL.Path)
					}
					var body map[string]interface{}
					if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
						t.Fatal(err)
					}
					if partner && (body["sp_mchid"] != "sp-mchid" || body["sub_mchid"] != "sub-mchid") {
						t.Fatalf("商户参数错误: %v", body)
					}
					if !partner {
						if body["mchid"] != "sp-mchid" || (!tc.close && body["appid"] != "appid") {
							t.Fatalf("直连商户参数错误: %v", body)
						}
						for _, field := range []string{"sp_appid", "sp_mchid", "sub_appid", "sub_mchid"} {
							if _, ok := body[field]; ok {
								t.Errorf("直连模式不应发送服务商字段: %s", field)
							}
						}
					}
					if tc.close {
						fields := 1
						if partner {
							fields = 2
						}
						if len(body) != fields {
							t.Fatalf("关单包体应只包含商户号: %v", body)
						}
					} else {
						if body["out_trade_no"] != "order-123" || body["amount"].(map[string]interface{})["total"] != float64(100) {
							t.Fatalf("订单参数错误: %v", body)
						}
						for _, field := range []string{"payer", "sub_appid", "time_expire", "detail", "scene_info", "settle_info"} {
							if _, ok := body[field]; ok {
								t.Errorf("不应发送未填写的字段: %s", field)
							}
						}
					}
					return &http.Response{StatusCode: tc.status, Header: make(http.Header), Body: io.NopCloser(strings.NewReader(tc.body))}, nil
				})},
			}
			var callErr error
			if tc.close {
				if partner {
					callErr = client.PaymentCloseOrder("order-123", "sp-mchid", "sub-mchid")
				} else {
					callErr = client.PaymentCloseOrder("order-123", "sp-mchid")
				}
			} else {
				data := custom.ReqNativeOrdersForPartner{
					SpAppid: "sp-appid", SpMchid: "sp-mchid", SubMchid: "sub-mchid",
					Description: "测试商品", OutTradeNo: "order-123", NotifyUrl: "https://example.com/notify",
				}
				data.Amount.Total = 100
				var resp *custom.RespNativeOrders
				if partner {
					resp, callErr = client.NativeOrdersForPartner(data)
				} else {
					directData := custom.ReqNativeOrders{
						Appid: "appid", Mchid: "sp-mchid", Description: data.Description,
						OutTradeNo: data.OutTradeNo, NotifyUrl: data.NotifyUrl,
					}
					directData.Amount.Total = 100
					resp, callErr = client.NativeOrders(directData)
				}
				if callErr == nil && (resp == nil || resp.CodeUrl != "weixin://wxpay/bizpayurl/up?pr=test") {
					t.Fatalf("二维码链接解析错误: %+v", resp)
				}
			}
			wantErr := tc.status == http.StatusBadRequest || tc.body == "invalid json"
			if (callErr != nil) != wantErr {
				t.Fatalf("错误返回不符: %v", callErr)
			}
			if calls != 1 {
				t.Fatalf("请求次数不符: %d", calls)
			}
		})
	}
}
