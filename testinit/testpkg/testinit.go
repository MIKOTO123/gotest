package testpkg

import (
	"fmt"
	"testinit/testpkg2"
)

func init() {
	fmt.Println("我是testpkg的init")
}

func Test() {
	testpkg2.Test2()
}
