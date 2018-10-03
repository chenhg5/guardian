package main

import (
	"net/http"
	"bytes"
	"encoding/json"
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
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	for key, header := range headers {
		header = GlobalVars.Replace(header)
		req.Header.Set(key, header)
	}

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	return res
}

func Get(url string, params interface{}, headers map[string]string) *http.Response {
	return Request("GET", url, params, headers)
}

func Post(url string, params interface{}, headers map[string]string) *http.Response {
	return Request("POST", url, params, headers)
}

func Delete(url string, params interface{}, headers map[string]string) *http.Response {
	return Request("DELETE", url, params, headers)
}

func Head(url string, params interface{}, headers map[string]string) *http.Response {
	return Request("HEAD", url, params, headers)
}

func Put(url string, params interface{}, headers map[string]string) *http.Response {
	return Request("PUT", url, params, headers)
}

func Options(url string, params interface{}, headers map[string]string) *http.Response {
	return Request("OPTIONS", url, params, headers)
}

func Request(method string, url string, params interface{}, headers map[string]string) *http.Response {
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

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}

	for key, header := range headers {
		header = GlobalVars.Replace(header)
		req.Header.Set(key, header)
	}

	res, _ := client.Do(req)

	if err != nil {
		panic(err)
	}
	return res
}

func FormGet(url string, params interface{}, headers map[string]string) *http.Response {
	return Form("GET", url, params, headers)
}

func FormPost(url string, params interface{}, headers map[string]string) *http.Response {
	return Form("POST", url, params, headers)
}

func Form(method string, url string, params interface{}, headers map[string]string) *http.Response {
	data := make(neturl.Values)
	formData := params.(map[string]interface{})

	var value string
	for k, v := range formData {
		value = v.(string)
		v = GlobalVars.Replace(value)
		data[k] = []string{value}
	}

	client := http.DefaultClient
	req, err := http.NewRequest(method, url, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for key, header := range headers {
		header = GlobalVars.Replace(header)
		req.Header.Set(key, header)
	}

	res, _ := client.Do(req)

	if err != nil {
		panic(err)
	}
	return res
}

func GetRequester(key string) RequestFunc {
	switch strings.ToLower(key) {
	case "json":
		return Json
	case "get":
		return Get
	case "post":
		return Post
	case "delete":
		return Delete
	case "put":
		return Put
	case "options":
		return Options
	case "formget":
		return FormGet
	case "formpost":
		return FormPost
	default:
		panic("wrong parameter")
	}
}
