package guardian

import "github.com/chenhg5/go-utils/database"

type Database struct {
	Port     string
	User     string
	Password string
	Charset  string
	Host     string
	Table    string
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
		UserName     : config.User,
		Password     : config.Password,
		Port         : config.Port,
		Ip           : config.Host,
		DatabaseName : config.Port,
		Charset      : config.Charset,
		MaxIdleConns : 20,
		MaxOpenConns : 50,
	})
}

func Query(sql string) []map[string]interface{} {
	res, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	return res
}