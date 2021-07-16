/*
 * @Author: linhuaqing
 * @Date: 2021-07-15 21:57:12
 * @LastEditTime: 2021-07-16 15:24:00
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-demo/http_mock_demo/http_request_demo.go
 * stay hungry stay foolish
 */
package http_request_demo

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func getAPIResponse(url string) (string, error) {
	var err error
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	myClient := &http.Client{}
	response, err := myClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return "", errors.New("response not 200!")
	}
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}
