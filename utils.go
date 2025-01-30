package main

import (
	"crypto/tls"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/liornabat-sealights/ginkgo-test-runner/lib/types"
	"time"
)

var restyClient = resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

func NewRestRequest() *resty.Request {

	return restyClient.NewRequest()
}

func call(path string, a, b string) (*types.ResultResponse, error) {
	time.Sleep(4 * time.Second)
	result := &types.ResultResponse{}
	resp, err := NewRestRequest().SetQueryParams(map[string]string{
		"a": a,
		"b": b,
	}).SetResult(result).
		Get("http://localhost:10000" + path)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("response: %s", resp.String())
	}
	return result, nil
}
