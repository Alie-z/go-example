package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func main() {

	startTime := time.Now()
	var v string
	for i := 0; i < 10000; i++ {
		v = fmt.Sprintf("%d", i)
	}
	fmt.Println(time.Now().Sub(startTime), v, reflect.TypeOf(v)) // 2.5018ms

	startTime = time.Now()
	for i := 0; i < 10000; i++ {
		v = strconv.Itoa(i)
	}
	fmt.Println(time.Now().Sub(startTime), v) // 700.6µs
}
