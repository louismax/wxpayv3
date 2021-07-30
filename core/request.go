package core

import (
	"bytes"
	"io"
	"net/http"
)

func WithBaseHeader(req *http.Request, authorization string) *http.Request {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", authorization)
	return req
}

func NewRequest(authorization string, url string, method string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	return WithBaseHeader(req, authorization), nil
}

func SimpleRequest(client *http.Client, url string, method string, authorization string, body []byte, platformSerialNo string) (*http.Response, error) {
	if client == nil {
		client = http.DefaultClient
	}
	var requestBody io.Reader
	if body != nil {
		requestBody = bytes.NewReader(body)
	}
	request, err := NewRequest(authorization, url, method, requestBody)
	if err != nil {
		return nil, err
	}
	if platformSerialNo != "" {
		request.Header.Set("Wechatpay-Serial", platformSerialNo)
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
