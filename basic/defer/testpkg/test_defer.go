package testpkg

import "fmt"

func Test1() {
	fmt.Println("start...")
	defer fmt.Println("step1...")
	defer fmt.Println("step2...")
	panic("我恐慌喽")
	defer fmt.Println("step3...")
	fmt.Println("end...")
}

func TestRecover() {
	test()
	fmt.Println("我其实是会继续执行的,,如果有recover的话")
}

func test() {
	defer func() {
		// revocer用于让程序恢复,这样在该函数的上一级函数继续执行
		//可以注释下面这个recover,对比看看效果
		if msg := recover(); msg != nil {
			fmt.Println(msg, "我恢复了..")
		}
	}()
	defer fmt.Println("4")
	test_func()
	fmt.Println("5")
}

func test_func() {
	defer fmt.Println("2")
	panic("我恐慌了")
	defer fmt.Println("3") //这个是不会执行的
}
