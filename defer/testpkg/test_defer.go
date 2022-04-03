package testpkg

import "fmt"

func Test1() {
	fmt.Println("start...")
	defer fmt.Println("step1...")
	defer fmt.Println("step2...")
	defer fmt.Println("step3...")
	fmt.Println("end...")
}
