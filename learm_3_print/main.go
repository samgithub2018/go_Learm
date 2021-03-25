package main

import (
	"fmt"
)

func main() {

	//打印布尔型
	var b = true
	fmt.Printf("b:%t \n", b)

	//打印整数,也是十进制
	var age = 10
	fmt.Printf("age:%d \n", age)

	//18的打印二进制 1010
	fmt.Printf("age:%b \n", age)

	//打印八进制 12
	fmt.Printf("age:%o \n", age)

	//打印十六进制，第十位就是a
	fmt.Printf("age:%x \n", age)

	//打印指针
	fmt.Printf("age:%p \n", &age)

	//打印字符串
	var str1 = "sam"
	fmt.Printf("str:%s \n", str1)

}
