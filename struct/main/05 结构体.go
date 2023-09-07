package main

import (
	"encoding/json"
	"fmt"
	"go_learn/go_learn/struct/struct_pkg"
)

func main() {
	stu := struct_pkg.Student{}
	stu.Add("hanxuecheng", 1, 1)

	stu1 := struct_pkg.Student{}
	stu1.Show()

	stu2 := &struct_pkg.Student{
		Name: "韩学成",
		Age:  18,
		Sex:  1,
	}
	stu2.Pointer()

	info := struct_pkg.Info{
		Student: struct_pkg.Student{
			Name: "韩学成",
		},
		Account: 10001,
	}
	info.Show()

	class := struct_pkg.Class{
		Name:          "15班",
		Age:           "高一",
		StudentNumber: 41,
	}
	fmt.Println(class)

	// 结构体 转 json
	jsonData, err := json.Marshal(class)
	if err != nil {
		return
	}
	fmt.Println(string(jsonData))
}
