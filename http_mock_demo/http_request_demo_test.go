/*
 * @Author: linhuaqing
 * @Date: 2021-07-15 22:03:45
 * @LastEditTime: 2021-07-16 15:59:09
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-demo/http_mock_demo/http_request_demo_test.go
 * stay hungry stay foolish
 */
package http_request_demo

import (
	"errors"
	"net/http"
	"strconv"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetAPIResponse(t *testing.T) {
	var api string
	var responseBody string
	var err error

	// url nil test
	api = "://新"
	responseBody, err = getAPIResponse(api)
	if err == nil {
		t.Errorf("url nil test failed! err: %v", err)
	}
	if responseBody != "" {
		t.Errorf("url nil test failed! Unexpect response!")
	}

	// http mock test
	var mockResponse string
	api = "http://www.baidu.com"
	mockResponse = "mock response body"
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", api, httpmock.NewStringResponder(200, string(mockResponse)))
	responseBody, err = getAPIResponse(api)
	if err != nil {
		t.Errorf("second test failed! err: %v", err)
	}
	if responseBody != mockResponse {
		t.Errorf("second test failed! Unexpect response!")
	}

	// request fail test
	api = "http://www.baidu.com/fail"
	httpmock.RegisterResponder("GET", api,
		func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				Status:        strconv.Itoa(200),
				StatusCode:    200,
				ContentLength: -1,
			}, errors.New("fail error!")
		})
	responseBody, err = getAPIResponse(api)
	if err == nil {
		t.Errorf("request fail test failed! err: %v", err)
	}
	if responseBody != "" {
		t.Errorf("request fail test failed! Unexpect response! response: %s", responseBody)
	}

	// wrong status test
	api = "http://www.baidu.com/wrongstatus"
	mockResponseWrong := "mock response body"
	httpmock.RegisterResponder("GET", api, httpmock.NewStringResponder(404, string(mockResponseWrong)))
	responseBody, err = getAPIResponse(api)
	if err == nil {
		t.Errorf("wrong request test failed! err: %v", err)
	}
	if responseBody != "" {
		t.Errorf("wrong request test failed! Unexpect response! response: %s", responseBody)
	}
}

// func TestGetAPIResponse(t *testing.T) {
// 	var api string
// 	var responseBody string
// 	var err error

// 	// first test
// 	api = "http://www.baidu.com"
// 	responseBody, err = getAPIResponse(api)
// 	if err != nil {
// 		t.Errorf("first test failed! err: %v", err)
// 	}
// 	if responseBody == "" {
// 		t.Errorf("first test failed! Unexpect response!")
// 	}
// }
