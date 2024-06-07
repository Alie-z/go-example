package main

import "fmt"

type Middleware interface {
	Execute()
}

type LoggingMiddleware struct{}

func (m LoggingMiddleware) Execute() {
	fmt.Println("Logging...")
}

type AuthenticationMiddleware struct{}

func (m AuthenticationMiddleware) Execute() {
	fmt.Println("Authenticating...")
}

func main() {
	// 定义一个中间件切片
	middlewares := []Middleware{
		LoggingMiddleware{},
		AuthenticationMiddleware{},
	}

	// 执行中间件
	for _, middleware := range middlewares {
		middleware.Execute()
	}
}
