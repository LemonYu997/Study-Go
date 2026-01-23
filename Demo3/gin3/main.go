package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

// 静态文件 html页面上用到的样式文件.css js文件图片
func main() {
	r := gin.Default()
	// 加载静态文件
	r.Static("/xxx", "./statics")
	// gin中给模版添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	// 模版解析
	//r.LoadHTMLFiles("template/index.tmpl")
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		// HTTP请求
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			// 模版渲染
			"title": "lemon",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		// HTTP请求
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			// 模版渲染
			"title": "<a href='https://www.baidu.com'>百度</a>",
		})
	})

	// 默认端口 8080 启动服务器
	// 监听 0.0.0.0:8080（Windows 下为 localhost:8080）
	r.Run()
}
