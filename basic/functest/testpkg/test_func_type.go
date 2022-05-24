package testpkg

import "fmt"

func FuncType() {
	type f1 func(int, int) int
	var ff f1
	ff = sum
	r := ff(1, 2)
	fmt.Printf("r: %v\n", r)

	ff = max
	r = ff(1, 2)
	fmt.Printf("r: %v\n", r)
}

func Test(name string, func1 func(string)) {
	func1(name)
}

func Cal(operation string) func(int, int) int {
	switch operation {
	case "+":
		return sum
	case "max":
		return max
	default:
		return nil
	}
}

func sayHello(name string) {
	fmt.Printf("Hello,%s", name)
}

func sum(a int, b int) int {
	return a + b
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
