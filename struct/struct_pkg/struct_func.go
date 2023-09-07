// 结构体方法

package struct_pkg

import "fmt"

type Student struct {
	Name string
	Age  int
	Sex  int
}

// 结构体继承

type Info struct {
	Account int
	Student
}

/* 创建方法 */

func (s Student) Add(name string, age int, sex int) {
	s.Name = name
	s.Age = age
	s.Sex = sex

	fmt.Printf("name=%v age=%v sex%v \n", s.Name, s.Age, s.Sex)
}

func (s Student) Show() Student {
	fmt.Printf("name=%v age=%v sex%v \n", s.Name, s.Age, s.Sex)
	return s
}

func (s *Student) Pointer() *Student {
	fmt.Printf("name=%v age=%v sex%v \n", s.Name, s.Age, s.Sex)
	return s
}

func (i Info) Show() Info {
	fmt.Printf("Accout=%v Student=%v", i.Account, i.Student)
	return i
}
