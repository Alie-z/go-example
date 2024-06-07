package main

import (
	"fmt"
	"reflect"
)

func main() {
	type S struct {
		A string
		B string
		C string
	}
	typ := reflect.TypeOf(S{})
	funcArr := make([]func() string, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		f := func() string {
			return typ.Field(i).Name
		}
		funcArr[i] = f
	}

	fmt.Println(funcArr[0](), funcArr[1](), funcArr[2]()) // error reflect: Field index out of bounds

	// type S struct {
	// 	A string
	// 	B string
	// 	C string
	// }
	// typ := reflect.TypeOf(S{})
	// fmt.Println(">>> typ", typ, typ.NumField())
	// funcArr := make([]func() string, typ.NumField())
	// for i := 0; i < typ.NumField(); i++ {
	// 	// index := i // assign to a new variable
	// 	f := func() string {
	// 		name := typ.Field(i).Name
	// 		return name
	// 	}
	// 	funcArr[i] = f
	// }

	// fmt.Println(funcArr[0]()) // A

	// go func() {
	// 	defer func() {
	// 		recover() // should do some thing here
	// 	}()

	// 	panic("oh...")
	// }()

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println("bye bye!")

	mp := map[int]int{}
	for i := 0; i < 20; i++ {
		mp[i] = i
	}

	for k, v := range mp {
		fmt.Println(k, v)
	}
	// 按照键的升序顺序输出键值对
	for i := 0; i < len(mp); i++ {
		fmt.Println(i, mp[i])
	}
	// 对键进行排序
	// sort.Ints(keys)

	// // 按照排序后的键顺序输出键值对
	// for _, k := range keys {
	// 	fmt.Println(k, mp[k])
	// }

}
