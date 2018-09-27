package main

import "strings"

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