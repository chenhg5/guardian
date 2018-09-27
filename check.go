package guardian

import (
	"net/http"
	"reflect"
	"encoding/json"
)

func CheckResponse(actual *http.Response, expect TableResponse) (bool, string) {

	var bodyByte []byte
	actual.Body.Read(bodyByte)
	defer actual.Body.Close()

	if expectBody, ok := expect.Body.(string); ok {
		if string(bodyByte) == expectBody {
			return true, ""
		} else {
			return false, "not match"
		}
	} else if expectBody, ok := expect.Body.(map[string]interface{}); ok{

		var actualBody map[string]interface{}
		json.Unmarshal(bodyByte, &actualBody)

		// float64, map[string]interface{}, string

		if CheckMap(actualBody, expectBody) {
			return true, ""
		} else {
			return false, "not match"
		}

	} else {
		panic("invalid body type")
	}
}

func CheckMysql(actual []map[string]interface{}, expect []map[string]interface{}) (bool, string) {
	return reflect.DeepEqual(actual, expect), ""
}

func CheckRedis(actual map[string]interface{}, expect map[string]interface{}) (bool, string) {

	// TODO: 比照redis，返回结果和结果字符串

	return true, ""
}

func CheckMap(a map[string]interface{}, b map[string]interface{}) bool {
	var r interface{}
	for key, value := range a {
		if _, ok := value.(string); ok {
			// TODO: 替代全局变量

		}
		if valueMap, ok := value.(map[string]interface{}); ok {
			if bMap, ok := b[key].(map[string]interface{}); ok {
				return CheckMap(valueMap, bMap)
			} else {
				return false
			}
		}
		return r == value
	}
}