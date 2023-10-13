package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func connectDb() (engine *xorm.Engine, err error) {
	engine, err = xorm.NewEngine("mysql", "root:qwer1234@tcp(47.93.116.220:3306)/gotest?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return engine, err
	}
	// 测试数据库
	err = engine.Ping()
	if err != nil {
		fmt.Println(err)
		return engine, err
	}
	fmt.Println("连接成功")
	return engine, err
}

func main() {
	engine, err := connectDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	sql := "select * from students"
	res, err := engine.QueryString(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	for i := 0; i < len(res); i++ {
		content := fmt.Sprintf("年龄: %v, 班级: %v, 性别: %v, 名称: %v",
			res[i]["age"], res[i]["class"], res[i]["gender"], res[i]["name"])
		fmt.Println(content)
	}
}
