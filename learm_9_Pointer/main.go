package main

import "fmt"

func main() {

	a := new(int)
	fmt.Println(a)
	fmt.Printf("%T \n",a)
	fmt.Println(*a)
	*a = 100
	fmt.Println(*a)

	//指针的指针变量
	fmt.Println(&a)

}
