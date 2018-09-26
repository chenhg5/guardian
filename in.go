package guardian

func Read(path string, engine *Engine) *Engine {
	engine.entrance = path

	// TODO: 读取文件

	// 初始化数据库
	InitDatabase()

	// TODO: 初始化测试表格集合

	// TODO: 初始化测试全局变量

	return engine
}
