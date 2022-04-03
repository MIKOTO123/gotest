package main

import (
	"datatype/testpkg"
	"fmt"
	"github.com/MIKOTO123/pikazocommonfunc"
	"math"
)

func main() {
	//testpkg.Assign()
	testpkg.Assign()
	//test2()
}

/**
数据类型,用于数字
*/
func test2() {

	pikazocommonfunc.PrintStartInFunc(3)
	defer pikazocommonfunc.PrintEndInFunc(3)
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77
	fmt.Printf("%d \n", b) // 63 ,,,8进制的77变成十进制就是66

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF
	fmt.Printf("%d \n", c) // 255 ,,,16进制的ff变成十进制就是255
}
