package testpkg

import "fmt"

func TestAmn() {
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}(1, 2)
	fmt.Printf("max:%v\n", max)
}
