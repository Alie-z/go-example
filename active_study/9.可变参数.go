package main

import "fmt"

func main() {
	fruits := []string{"apple", "pear", "orange", "watermelon"}
	res4 := joinStrings("||", fruits...)
	fmt.Println(">>> res4", res4)

	res3 := joinStrings("|", "anna", "lee", "chen")
	fmt.Println(">>> res3", res3)
}

// joinStrings joins all the input tokens using the specified sep
func joinStrings(sep string, tokens ...string) string {
	fmt.Printf("Type of args is : %T", tokens) // output：Type of args is : []string
	fmt.Println()

	r := ""
	for i, token := range tokens {
		if i != 0 {
			r += sep
		}
		r += token
	}
	return r
}
