package main

import "fmt"

func main() {

	//fmt.Println("请输入你想输入的字符或数字，然后按回车完成！")
	//var s1 string
	//fmt.Scan(&s1)
	//fmt.Println(s1)

	//fmt.Println("请输入你的姓名和年龄，请输入在一行，并用空格分开！")
	//var (
	//	name string
	//	age  int
	//)
	//fmt.Scanf("%s%d", &name, &age)
	//fmt.Printf("name:%s age:%d",name, age)

	//fmt.Println("请输入你的姓名和年龄，请输入在一行，并用空格分开！")
	//var (
	//	name string
	//	age  int
	//)
	//fmt.Scanln(&name,&age)
	//fmt.Printf("name:%s age:%d",name, age)

	s := "hello"
	for _, v := range s {
		if v=='l' {
			fmt.Println("l")
		}
	}

}
