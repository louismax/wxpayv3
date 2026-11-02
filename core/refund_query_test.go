package core

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestPaymentQueryRefundForPartner(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	for _, state := range []string{"SUCCESS", "PROCESSING", "ABNORMAL", "CLOSED"} {
		t.Run(state, func(t *testing.T) {
			client := &PayClient{MchId: "service-mch", ApiSerialNo: "serial", ApiPrivateKey: key}
			client.HttpClient = &http.Client{Transport: nativePaymentTransport(func(req *http.Request) (*http.Response, error) {
				if req.Method != http.MethodGet || req.URL.Path != "/v3/refund/domestic/refunds/refund-1" || req.URL.Query().Get("sub_mchid") != "sub-mch" {
					t.Fatalf("退款查询没有使用正确的服务商请求: %s %s", req.Method, req.URL)
				}
				body := fmt.Sprintf(`{"out_refund_no":"refund-1","transaction_id":"transaction-1","status":"%s","amount":{"total":100,"refund":100}}`, state)
				return &http.Response{StatusCode: http.StatusOK, Header: http.Header{}, Body: io.NopCloser(strings.NewReader(body))}, nil
			})}
			result, err := client.PaymentQueryRefund("refund-1", "sub-mch")
			if err != nil || result.Status != state || result.TransactionId != "transaction-1" || result.Amount.Refund != 100 {
				t.Fatalf("退款结果解析不正确: %+v err=%v", result, err)
			}
		})
	}
}
