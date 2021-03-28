package main

import (
	"fmt"
	"strings"
)

func main() {

	//
	//for i := 1; i < 10; i++ {
	//	fmt.Println(i)
	//}
	//
	//i := 0
	//for ; i < 10; i++ {
	//	fmt.Println(i)
	//}
	//
	//for i < 20 {
	//	fmt.Println(i)
	//	i++
	//}
	//
	//for true {
	//	fmt.Println("无限循环")
	//}

	str := "hello 打本"
	strArr := strings.Split(str, "")
	fmt.Println(strArr)

}
