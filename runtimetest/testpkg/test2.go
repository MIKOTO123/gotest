package testpkg

import (
	"fmt"
	"runtime"
	"time"
)

func RuntimeTest2() {
	go show2()
	time.Sleep(time.Second)
}

func show2() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			runtime.Goexit()
		}
	}
}
