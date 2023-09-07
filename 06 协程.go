// 主线程结束 协程就会结束

package main

import (
	"fmt"
	"time"
)

func SendMsg() {
	fmt.Println("发送一条短信")
	time.Sleep(3 * time.Second)
	fmt.Println("短信发送完毕")
}

func main() {
	fmt.Println("用户等待接受验证码")
	go SendMsg()
	fmt.Println("验证码已发送")
}
