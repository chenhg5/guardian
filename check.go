package guardian

import (
	"net/http"
	"reflect"
)

func CheckResponse(actual *http.Response, expect TableResponse) (bool, string) {

	// TODO: 比照响应，返回结果和结果字符串

	return true, ""
}

func CheckMysql(actual []map[string]interface{}, expect []map[string]interface{}) (bool, string) {

	// TODO: 比照mysql，返回结果和结果字符串

	return true, ""
}

func CheckRedis(actual map[string]interface{}, expect map[string]interface{}) (bool, string) {

	// TODO: 比照redis，返回结果和结果字符串

	return true, ""
}

func DeepCheckMap(actual map[string]interface{}, expect map[string]interface{}) bool {
	return reflect.DeepEqual(actual, expect)
}