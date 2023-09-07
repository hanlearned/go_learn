// 接口使用

package struct_pkg

import "fmt"

// 定义一个接口

type Animal interface {
	Sing() // 注意接口声明是 区分大小写的
	Jump()
	Rap()
}

type Chicken struct {
	Name string
}

func (c Chicken) Sing() {
	fmt.Printf("%v 会唱歌 \n", c.Name)
}

func (c Chicken) Jump() {
	fmt.Printf("%v 会跳 \n", c.Name)
}

func (c Chicken) Rap() {
	fmt.Printf("%v 会rap \n", c.Name)
}
