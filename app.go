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

	// 测试表格
	tables Suits

	// 测试结果
	result map[string]Results

	// 竞争锁
	reslock sync.Mutex
}

func New(path string) *Engine {
	return Read(path, new(Engine))
}

func (eng *Engine) Run() {
	// 测试
	var wg sync.WaitGroup
	for key, suit := range eng.tables {
		wg.Add(1)
		result := suit.Run()
		go func(eng *Engine, key string, result Results) {
			eng.reslock.Lock()
			eng.result[key] = result
			eng.reslock.Unlock()
			wg.Add(-1)
		}(eng, key, result)
	}

	wg.Wait()

	// 输出结果
	Output(eng.result)
}

func (suit *Suit) Run() Results {

	// 发请求 =》得到响应 =》对比响应 =》对比数据库与redis结果 =》记录结果返回

	var (
		actual           *http.Response
		actualChan       chan *http.Response
		pass             = true
		resPass          = true
		dataPass         = true
		resDesc          = ""
		finalDesc        = ""
		checkResOk       bool
		checkResResult   string
		checkMysqlOk     bool
		checkMysqlResult string
		resultList       = make([]Result, 0)
		url              string
	)
	for _, table := range *suit {

		resDesc = ""

		// 请求

		url = GlobalVars.Replace(table.Request.Url)

		if table.Concurrent == 0 {
			go func() {
				actualChan <- GetRequester(table.Request.Method)(url, table.Request.Param, table.Request.Header)
			}()
			count := 0
			for count < table.Concurrent + 1 {
				actual = <-actualChan
				count++
			}
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
			resDesc += checkResResult
		}

		// 对比数据
		for _, data := range table.Data {
			checkMysqlOk, checkMysqlResult = CheckMysql(Query(data.Sql), data.Result)
			if !checkMysqlOk {
				pass = false
				dataPass = false
				resDesc += checkMysqlResult
			}
		}

		resultList = append(resultList, Result{
			ResPass:     resPass,
			DataPass:    dataPass,
			Description: resDesc,
			Title:       table.Info.Title,
		})
	}

	return Results{
		Pass:        pass,
		Description: finalDesc,
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
	ResPass     bool
	DataPass    bool
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
	Info       Info
	Concurrent int
	Request    TableRequest
	Response   TableResponse
	Data       []TableData
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
	Header map[string]string
	Body   interface{}
}

// 数据信息
type TableData struct {
	Sql    string
	Result []map[string]interface{}
}
