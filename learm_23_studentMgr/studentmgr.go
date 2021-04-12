package main

import "fmt"

//学生结构
type Student struct {
	Name      string
	Age       int
	ClassName string
}

//学生管理
type StudentMgr struct {
	allStudent map[string]Student
}

//创建全局
var StuMgr StudentMgr

//构建函数
func newStudent(name, className string, age int) *Student {
	return &Student{
		Name:      name,
		Age:       age,
		ClassName: className,
	}
}

//添加学生
func Add() {
	var (
		name, className string
		age             int
	)
	fmt.Println("请输入学生姓名：")
	fmt.Scan(&name)
	fmt.Println("请输入学生年龄：")
	fmt.Scan(&age)
	fmt.Println("请输入学生班级：")
	fmt.Scan(&className)
	StuMgr.allStudent[name] = *newStudent(name, className, age)
}

//显示所有添加的学生
func ShowAll() {
	for _, v := range StuMgr.allStudent {
		fmt.Printf("姓名：%s, 年龄：%d, 班级：%s \n", v.Name, v.Age, v.ClassName)
	}
}
