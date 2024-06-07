package main

import "fmt"

func main() {
	appendAnyF := func(t []any, toAppend ...any) []any {
		ret := append(t, toAppend...)
		return ret
	}

	emptySlice := []any{}
	slice2 := []any{"hello", "world"}

	// bug append slice as a element
	emptySlice = appendAnyF(emptySlice, slice2)
	fmt.Println(emptySlice) // only 1 element [[hello world]]

	// emptySlice = []any{}
	emptySlice = appendAnyF(emptySlice, slice2...)
	fmt.Println(emptySlice) // [hello world]

	arr := [3]int{0, 1, 2}
	f := func(v [3]int) {
		v[0] = 100
	}
	f(arr)           // no modify to arr
	fmt.Println(arr) // [0 1 2]

	// arr := [3]int{0, 1, 2}
	// f := func(v *[3]int) {
	// 	v[0] = 100
	// }
	// f(&arr)
	// fmt.Println(arr)

	// arr := []int{0, 1, 2}
	// f := func(v []int) {
	// 	v[0] = 100 // can modify origin array
	// 	// v = append(v, 4) // new memory allocated
	// 	// v[0] = 50        // no modify to origin array
	// }
	// f(arr)
	// fmt.Println(arr) // [100 1 2]
	// mp := make(map[string]int) // 使用make函数初始化map
	// mp["dd"] = 1
	// fmt.Println(mp) // 输出: map[dd:1]

	var tmp = make(map[string]int)
	mp := &tmp
	(*mp)["dd"] = 1 // 注意这里需要先解引用指针
	fmt.Println(">>> tmp", tmp["dd"], (*mp)["dd"])
}
