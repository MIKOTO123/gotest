package testpkg

import "fmt"

func Test1() {
	fmt.Println("我既没有参数，也没有返回值...")
}

func Sum2(a int, b int) (ret int) {
	ret = a + b
	return //这里直接就return,因为上面已经定义了ret要return
}

func Sum3(a int, b int) int {
	// ret := a + b
	return a + b
}

func Return2(age int, name string) (int, string) {
	return age, name
}
