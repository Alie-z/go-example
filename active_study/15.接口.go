package main

import (
	"fmt"
)

// 定义一个接口
type Switch interface {
	TurnOn()
	TurnOff()
}

// 灯实现了开关接口
type Light struct {
	Name string
}

func (l Light) TurnOn() {
	fmt.Printf("%s 灯已经打开\n", l.Name)
}

func (l Light) TurnOff() {
	fmt.Printf("%s 灯已经关闭\n", l.Name)
}

// 风扇也实现了开关接口
type Fan struct {
	Name string
}

func (f Fan) TurnOn() {
	fmt.Printf("%s 风扇已经打开\n", f.Name)
}

func (f Fan) TurnOff() {
	fmt.Printf("%s 风扇已经关闭\n", f.Name)
}

func main() {
	// 创建一个灯对象
	light := Light{Name: "客厅"}
	// 创建一个风扇对象
	fan := Fan{Name: "卧室"}

	// 使用接口调用灯和风扇的开关方法
	switchOnOff(light)
	switchOnOff(fan)
	TestConvert()
}

// 接收一个开关接口类型的参数
func switchOnOff(sw Switch) {
	sw.TurnOn()
	sw.TurnOff()
}

type GError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (g *GError) Error() string {
	return fmt.Sprintf("Code:%s, Message:%s", g.Code, g.Message)
}

func callapi() error {
	return &GError{}
}

func TestConvert() {
	err := callapi()
	if err != nil {
		if gErr, ok := err.(*GError); ok {
			fmt.Println(gErr.Error())
		} else {
			fmt.Println("system error")
		}
	}
}
