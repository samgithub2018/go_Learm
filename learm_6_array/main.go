package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3}
	fmt.Println(a1)

	//根据元素个数创建数组
	a2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a2)

	//根据索引给数组赋值
	a3 := [5]int{0: 1, 3: 4}
	fmt.Println(a3)

	for i := 0; i < len(a3); i++ {
		fmt.Println(a3[i])
	}

	//多维数组
	a4 := [3][2]int{{1, 10}, {2, 20}, {9, 30}}
	fmt.Println(a4)

	//使用range遍历数组
	for _, v := range a4 {
		fmt.Println(v)
	}

}
