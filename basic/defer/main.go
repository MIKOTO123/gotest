package main

import (
	"defer/testpkg"
	"fmt"
)

func main() {
	defer fmt.Println("我是main里面的defer")
	testpkg.Test1()
}
