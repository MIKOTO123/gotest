package testpkg

import (
	"fmt"
	"runtime"
	"time"
)

func RuntimeTest3() {
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU()) //显示cpu核心数
	runtime.GOMAXPROCS(1)                                  //设置只有一个携程,打印结果,要么全a,要么全b,如果ab不睡觉,如果ab开始了睡觉,,,,在a睡觉的时候b会执行,b睡觉的时候a会执行
	go a()
	go b()

	time.Sleep(time.Second)
}

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("a:", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("b:", i)
		time.Sleep(time.Millisecond * 100)
	}
}
