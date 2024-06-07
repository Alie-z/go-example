package main

import (
	"encoding/json"
	"fmt"
)

type ModelListItem struct {
	Value string
	Name  string
}

var Yiyan35 = ModelListItem{"yiyan", "ERNIE-Bot 3.5"}
var YiyanEB4 = ModelListItem{"yiyan_eb_4", "ERNIE-Bot 4.0"}
var YiyanEB = ModelListItem{"yiyan_eb", "ERNIE-bot eb"}
var YiyanEB1 = ModelListItem{"yiyan_eb1111", "ERNIE-bot eb111"}

func getDefaultModelListItems() []*ModelListItem {
	return []*ModelListItem{
		&Yiyan35,
		&YiyanEB,
	}
}
func getMultimodalChatBaseModelListItems() []*ModelListItem {
	baseItems := getDefaultModelListItems()
	baseItems = append(baseItems,
		&YiyanEB4,
	)
	return baseItems
}

func getCustomModelListItems(v ...*ModelListItem) []*ModelListItem {
	baseItems := getDefaultModelListItems()
	baseItems = append(baseItems,
		v...,
	)
	return baseItems
}
func main() {
	v, _ := json.Marshal(getDefaultModelListItems())
	v1, _ := json.Marshal(getMultimodalChatBaseModelListItems())
	v2, _ := json.Marshal(getCustomModelListItems(&YiyanEB1))

	fmt.Println(">>> 1", string(v))
	fmt.Println(">>> 2", string(v1))
	fmt.Println(">>> 2", string(v2))
}
