
 # 🚀 Go 语言中的 gin 框架 💪
Gin 是一个用 Go (Golang) 编写的 HTTP web 框架。

官方文档：
https://gin-gonic.com/zh-cn/docs/

如果还没有安装Go语言开发环境，参考一下我的这篇文章：

[第 N 次Hello World ! Ready ？Go !](https://juejin.cn/post/7165669292522864676)

# 新建项目初始化，以及安装

## 新建文件夹后，初始化 go.mod 文件
```
//初始化
go mod init gin
```

## 下载安装Gin
```
//下载
go get -u github.com/gin-gonic/gin
```

## tips：踩坑指南
`go get 命令下载gin包时出现报错`

go在1.13版本后，默认开启了：
GOSUMDB=sum.golang.org
而这个网址sum.golang.org 在国内是无法访问，故需要关闭

## 解决方案：
go env -w GOSUMDB=off

## 更换国内源
go env -w GO111MODULE=on
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

首先，使用Gin输出一个 Ready?Go!

## 新建文件
在文件夹中新建 `main.go` 文件

## 导入Gin
```
import "github.com/gin-gonic/gin"
```

如果使用诸如 `http.StatusOK` 之类的常量，则需要引入 net/http 包
```
import "net/http"
```

## 官网的demo
```
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
```

## 自己的代码
简单分析一下，每一步做了什么
```
package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

func main() {
   // 创建服务，Default返回一个默认的路由引擎
   ginServer := gin.Default()
   // GET 方法, 添加路由地址和回调函数
   // Gin进行了封装，把request和response都封装到了gin.Context的上下文环境中。
   ginServer.GET("/", func(context *gin.Context) {
      context.String(http.StatusOK, "Ready?Go!")
   })
   // 服务器端口 如果不传端口号，默认8080
   ginServer.Run(":8082")
}
```


运行项目，通过命令启动，或者直接在IDE开发工具中启动。
```
go run main.go
```


浏览器访问地址：
http://localhost:8082/

成功输出：Ready?Go!

看一下上面的代码，回调函数中的 gin.Context 
除了返回 Json 和 String 当然还可以返回 HTML页面。

## 响应一个前端页面
```
 // 响应一个前端页面
ginServer.GET("/index", func(context *gin.Context) {
  context.HTML(http.StatusOK, "index.html", gin.H{"title": "gin web页面"})
})
```

## 使用 ico 文件
安装命令
```
go get "github.com/thinkerou/favicon"
```
```
// 使用 favicon.ico 文件
ginServer.Use(favicon.New("./favicon.ico"))
```

## 加载静态页面
```
// 加载静态页面
ginServer.LoadHTMLGlob("templates/*")
```

## 加载资源文件
```
// 加载资源文件
ginServer.Static("/static", "./static")
```

## Restful API
Gin 是一个标准的 Web 服务框架
遵循 Restful API 接口规范
```
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
```

## 参数处理
Query接收参数，通过 gin.Context 的 Query 函数获取参数。
```
// 接收前端传递过来的参数
ginServer.GET("/user/info", func(context *gin.Context) {
   userid := context.Query("userid")
   username := context.Query("username")
   context.JSON(http.StatusOK, gin.H{
      "userid":   userid,
      "username": username,
   })
})
```
path 参数，可通过 Param 函数获取请求 Path 中的参数。
Path 路径中参数以:开头。
```
// userinfo/info/123/orzr3
ginServer.GET("/user/info/:userid/:username", func(context *gin.Context) {
   userid := context.Param("userid")
   username := context.Param("username")
   context.JSON(http.StatusOK, gin.H{
      "userid":   userid,
      "username": username,
   })
})
```

## 传递 json 数据
```
// 传递 json 数据
ginServer.POST("/json", func(context *gin.Context) {
   // request.body
   data, _ := context.GetRawData()
   var m map[string]interface{}
   // 包装成json 数据
   _ = json.Unmarshal(data, &m)
   context.JSON(http.StatusOK, m)
})
```

## 表单提交示例

web页面代码
```
<form action="/user/add" method="post">
   <p>username: <input type="text" name="username" /></p>
   <p>password: <input type="text" name="password" /></p>
   <button type="submit">提交</button>
 </form>
```

go代码
```
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
```


点击提交后，触发action请求： /user/add

{"msg":"ok","password":"222","username":"11"}

## 路由重定向
```
// 路由重定向
ginServer.GET("/redirect", func(context *gin.Context) {
   // 重定向 301 http.StatusMovedPermanently
   context.Redirect(http.StatusMovedPermanently, "https://www.bilibili.com")
})
```

## 404 NoRoute
```
// 404 NoRoute
ginServer.NoRoute(func(context *gin.Context) {
   // context.JSON(http.StatusNotFound, gin.H{
   //     "code": 404,
   //     "msg":  "404 not found",
   // })
   context.HTML(http.StatusNotFound, "404.html", nil)
})
```

## 路由组
```
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
```

## 中间件
```
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
```
```
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
```