package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// querystring传参
func main() {
	// 创建带默认中间件（日志与恢复）的 Gin 路由器
	r := gin.Default()

	// 定义简单的 GET 路由
	// 测试地址：http://127.0.0.1:8080/web?query=AAA&age=18
	r.GET("/web", func(c *gin.Context) {
		// 获取浏览器发请求携带的 query string 参数
		//name := c.Query("query")
		age := c.Query("age")
		// 获取query参数，没有时使用默认值
		//name := c.DefaultQuery("query", "someone")
		// 带结果的获取方法
		name, ok := c.GetQuery("query")
		if !ok {
			// 取不到
			name = "someone"
		}

		// 返回 JSON 响应
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	// 默认端口 8080 启动服务器
	// 监听 0.0.0.0:8080（Windows 下为 localhost:8080）
	r.Run()
}
