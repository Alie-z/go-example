// package main

// import "fmt"

// func main() {
// 	fmt.Println(backwards(1))
// 	fmt.Println(backwards(0))
// }

// func backwards(input int) (res int, err error) {

// 	defer func() {
// 		if p := recover(); p != nil {
// 			err = fmt.Errorf("number %d is illegal input", input)
// 		}
// 	}()

// 	return 1 / input, nil
// }

package main

import "fmt"

func main() {
	panicParam()
}

func panicParam() {

	defer func() {
		if p := recover(); p != nil {
			fmt.Println(">>>", p)
		}
	}()

	panic("some message")
}
