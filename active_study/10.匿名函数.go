// package main

// import "fmt"

// func getAction() func() {
// 	return func() {
// 		fmt.Println("Hello World")
// 	}
// }

// func main() {
// 	action := getAction()
// 	action()
// }

package main

import "fmt"

func do(action func()) {
	action()
}

func main() {
	sayHello := func() {
		fmt.Println("Hello World")
	}
	do(sayHello)
}
