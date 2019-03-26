package main

import (
	"os"
	"encoding/json"
	"path"
)

func Read(entrancePath string, engine *Engine) *Engine {
	engine.entrance = entrancePath
	engine.suits = make(Suits, 0)
	engine.result = make(map[string]Results, 0)

	// 入口文件读取配置
	entrance, err := os.Open(entrancePath)
	defer entrance.Close()
	if err != nil {
		panic(err)
	}
	fileinfo, err := entrance.Stat()
	if err != nil {
		panic(err)
	}
	fileSize := fileinfo.Size()
	configByte := make([]byte, fileSize)
	_, err = entrance.Read(configByte)
	if err != nil {
		panic(err)
	}

	var config Config
	if err := json.Unmarshal(configByte, &config); err != nil {
		panic(err)
	}

	dir := path.Dir(entrancePath)

	// 初始化数据库
	InitDatabase(config.Database)

	// 读取表格信息
	for suit, tables := range config.Tables {
		var tableList = make([]Table, 0)
		for i := 0; i < len(tables.Tables); i++ {
			tableFile, err := os.Open(dir + "/" + tables.Tables[i])
			if err != nil {
				tableFile.Close()
				panic(err)
			}
			fileinfo, err := tableFile.Stat()
			if err != nil {
				panic(err)
			}
			fileSize := fileinfo.Size()
			tableByte := make([]byte, fileSize)
			_, err = tableFile.Read(tableByte)
			if err != nil {
				tableFile.Close()
				panic(err)
			}
			tableFile.Close()
			var table Table
			json.Unmarshal(tableByte, &table)
			tableList = append(tableList, table)
		}
		engine.suits.Add(suit, &Suit{
			Tables:    tableList,
			Name:      tables.Name,
			PreSqls:   tables.PreSqls,
			AfterSqls: tables.AfterSqls,
		})
	}

	// 设置全局变量
	engine.vars = config.Vars
	GlobalVars = config.Vars
	engine.debug = config.Debug

	return engine
}
