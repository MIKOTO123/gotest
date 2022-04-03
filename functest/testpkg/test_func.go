package testpkg

func Sum(a int, b int) int {
	return a + b
}

//这里的max,因为在前面已经定义了,所以西面不用 max:= 而是直接=
func Comp(a int, b int) (max int) {
	if a > b {
		max = a
	} else {
		max = b
	}
	return max
}
