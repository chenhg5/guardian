package main

import (
	"net/http"
	"bytes"
	"github.com/gin-gonic/gin/json"
	neturl "net/url"
	"strings"
)

type RequestFunc func(url string, params interface{}, headers map[string]string) *http.Response

func Json(url string, params interface{}, headers map[string]string) *http.Response {
	var (
		postDataByte []byte
		err          error
	)

	if postDataByte, err = json.Marshal(params); err != nil {
		return nil
	}

	// 替代全局变量
	postDataByte = []byte(GlobalVars.Replace(string(postDataByte)))

	client := http.DefaultClient

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postDataByte))
	if err != nil {
		return nil
	}
	req.Header.Set("Content-Type", "application/json")

	for key, header := range headers {
		header = GlobalVars.Replace(header)
		req.Header.Set(key, header)
	}

	res, err := client.Do(req)

	if err != nil {
		return nil
	}

	return res
}

func Get(url string, params interface{}, headers map[string]string) *http.Response {
	var count = 0
	data := params.(map[string]interface{})
	var value string
	for k, v := range data {
		value = v.(string)
		v = GlobalVars.Replace(value)
		if count == 0 {
			url += "?" + k + "=" + value
		} else {
			url += "&" + k + "=" + value
		}
		count++
	}

	client := http.DefaultClient

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}

	for key, header := range headers {
		header = GlobalVars.Replace(header)
		req.Header.Set(key, header)
	}

	res, _ := client.Do(req)

	if err != nil {
		return nil
	}
	return res
}

func FormGet(url string, params interface{}, headers map[string]string) *http.Response {
	data := make(neturl.Values)
	formData := params.(map[string]interface{})
	var value string
	for k, v := range formData {
		value = v.(string)
		v = GlobalVars.Replace(value)
		data[k] = []string{value}
	}

	client := http.DefaultClient
	req, err := http.NewRequest("GET", url, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for key, header := range headers {
		header = GlobalVars.Replace(header)
		req.Header.Set(key, header)
	}

	res, _ := client.Do(req)

	if err != nil {
		return nil
	}
	return res
}

func FormPost(url string, params interface{}, headers map[string]string) *http.Response {
	data := make(neturl.Values)
	formData := params.(map[string]interface{})

	var value string
	for k, v := range formData {
		value = v.(string)
		v = GlobalVars.Replace(value)
		data[k] = []string{value}
	}

	client := http.DefaultClient
	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for key, header := range headers {
		header = GlobalVars.Replace(header)
		req.Header.Set(key, header)
	}

	res, _ := client.Do(req)

	if err != nil {
		return nil
	}
	return res
}

func GetRequester(key string) RequestFunc {
	switch key {
	case "json":
		return Json
	case "get":
		return Get
	case "formget":
		return FormGet
	case "formpost":
		return FormPost
	default:
		panic("wrong parameter")
	}
}
