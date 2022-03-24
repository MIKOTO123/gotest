package main

import (
	"fmt"
	"github.com/MIKOTO123/pikazocommonfunc"
	"math"
)

func main() {
	test1()
	test2()
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

func test1() {
	pikazocommonfunc.PrintStartInFunc(3)
	defer pikazocommonfunc.PrintEndInFunc(3)
	name := "pikazo"
	age := 20
	b := true
	p := &age
	array := [...]int{1, 2, 3, 4} //数组是有固定长度的,
	slice := []int{4, 5, 6, 7}
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", array)
	fmt.Printf("%T\n", slice)
}
