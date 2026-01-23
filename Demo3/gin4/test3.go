package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// form表单传参
func main() {
	// 创建带默认中间件（日志与恢复）的 Gin 路由器
	r := gin.Default()

	r.LoadHTMLFiles("./statics/login.html", "./statics/home.tmpl")
	// login get
	r.GET("/login", func(c *gin.Context) {
		// 返回 HTML 页面
		c.HTML(http.StatusOK, "login.html", nil)
	})
	// login post
	r.POST("/login", func(c *gin.Context) {
		// 获取 form 表单提交的数据
		//username := c.PostForm("username")
		//password := c.PostForm("password")

		// 表单不填时实际返回的是空字符串，并不是 nil
		//username := c.DefaultPostForm("username", "admin")
		//password := c.DefaultPostForm("password", "123456")
		//xxx := c.DefaultPostForm("xxx", "默认值")

		// 带结果的方式
		username, ok := c.GetPostForm("username")
		if !ok {
			username = "admin"
		}
		password, ok := c.GetPostForm("password")
		if !ok {
			password = "123456"
		}
		xxx, ok := c.GetPostForm("xxx")
		if !ok {
			xxx = "默认值"
		}
		c.HTML(http.StatusOK, "home.tmpl", gin.H{
			"Name":     username,
			"Password": password,
			"xxx":      xxx,
		})
	})

	// 默认端口 8080 启动服务器
	// 监听 0.0.0.0:8080（Windows 下为 localhost:8080）
	r.Run()
}
