package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	inspect(p)
}

func inspect(x interface{}) {
	// 获取接口值的反射对象
	v := reflect.ValueOf(x)
	fmt.Println(">>> v", v.Kind())
	// 如果接口值不是结构体或者指向结构体的指针，则直接返回
	if v.Kind() != reflect.Struct {
		fmt.Println("类型不是结构体")
		return
	}

	// 遍历结构体的字段
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fmt.Printf("字段 %d: 名称=%s, 类型=%s, 值=%v\n", i+1, v.Type().Field(i).Name, field.Type(), field.Interface())
	}
}
