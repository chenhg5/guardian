package main

import (
	"strings"
	"encoding/json"
	"strconv"
)

type Vars map[string]interface{}

func (v Vars) Get(key string) interface{} {
	return v[key]
}

func (v Vars) Add(key string, value interface{}) {
	kv, ok := v[key]
	if !ok {
		v[key] = value
		return
	}
	if _, ok := kv.(string); ok {
		if _, ok2 := value.(string); ok2 {
			v[key] = value
			return
		}
		if val, ok2 := value.(bool); ok2 {
			if val {
				v[key] = "true"
				return
			} else {
				v[key] = "false"
				return
			}
		}
		if val, ok2 := value.(float64); ok2 {
			if int64(val*100000) == int64(val)*100000 {
				v[key] = strconv.FormatInt(int64(val), 10)
				return
			}
			v[key] = strconv.FormatFloat(val, 'g', 5, 64)
			return
		}
	}
	if _, ok := kv.(bool); ok {
		if _, ok2 := value.(bool); ok2 {
			v[key] = value
			return
		}
		if val, ok2 := value.(string); ok2 {
			if val == "false" {
				v[key] = false
				return
			} else {
				v[key] = true
				return
			}
		}
		if val, ok2 := value.(float64); ok2 {
			if val > 0 {
				v[key] = true
				return
			} else {
				v[key] = false
				return
			}
		}
	}
	if _, ok := kv.(float64); ok {
		if _, ok2 := value.(float64); ok2 {
			v[key] = value
			return
		}
		if val, ok2 := value.(string); ok2 {
			flo, err := strconv.ParseFloat(val, 64)
			if err != nil {
				v[key] = 0
				return
			}
			if int64(flo*100000) == int64(flo)*100000 {
				v[key] = int64(flo)
				return
			} else {
				v[key] = flo
				return
			}
		}
		if val, ok2 := value.(bool); ok2 {
			if val {
				v[key] = 1
				return
			} else {
				v[key] = 0
				return
			}
		}
	}
}

func (v Vars) Refresh(new Vars) {
	for key, value := range new {
		v.Add(key, value)
	}
}

var GlobalVars Vars

func (v Vars) ReplaceString(value string) string {
	var (
		newVal string
	)
	for key, val := range v {
		_, newVal = InterfaceToString(val)
		value = strings.Replace(value, "{{"+key+"}}", newVal, -1)
	}
	return value
}

func (v Vars) Replace(value string) string {
	var (
		isString bool
		newVal   string
	)
	for key, val := range v {
		isString, newVal = InterfaceToString(val)
		if isString {
			value = strings.Replace(value, "{{"+key+"}}", newVal, -1)
		} else {
			value = strings.Replace(value, `"{{`+key+`}}"`, newVal, -1)
			value = strings.Replace(value, `{{`+key+`}}`, newVal, -1)
		}
	}
	return value
}

func (v Vars) ReplaceSql(value string) string {
	var (
		newVal string
	)
	for key, val := range v {
		_, newVal = InterfaceToString(val)
		value = strings.Replace(value, "{{"+key+"}}", `"`+newVal+`"`, -1)
	}
	return value

}

func (v Vars) ReplaceMap(mapValue []map[string]interface{}) []map[string]interface{} {

	valueByte, _ := json.Marshal(mapValue)
	value := string(valueByte)

	var (
		isString bool
		newVal   string
	)
	for key, val := range v {
		isString, newVal = InterfaceToString(val)
		if isString {
			value = strings.Replace(value, "{{"+key+"}}", newVal, -1)
		} else {
			value = strings.Replace(value, `"{{`+key+`}}"`, newVal, -1)
		}
	}

	var data []map[string]interface{}
	json.Unmarshal([]byte(value), &data)

	return data
}

// string/bool/float64
func InterfaceToString(value interface{}) (bool, string) {
	if value, ok := value.(string); ok {
		return true, value
	}
	if value, ok := value.(bool); ok {
		if value {
			return false, "true"
		} else {
			return false, "false"
		}
	}
	if value, ok := value.(float64); ok {
		if int64(value*100000) == int64(value)*100000 {
			return false, strconv.FormatInt(int64(value), 10)
		}
		return false, strconv.FormatFloat(value, 'g', 5, 64)
	}
	return false, ""
}
