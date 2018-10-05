package main

import (
	"net/http"
	"encoding/json"
	"regexp"
	"strings"
	"io"
	"io/ioutil"
	"compress/gzip"
)

func CheckResponse(actual *http.Response, expect TableResponse) (bool, string) {

	var (
		reader   io.ReadCloser
		err      error
		bodyByte []byte
	)
	if actual.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(actual.Body)
		if err != nil {
			panic(err)
		}
	} else {
		reader = actual.Body
	}

	defer actual.Body.Close()

	bodyByte, err = ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	if expectBody, ok := expect.Body.(string); ok {
		if string(bodyByte) == expectBody {
			return true, ""
		} else {
			return false, string(bodyByte) + "\n\n" + expectBody
		}
	} else if expectBody, ok := expect.Body.(map[string]interface{}); ok {

		var actualBody map[string]interface{}
		json.Unmarshal(bodyByte, &actualBody)

		// float64, map[string]interface{}, string

		if CheckMap(actualBody, expectBody) {
			return true, ""
		} else {
			expectByte, _ := json.Marshal(expectBody)
			return false, string(bodyByte) + "\n\n" +  string(expectByte)
		}

	} else {
		panic("invalid body type")
	}
}

func CheckMysql(actual []map[string]interface{}, expect []map[string]interface{}) (bool, string) {
	if len(actual) != len(expect) || len(actual) == 0 {
		actualByte, _ := json.Marshal(actual)
		expectByte, _ := json.Marshal(expect)
		return false, string(expectByte) + "\n\n" +  string(actualByte)
	}

	for i := 0; i < len(actual); i++ {
		if !CheckMapIgnoreInt64(actual[i], expect[i]) {
			actualByte, _ := json.Marshal(actual)
			expectByte, _ := json.Marshal(expect)
			return false, string(expectByte) + "\n\n" +  string(actualByte)
		}
	}

	return true, ""
}

func CheckRedis(actual map[string]interface{}, expect map[string]interface{}) (bool, string) {

	// TODO: 比照redis，返回结果和结果字符串

	return true, ""
}

func CheckMap(a map[string]interface{}, e map[string]interface{}) bool {

	var (
		find string
		reg  = regexp.MustCompile("{{(.*?)}}")
	)

	if len(a) == 0 {
		return false
	}

	for key, value := range a {
		if _, ok := value.(string); ok {
			if _, ok := e[key].(string); ok {
				find = reg.FindString(e[key].(string))
				if find != "" {
					find = strings.Replace(find, "{{", "", -1)
					find = strings.Replace(find, "}}", "", -1)
					GlobalVars.Add(find, value.(string))
					continue
				}
			} else {
				return false
			}
		}
		if valueMap, ok := value.(map[string]interface{}); ok {
			if bMap, ok := e[key].(map[string]interface{}); ok {
				return CheckMap(valueMap, bMap)
			} else {
				if eStr, ok := e[key].(string); ok && eStr == "*" {
					continue
				} else {
					return false
				}
			}
		}
		if eStr, ok := e[key].(string); ok && eStr == "*" {
			continue
		}
		if e[key] != value {
			return false
		}
	}
	return true
}

func CheckMapIgnoreInt64(a map[string]interface{}, e map[string]interface{}) bool {
	if len(a) == 0 {
		return false
	}

	for key, value := range a {
		if intValue, ok := value.(int64); ok {
			if eFloat, ok2 := e[key].(float64); ok2 {
				if eFloat != float64(intValue) {
					return false
				} else {
					continue
				}
			}
		}
		if FloatValue, ok := value.(float64); ok {
			if eInt, ok2 := e[key].(int64); ok2 {
				if float64(eInt) != FloatValue {
					return false
				} else {
					continue
				}
			}
		}
		if valueMap, ok := value.(map[string]interface{}); ok {
			if bMap, ok := e[key].(map[string]interface{}); ok {
				return CheckMapIgnoreInt64(valueMap, bMap)
			} else {
				return false
			}
		}
		if e[key] != value {
			return false
		}
	}
	return true
}
