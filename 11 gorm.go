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

func dbConnect() *gorm.DB {
	dsn := "root:qwer1234@tcp(47.93.116.220:3306)/gotest?charset=utf8mb4&parseTime=True&loc=Local"
	var connect = mysql.Open(dsn)
	db, err := gorm.Open(connect, &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return db
}

// 增
func insertData(db *gorm.DB) {
	// 插入一条数据
	stu := &Students{7, "han234", "男", 20, "class B"}
	db.Create(&stu)
	db.Commit()
}

// 删
func deleteData(db *gorm.DB) {
	var firstStudent Students
	db.First(&firstStudent) // select * from students order by student_id limit 1
	db.Delete(&firstStudent)
}

// 改
func updateData(db *gorm.DB) {
	var student Students
	db.Take(&student)
	student.Name = "韩学成"
	student.Gender = "Female"
	student.Age = 25
	db.Save(&student)
}

func findAll(db *gorm.DB) {
	// 获取所有数据 把结果返回 allStudents 变量中
	var allStudents []Students
	db.Find(&allStudents) // select * from students
	fmt.Println(allStudents)
}

// 查
func findData(db *gorm.DB) {

	// 获取第一条数据
	var firstStudent Students
	db.First(&firstStudent) // select * from students order by student_id limit 1
	fmt.Println(firstStudent)

	// 获取最后一条数据
	var lastStudent Students
	db.Last(&lastStudent) // select * from students order by student_id desc limit 1
	fmt.Println(lastStudent)

	// 根据主键查询, 获取单条数据
	var singleStudent Students
	db.Take(&singleStudent, 2) // select * from students where student_id = 2
	fmt.Println(singleStudent)

	// 按条件查询
	var queryStudent []Students
	db.Find(&queryStudent, "class = ?", "Class A") // select * from students where class = "Class A"
	fmt.Println(queryStudent)

	// where 条件查询
	var whereStudent []Students
	db.Where("name = ?", "Alice").Find(&whereStudent)
	fmt.Println(whereStudent)
}

func main() {
	db := dbConnect()
	//findData(db)
	updateData(db)
	findAll(db)
	deleteData(db)
	findAll(db)
}
