package main

import "fmt"

//学生结构体
type Student struct {
	name      string
	className string
}

//构造函数
func newStudent(name, className string) *Student {
	return &Student{
		name:      name,
		className: className,
	}
}

//特定类型的函数
func (s *Student) Show() {
	fmt.Printf("姓名：%s 班级：%s", s.name, s.className)
}

//特定类型的函数 和C#的扩展函数类似
func main() {
	newStudent("张三", "一班").Show()
}
