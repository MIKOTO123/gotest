package testpkg

import "fmt"

func AddContinue() func(int) int {
	var x int
	return func(a int) int {
		x += a
		return x
	}
}

func TestClosure() {
	var f = AddContinue()
	fmt.Printf("f(10):%v\n", f(10))
	fmt.Printf("f(20):%v\n", f(20))
	fmt.Printf("f(30):%v\n", f(30))
}
