package main

import "fmt"

// Address 地址结构体
type Address struct {
	Province   string
	City       string
	CreateTime string
}

// Email 邮箱结构体
type Email struct {
	Account    string
	CreateTime string
}

// Person 结构体
type Person struct {
	Name    string
	Age     int
	Address // 匿名字段，嵌套结构体
	Email   // 匿名字段，嵌套结构体
}

func main() {
	// p1 := new(Person)
	p1 := &Person{}
	p1.Name = "barry"
	p1.Age = 26
	//p1.CreateTime = "2019" // ambiguous selector p8.CreateTime，产生字段歧义
	p1.Address.CreateTime = "2021" // 指定Address结构体中的CreateTime
	p1.Email.CreateTime = "20211"  // 指定Email结构体中的CreateTime
	fmt.Println(*p1)
}
