package main

import (
	"demo/myTools"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

// 自定义Go中间件 拦截器
func MyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Do something before request
		// 通过自定义的中间件，设置的值，在后续处理只要调用了这个中间件，就可以获取到这个值
		c.Set("usersession", "orzr3-123456")
		c.Next() // 一定要调用该方法，否则后续的处理不会被执行
		// c.Abort() // 一旦调用了该方法，后续的处理都不会被执行
		// Do something after request
	}
}

func main() {
	// 创建服务
	ginServer := gin.Default()
	// 使用 favicon.ico 文件
	ginServer.Use(favicon.New("./favicon.ico"))

	// 加载静态页面
	ginServer.LoadHTMLGlob("templates/*")
	// 加载资源文件
	ginServer.Static("/static", "./static")

	// 或者 LoadHTMLFiles
	// ginServer.LoadHTMLFiles("templates/index.html")
	// 访问地址，处理请求
	ginServer.GET("/hello", func(context *gin.Context) {
		sum := myTools.Add(3, 2)
		// sum = 1
		context.JSON(200, gin.H{"message": sum})
	})
	// Gin RESTful API
	ginServer.POST("/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "post user"})
	})

	ginServer.PUT("/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "put user"})
	})

	ginServer.DELETE("/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "delete user"})
	})

	// 响应一个前端页面
	ginServer.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{"title": "gin web页面"})
	})

	// 接收前端传递过来的参数
	ginServer.GET("/user/info", func(context *gin.Context) {
		userid := context.Query("userid")
		username := context.Query("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})

	// 接收前端传递过来的参数，并且使用中间件
	ginServer.GET("/user/info/handler", MyMiddleware(), func(context *gin.Context) {
		// 获取中间件中设置的值
		userSession := context.MustGet("usersession").(string)
		// 打印
		log.Println("userSession:================>", userSession)

		userid := context.Query("userid")
		username := context.Query("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":      userid,
			"username":    username,
			"userSession": userSession,
		})
	})

	// userinfo/info/123/orzr3
	ginServer.GET("/user/info/:userid/:username", func(context *gin.Context) {
		userid := context.Param("userid")
		username := context.Param("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})

	// 前端给后端传递 json 数据
	ginServer.POST("/json", func(context *gin.Context) {
		// request.body
		data, _ := context.GetRawData()
		var m map[string]interface{}
		// 包装成json 数据
		_ = json.Unmarshal(data, &m)
		context.JSON(http.StatusOK, m)
	})

	// user/add 添加数据
	// 支持函数式编程 java不能把函数作为参数传递，go可以
	ginServer.POST("/user/add", func(context *gin.Context) {
		// 获取前端传递过来的参数
		username := context.PostForm("username")
		password := context.PostForm("password")
		if username == "" || password == "" {
			context.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "参数错误",
			})
			return
		}
		// 返回给前端
		context.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"password": password,
		})
	})

	// 路由重定向
	ginServer.GET("/redirect", func(context *gin.Context) {
		// 重定向 301	http.StatusMovedPermanently
		context.Redirect(http.StatusMovedPermanently, "https://www.bilibili.com")
	})

	// 404 NoRoute
	ginServer.NoRoute(func(context *gin.Context) {
		// context.JSON(http.StatusNotFound, gin.H{
		// 	"code": 404,
		// 	"msg":  "404 not found",
		// })
		context.HTML(http.StatusNotFound, "404.html", nil)
	})

	// 路由组 /system/add 路由组的写法
	systemGroup := ginServer.Group("/system")
	{
		systemGroup.GET("/add", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "system add",
			})

		})
		systemGroup.DELETE("/delete", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "system delete",
			})
		})
	}

	// 服务器端口
	ginServer.Run(":8082")
}

/*
 tips:
 为什么需要一个新的语言
 需要一种高效的执行速度，编译速度和开发速度的编程语言

 Docker 就是使用Go进行开发的

 Go编译的程序，可以媲美 C / C++ 的速度，而且更加安全
 它是21世纪的C语言
*/
