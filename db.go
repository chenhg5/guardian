package guardian

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

func InitDatabase(database Database) {
	// TODO: 初始化数据库
}

func Query(sql string) []map[string]interface{} {
	// TODO: sql执行
	return []map[string]interface{}{}
}