package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	//json序列化是必须以大写开头的字段，否则外部无法访问到该字段
	Name string
	//标记Age，序列化的时候名字改成age
	Age  int `json:"age"`
	//该字段不会被访问到，序列化到
	className string
}

func main() {
	stu1 := Student{
		Name:      "sam",
		Age:       18,
		className: "研一",
	}
	fmt.Println(stu1)
	s1, err1 := json.Marshal(stu1)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	//将byte转成字符串
	fmt.Println(string(s1))
	s2 := string(s1)

	//将字符串转成结构
	stu2 := new(Student)
	err2 := json.Unmarshal([]byte(s2), stu2)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(*stu2)
}
