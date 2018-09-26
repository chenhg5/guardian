package guardian

import "net/http"

type RequestFunc func(url string, params interface{}, header map[string]string) http.Response

func Json(url string, params interface{}, header map[string]string) http.Response {
	// TODO: 发送json post请求
	return http.Response{}
}

func Get(url string, params interface{}, header map[string]string) http.Response {
	// TODO: 发送get请求
	return http.Response{}
}

func FormGet(url string, params interface{}, header map[string]string) http.Response {
	// TODO: 发送form get请求
	return http.Response{}
}

func FormPost(url string, params interface{}, header map[string]string) http.Response {
	// TODO: 发送form post请求
	return http.Response{}
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