package main

import "fmt"

type Student struct {
	name      string
	className string
	sex       string
	age       int8
	code      int8
}


//构造函数，返回一个Student对象
func newStudent(name, className, sex string, age, code int8) *Student {
	return &Student{
		name:      name,
		className: className,
		sex:       sex,
		age:       age,
		code:      code,
	}
}

func main() {
	stu1 := Student{
		name:      string("3"),
		className: string("3"),
		sex:       string("男"),
		age:       18,
		code:      20,
	}
	fmt.Println(stu1)

	//连续的内存
	fmt.Printf("age=%v\n", &stu1.age)               // age= 0xc00006e3d0
	fmt.Printf("code=%v\n", &stu1.code)             // code=0xc00006e3d1
	fmt.Printf("sex=       %v \n", &stu1.sex)       // sex=       0xc000042060
	fmt.Printf("name=      %v \n", &stu1.name)      // name=      0xc000042040
	fmt.Printf("className= %v \n", &stu1.className) // className= 0xc000042050

	var stu2 = Student{
		name:      string("joh"),
		age:       18,
		className: string("3"),
	}
	fmt.Println(stu2)

	//创建一个Student 返回指针类型 *Student
	stu3 := new(Student)
	stu3.name = string("joh")
	stu3.age = 18
	fmt.Println(stu3)
	fmt.Printf("%p \n", stu3)

	//创建Student，返回指针类型
	stu4 := &Student{
		name:      string("joh"),
		age:       18,
		className: string("3"),
	}
	fmt.Printf("stu4 = %p \n", stu4)
	//测试是引用 还是 值类型
	f1(*stu4)
	f3(stu4)

	//匿名结构体
	type stu struct {
		name      string
		className string
		sex       string
		age       int8
		code      int8
	}

	s1 := new(stu)
	s1.name = "joh"
	s1.code = 1

	fmt.Println(s1)

}

func f1(stu Student) {
	fmt.Printf("f1 = %p \n", &stu) //f1 = 0xc0000421c0
	f2(stu)
}

func f2(stu Student) {
	fmt.Printf("f2 = %p \n", &stu) //f2 = 0xc000042200
}

func f3(stu *Student) {
	fmt.Printf("f3 = %p \n", stu)
}
