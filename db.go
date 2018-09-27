package main

import (
	"github.com/chenhg5/go-utils/database"
)

type Database struct {
	Port     string
	User     string
	Password string
	Charset  string
	Host     string
	Database string
}

type Redis struct {
	Database string
	User     string
	Password string
	Host     string
}

var db *database.SqlDBStruct

func InitDatabase(config Database) {
	db = database.InitDefaultDB(database.Config{
		UserName:     config.User,
		Password:     config.Password,
		Port:         config.Port,
		Ip:           config.Host,
		DatabaseName: config.Database,
		Charset:      config.Charset,
		MaxIdleConns: 20,
		MaxOpenConns: 50,
	})
}

func Query(sql string) []map[string]interface{} {
	sql = GlobalVars.ReplaceSql(sql)
	res, _ := db.Query(sql)
	return res
}
