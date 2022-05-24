package testpkg

import "fmt"

//阶乘
func Factorial(a int) int {
	if a == 1 {
		return 1
	}
	return a * Factorial(a-1)
}

func TestDigui() {
	fmt.Printf("Factorial(5):%v\n", Factorial(5))
}
