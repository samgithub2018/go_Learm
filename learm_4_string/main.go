package main

import (
	"fmt"
	"strings"
)

func main() {

	var str1 = "sam"
	var str2 = "joh"

	//字符串拼接
	str3 := fmt.Sprintf("%s+%s", str1, str2)
	fmt.Printf(str3)
	fmt.Println()

	//字符串分割
	strArr := strings.Split(str3, "+")
	fmt.Println(strArr)

	//字符串包含
	b1 := strings.Contains(str3, "+")
	fmt.Println(b1)

	//前缀是否以sam 开头
	b2 := strings.HasPrefix(str3, "sam")
	fmt.Println(b2)

	//后缀是否已 joh 开头
	b3 := strings.HasSuffix(str3, "joh")
	fmt.Println(b3)

	//字符串简单处理
	s1 := "hello小王子"
	s2 := strings.Replace(s1,"hello","",1)
	fmt.Println(len(s2)/3)

 	s3 := strings.Split(s1,"")
 	fmt.Println(s3)
	for i := range s3 {
		fmt.Printf("%s%d\t",s3[i],len(s3[i]))
	}


}
