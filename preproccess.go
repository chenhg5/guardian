package main

func preExcecute(sqls PreSql) bool {
	for _, sql := range sqls {
		db.Exec(sql)
	}
	return true
}
