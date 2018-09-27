package guardian

import "net/http"

type Engine struct {
	// 入口文件
	entrance string

	// 全局变量
	vars Vars

	// 测试表格
	tables Suits

	// 测试结果
	result map[string]Results
}

func New(path string) *Engine {
	return Read(path, new(Engine))
}

func (eng *Engine) Run() {
	// 测试
	for key, suit := range eng.tables {
		go func(eng *Engine, key string) {
			eng.result[key] = suit.Run()
		}(eng, key)
	}

	// 输出结果
	Output(eng.result)
}

func (suit *Suit) Run() Results {

	// TODO: 运行集合测试
	// 发请求 =》得到响应 =》对比响应 =》对比数据库与redis结果 =》记录结果返回

	var (
		actual           *http.Response
		pass             = true
		description      = ""
		checkResOk       bool
		checkResResult   string
		checkMysqlOk     bool
		checkMysqlResult string
		resultList       = make([]Result, 0)
	)
	for _, table := range *suit {

		// 请求
		actual = GetRequester(table.Request.Method)(table.Request.Url, table.Request.Param, table.Request.Header)

		if actual == nil {
			pass = false
			continue
		}

		// 对比响应
		checkResOk, checkResResult = CheckResponse(actual, table.Response)

		if !checkResOk {
			pass = false
			description += checkResResult
		}

		// 对比数据
		for _, data := range table.Data {
			checkMysqlOk, checkMysqlResult = CheckMysql(Query(data.Sql), data.Result)
			if !checkMysqlOk {
				pass = false
				description += checkMysqlResult
			}
		}

		resultList = append(resultList, Result{
			Pass: pass,
			Description: description,
		})

		// TODO: 设置全局变量
	}

	return Results{
		Pass:        pass,
		Description: description,
		List:        resultList,
	}
}

// 入口配置
type Config struct {
	Tables   map[string][]string
	Database Database
	Redis    Redis
	Vars     map[string]string
}

type Results struct {
	List        []Result
	Description string
	Title       string
	Pass        bool
}

type Result struct {
	Pass        bool
	Description string
	Title       string
}

type Suits map[string]*Suit

func (su Suits) Add(key string, suit *Suit) {
	su[key] = suit
}

// 测试集
type Suit []Table

// 测试表格
type Table struct {
	Info     Info
	Request  TableRequest
	Response TableResponse
	Data     []TableData
}

// 表格信息
type Info struct {
	Title       string
	Description string
}

// 请求信息
type TableRequest struct {
	Url    string
	Method string
	Param  map[string]interface{}
	Header map[string]string
}

// 响应信息
type TableResponse struct {
	Header   map[string]string
	Body     interface{}
	BodyType string
}

// 数据信息
type TableData struct {
	Sql    string
	Result []map[string]interface{}
}
