// package main

// import (
// 	"fmt"
// )

// type data int

// func (d1 data) add(d2 data) data {
// 	return d1 + d2
// }

// func main() {
// 	d1 := data(1)
// 	d2 := data(1)
// 	result := d1.add(d2.add(d1))
// 	fmt.Println(result)
// }

// package main

// import (
// 	"fmt"
// )

// type People struct {
// 	Age    int
// 	Gender string
// 	Name   string
// }

// func (p People) printPeople() {
// 	fmt.Println("People ", p.Name, ": ", p.Age, "岁, 性别:", p.Gender)
// }

// type Student struct {
// 	People // 嵌入匿名People结构体
// 	Grade  string
// }

// func (s Student) printStudent(p People) {
// 	fmt.Println("People ", p.Name, ": ", p.Age, "岁, 性别:", p.Gender, " 年级：", s.Grade)
// }

// func main() {
// 	p := People{10, "男", "小明"}
// 	s := Student{p, "三年级"}
// 	s.printPeople()                      // People  小明 :  10 岁, 性别: 男
// 	fmt.Println("Student's age:", s.Age) // Student's age: 10

// 	s.printStudent(p)
// }

package main

import (
	"fmt"
)

type People struct {
	Name string
}

func (p People) setName(name string) {
	p.Name = name
}

func (p *People) setName2(name string) {
	p.Name = name
}

func main() {
	p := People{"xiaoxiao"}
	p.setName("xiaoming")
	fmt.Println(p.Name) // xiaoxiao

	p.setName2("xiaoliang")
	fmt.Println(p.Name) // xiaoliang

	p1 := &People{"xiaobei"}
	p1.setName("xiaoming") // 编译器转变为(*p1).setName("xiaoming")
	fmt.Println(p1.Name)   // xiaobei

	p1.setName2("xiaoliang")
	fmt.Println(p1.Name) // xiaoliang
}
