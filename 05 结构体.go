package main

import "fmt"

// 创建User结构体

type User struct {
	Name string
	Ids  uint32
}

type Sex struct {
}

// Account 继承 User

type Account struct {
	User
	money float32
	Name  string
}

func main() {
	// 实例化
	ac := Account{
		money: 10.1,
		Name:  "hanxuecheng",
		User: User{
			Name: "韩学成",
			Ids:  10001,
		},
	}
	fmt.Println(ac)

}
