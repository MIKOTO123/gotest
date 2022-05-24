package testpkg

import (
	"fmt"
	"runtime"
)

func RuntimeTest() {
	go show("java") // 子协程来运行
	for i := 0; i < 5; i++ {
		runtime.Gosched() // 我有权利执行任务了，让给你（其他子协程来执行）
		fmt.Printf("\"golang\": %v\n", "golang")
	}
	fmt.Println("end...")
}

func show(s string) {
	for i := 0; i < 2200; i++ {
		fmt.Printf("msg: %v\n", s)
	}
}
