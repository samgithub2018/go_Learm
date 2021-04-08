package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	className string
	//结构体嵌套，能实现一些简单的字段继承，和方法继承
	//如果嵌套多个结构体，那么如果存在相同的字段，那么必须有各自的"名字" , p Person
	Person
}

func (p Person) showName() {
	fmt.Printf("你的姓名是：%s \n", p.name)
}

func (s Student) showClass() {
	fmt.Printf("你的班级是%s \n", s.className)
}

func main() {
	stu := Student{
		className: "研发一班",
		Person: Person{
			name: "sam",
			age:  18,
		},
	}

	fmt.Println(stu)
	//可以直接使用person的字段，但一个结构体中嵌套不能出现同名的字段
	fmt.Println(stu.name)

	//将嵌套结构体的，扩展函数继承过来了
	stu.showName()

	//自己的扩展函数
	stu.showClass()

}
