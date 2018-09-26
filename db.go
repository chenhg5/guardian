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

func Exec(sql string) []map[string]interface{} {
	return []map[string]interface{}{}
}