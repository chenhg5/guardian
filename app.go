package main

import (
	"net/http"
	"sync"
)

type Engine struct {
	// 入口文件
	entrance string

	// 全局变量
	vars Vars

	// 测试表格集
	suits Suits

	// 测试结果
	result map[string]Results

	// 竞争锁
	reslock sync.Mutex

	// 是否是调试模式
	debug bool
}

func New(path string) *Engine {
	return Read(path, new(Engine))
}

var IsDebug bool

func (eng *Engine) Run() {

	IsDebug = eng.debug
	LogFirst()

	// 测试
	var wg sync.WaitGroup
	for key, suit := range eng.suits {
		wg.Add(1)
		if suit.Name != "" {
			LogTitle(suit.Name)
		} else {
			LogTitle(key)
		}
		result := suit.Run()
		go func(eng *Engine, key string, result Results) {
			eng.reslock.Lock()
			eng.result[key] = result
			eng.reslock.Unlock()
			wg.Add(-1)
		}(eng, key, result)
	}

	wg.Wait()

	// 输出最终结果
	LogFinal()
}

func (suit *Suit) Run() Results {

	// 预处理 => 发请求 => 得到响应 => 对比响应 => 对比数据库与redis结果 => 记录结果返回 => 变量更新

	var (
		actual           *http.Response
		actualChan       = make(chan *http.Response, 0)
		resultList       = make([]Result, 0)
		pass             = true
		resPass          = true
		dataPass         = true
		resDesc          = ""
		sqlDesc          = ""
		finalDesc        = ""
		checkResOk       bool
		checkResResult   string
		checkMysqlOk     bool
		checkMysqlResult string
		url              string
	)

	// 预处理
	Excecute((*suit).PreSqls)

	for _, table := range (*suit).Tables {

		resDesc = ""
		sqlDesc = ""
		resPass = true
		dataPass = true

		// 发请求
		url = GlobalVars.ReplaceString(table.Request.Url)

		if table.Concurrent == 0 {
			actual = GetRequester(table.Request.Method)(url, table.Request.Param, table.Request.Header)
		} else {
			for i := 0; i < table.Concurrent; i++ {
				go func() {
					actualChan <- GetRequester(table.Request.Method)(url, table.Request.Param, table.Request.Header)
				}()
			}
			count := 0
			for count < table.Concurrent {
				actual = <-actualChan
				count++
			}
		}

		if actual == nil {
			pass = false
			continue
		}

		// 对比响应
		checkResOk, checkResResult = CheckResponse(actual, table.Response)

		if !checkResOk {
			pass = false
			resPass = false
		}

		resDesc += checkResResult

		// 对比数据
		for _, data := range table.Data {
			checkMysqlOk, checkMysqlResult = CheckMysql(Query(data.Sql), GlobalVars.ReplaceMap(data.Result))
			if !checkMysqlOk {
				pass = false
				dataPass = false
			}
			sqlDesc += checkMysqlResult
		}

		// 记录结果
		resultList = append(resultList, Result{
			ResPass:  resPass,
			DataPass: dataPass,
			ResDesc:  resDesc,
			SqlDesc:  sqlDesc,
			Title:    table.Info.Title,
		})

		// 变量更新
		GlobalVars.Refresh(table.After)

		LogResult(dataPass, resPass, IsDebug, sqlDesc, resDesc, table.Info.Title)
	}

	// 后续处理
	Excecute((*suit).AfterSqls)

	return Results{
		Pass:        pass,
		Description: finalDesc,
		List:        resultList,
	}
}

// 入口配置
type Config struct {
	Tables   map[string]SuitConfig
	Database Database
	Redis    Redis
	Vars     map[string]interface{}
	Debug    bool
}

type SuitConfig struct {
	Tables    []string `json:"list"`
	Name      string   `json:"name"`
	PreSqls   ExceSql  `json:"pre-execution"`
	AfterSqls ExceSql  `json:"after-execution"`
}

type Results struct {
	List        []Result
	Description string
	Title       string
	Pass        bool
}

type Result struct {
	ResPass  bool
	DataPass bool
	ResDesc  string
	SqlDesc  string
	Title    string
}

type Suits map[string]*Suit

func (su Suits) Add(key string, suit *Suit) {
	su[key] = suit
}

// 测试集
type Suit struct {
	Tables    []Table
	Name      string
	PreSqls   ExceSql
	AfterSqls ExceSql
}

// 测试表格
type Table struct {
	Info       Info
	Concurrent int
	Request    TableRequest
	Response   TableResponse
	Data       []TableData
	After      Vars
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
	Header     map[string]string
	Body       interface{}
	StatusCode int
}

// 数据信息
type TableData struct {
	Sql    string
	Result []map[string]interface{}
}

type ExceSql []string
