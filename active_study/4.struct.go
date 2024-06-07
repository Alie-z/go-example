package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	var p *Person = new(Person)
	fmt.Println(*p)
	// (*p).Name = "smith" 等价  p.Name = "smith"
	// go设计者为了方便，底层对 p.Name = "smith" 进行处理，等价于(*p).Name = "smith"
	p.Name = "john"
	(*p).Age = 30
	fmt.Println(*p)

	p1 := Person{
		Name:  "smith",
		Age:   20,
		Email: "<EMAIL>",
	}
	p1.Email = "123@.com"
	fmt.Println(p1)

	nick := &Person{
		Name:  "Tom",
		Age:   27,
		Email: "tjtulong@xxx.com",
	}
	fmt.Println(*nick)
}
