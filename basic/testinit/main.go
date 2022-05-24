package main

import (
	"fmt"
	"testinit/testpkg"
	"testinit/testpkg3"
)

//init是没有参数
func init() {
	fmt.Println("我是init")
}

//一个员文件里面可以有多个init,但是意义不名
func init() {
	fmt.Println("我是init2")
}

var smgui = printtheorder()

func printtheorder() interface{} {

	fmt.Println("我是var定义")
	return 1
}

func main() {
	fmt.Println("我是main")
	testpkg.Test()
	testpkg3.Test3()
}
