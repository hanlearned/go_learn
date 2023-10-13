package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Students struct {
	//gorm.Model
	StudentId int `gorm:"primaryKey"`
	Name      string
	Gender    string
	Age       int
	Class     string
}

func main() {
	dsn := "root:qwer1234@tcp(47.93.116.220:3306)/gotest?charset=utf8mb4&parseTime=True&loc=Local"
	var connect = mysql.Open(dsn)
	db, err := gorm.Open(connect, &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	//stu := &Students{6, "han123", "男", 19, "class A"}
	//db.Create(&stu)

	db.Commit()
}
