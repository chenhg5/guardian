package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 这是一个简单的示例，先导入guardian.sql，接着修改 tests/entrance.json
// 中sql配置选项，然后再运行：
//
// go run main.go
// guardian --tests=tests/entrance.json

func main() {
	app := gin.Default()

	app.POST("/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "ok",
		})
		return
	})

	app.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": gin.H{
				"user": gin.H{
					"name":  "Jack",
					"sex":   1,
					"phone": "038-09829010",
					"job":   "engineer",
				},
			},
		})
		return
	})

	app.POST("/order", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": gin.H{
				"orderNum": "PYpaXEQvw4wqJntDRT8PUwXBvyTeccbf",
			},
		})
		return
	})

	app.GET("/order/:order_num", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": gin.H{
				"order": gin.H{
					"amount":     15.45,
					"created_at": "2018-10-03 00:00:00",
					"state":      1,
				},
			},
		})
		return
	})

	app.Run(":1235")
}
