package main

import "fmt"

type animal interface {
	call()
}

type dog struct {
}

func (d dog) call() {
	fmt.Println("this dog call")
}

type cat struct {
}

func (c *cat) call() {
	fmt.Println("this cat call")
}



func main() {
	var an animal

	/*
		特定函数，实现接口的时候会有两种情况，参数接收有两种，分别是值（结构是值类型）和指针
		接收值类型，那么是下面情况【1】，实例化的时候可以是指针，也可以是值类型
		接收指针类型，那么下面情况【2】，实例化的时候只能接收指针类型
	*/

	//【1】
	an = dog{}
	an = &dog{}
	fmt.Printf("%T \n", an) //*main.dog
	an.call()

	//【2】
	//an = cat{} 报错
	an = &cat{}
	fmt.Printf("%T \n", an) //*main.call
	an.call()
}
