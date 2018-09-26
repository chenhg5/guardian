package guardian

import (
	"os"
	"encoding/json"
)

func Read(path string, engine *Engine) *Engine {
	engine.entrance = path

	// 入口文件读取配置
	entrance, err := os.Open(path)
	defer entrance.Close()
	if err != nil {
		panic("")
	}
	configByte := make([]byte, 0)
	_, err = entrance.Read(configByte)
	if err != nil {
		panic("")
	}

	var config Config
	json.Unmarshal(configByte, &config)

	// 初始化数据库
	InitDatabase(config.Database)

	// 读取表格信息
	for suit, tables := range config.Tables {
		var tableList = make(Suit, 0)
		for i := 0; i < len(tables); i++ {
			tableFile, err := os.Open(tables[i])
			if err != nil {
				tableFile.Close()
				panic("")
			}
			tableByte := make([]byte, 0)
			_, err = tableFile.Read(tableByte)
			if err != nil {
				tableFile.Close()
				panic("")
			}
			tableFile.Close()
			var table Table
			json.Unmarshal(tableByte, &table)
			tableList = append(tableList, table)
		}
		engine.tables.Add(suit, &tableList)
	}

	// 读取表格信息
	engine.vars = config.Vars

	return engine
}
