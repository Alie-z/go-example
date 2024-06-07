package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	RawMessage json.RawMessage `json:"data"`
}

func main() {
	// 假设我们有一个未知结构的 JSON 数据
	rawJSON := []byte(`{"data": {"name": "John", "age": 30}}`)

	// 将 JSON 数据解码到 Data 结构中
	var data Data
	if err := json.Unmarshal(rawJSON, &data); err != nil {
		fmt.Println("解码 JSON 出错:", err)
		return
	}

	// 输出未解析的 JSON 数据
	fmt.Println("未解析的 JSON 数据:", string(data.RawMessage))

	// 可以稍后解析未解析的 JSON 数据
	var parsedData map[string]interface{}
	if err := json.Unmarshal(data.RawMessage, &parsedData); err != nil {
		fmt.Println("解码未解析的 JSON 出错:", err)
		return
	}

	// 输出解析后的数据
	fmt.Println("解析后的数据:", parsedData)
}
