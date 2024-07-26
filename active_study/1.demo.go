package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type Data struct {
	RawMessage json.RawMessage `json:"data"`
}

func main() {
	allPage := float64(666) / float64(200)
	allPage = math.Ceil(allPage)
	fmt.Println(">>> n", allPage)
	for i := 2; i <= int(allPage); i++ {
		fmt.Println(">>> i", i)
	}
}
