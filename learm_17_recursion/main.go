package main

import "fmt"

//n个台阶,一次可以走1步,也可以走2步,有多少种走法。
func f1(y int) int {
	// 分析
	// 1. 有个台阶数
	// 2. 有每次走几层台阶
	// 3. 有多少种走法
	if y == 1 ||y == 2 {
		return 1
	}
	r := f1(y-1) + f1(y-2)
	return r
}

//1 1 2 3 5 8 13 21 34 ...  取第30为的数
func f2(y int) int {
	if y == 1 || y == 2 {
		return 1
	}
	r := f2(y-1) + f2(y-2)
	return r
}

func main() {
	fmt.Println( f1(30))
	fmt.Println(f2(30))
}
