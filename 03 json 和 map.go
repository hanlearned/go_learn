package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 创建一个 map
	var map1 = make(map[string]interface{})
	map1["name"] = "hanxuecheng"
	map1["age"] = 10

	fmt.Println(map1)

	// 将 map1 转换为 json1
	json1, err := json.Marshal(map1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(json1))

	// 创建一个 map2 ，将 json1 转换为 map2
	var map2 = make(map[string]interface{})
	err = json.Unmarshal([]byte(string(json1)), &map2)
	if err != nil {
		return
	}
	fmt.Println(map2)
}
