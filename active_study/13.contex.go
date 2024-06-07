package main

import (
	"context"
	"fmt"
)

type V struct {
	Name string
}

func main() {

	empCtx := context.Background()
	valCtx := context.WithValue(empCtx, "key", &V{"matt"})

	// bad case
	f := func(c context.Context) {
		// here need change context "key" to pass to another
		if v, ok := valCtx.Value("key").(*V); ok {
			v.Name = "new name"
		}
	}
	f(valCtx)

	v := valCtx.Value("key").(*V)
	fmt.Println(v.Name) // new name

	// good case
	f2 := func(c context.Context) context.Context {
		// here need change context "key" to pass to another
		newCtx := context.WithValue(c, "key", "new name")
		return newCtx
	}
	newCtx := f2(valCtx)
	fmt.Println(newCtx.Value("key")) // new name
}
