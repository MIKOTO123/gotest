package main

import (
	"consttest/common"
	"fmt"
)

const PI = 3.14

func main() {

	const (
		REMOTE_ARR = "www.pikazo.com"
		PORT       = 3306
		BASE_URI   = "/caldav/"
	)

	println(common.A_ADDR)
	test()

}

func test() {
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
		_        //_只可意会
		k        //10
	)
	fmt.Println(a, b, c, d, e, f, g, h, i, k)
}

func test2() {
	println(common.A_ADDR)
}
