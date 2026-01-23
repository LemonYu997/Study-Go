package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建带默认中间件（日志与恢复）的 Gin 路由器
	r := gin.Default()

	// 定义简单的 GET 路由
	r.GET("/json", func(c *gin.Context) {
		// 方法1：使用map
		//data := map[string]interface{}{
		//	"name":    "小王子",
		//	"message": "hello world",
		//	"age":     18,
		//}
		// 可以直接使用 gin 生成map
		data := gin.H{"name": "小王子", "message": "hello world!", "age": 18}

		// 返回 JSON 响应
		c.JSON(http.StatusOK, data)
	})

	// 方法2：结构体
	type msg struct {
		// 字段名必须大写开头，可以使用tag来定制化结构体字段 json序列化时生效
		Name    string `json:"name"`
		Message string
		Age     int
	}
	r.GET("/json2", func(c *gin.Context) {
		data := msg{
			"小王子",
			"Hello World!",
			18}
		// json 的序列化
		c.JSON(http.StatusOK, data)
	})

	// 默认端口 8080 启动服务器
	// 监听 0.0.0.0:8080（Windows 下为 localhost:8080）
	r.Run()
}
