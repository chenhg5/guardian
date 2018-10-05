package main

import (
	"strings"
	"encoding/json"
)

type Vars map[string]string

func (v Vars) Get(key string) string {
	return v[key]
}

func (v Vars) Add(key string, value string) {
	v[key] = value
}

var GlobalVars Vars

func (v Vars) Replace(value string) string {
	for key, val := range v {
		value = strings.Replace(value, "{{" + key + "}}",  val, -1)
	}
	return value
}

func (v Vars) ReplaceSql(value string) string {
	for key, val := range v {
		value = strings.Replace(value, "{{" + key + "}}",  `"` + val + `"`, -1)
	}
	return value
}

func (v Vars) ReplaceMap(mapValue []map[string]interface{}) []map[string]interface{} {

	valueByte, _ := json.Marshal(mapValue)
	value := string(valueByte)

	for key, val := range v {
		value = strings.Replace(value, "{{" + key + "}}",  val, -1)
		value = strings.Replace(value, "{{" + key + "|string}}",  val, -1)
		value = strings.Replace(value, `"{{` + key + `|number}}"`,  val, -1)
	}

	var data []map[string]interface{}
	json.Unmarshal([]byte(value), &data)

	return data
}