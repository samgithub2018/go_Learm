package main

//此项目需要单独打开
//并将项目目录添加到GOPATH

func main() {
	StuMgr = StudentMgr{
		allStudent: make(map[string]Student, 0),
	}
	Add()
	ShowAll()
}
