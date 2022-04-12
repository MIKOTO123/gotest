package testpkg

import (
	"fmt"
	"time"
)

func ShowMsgWithTime(msg interface{}, a int) {
	for i := 0; i < a; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Printf("%v\n", msg)
	}
}

func GoTest() {
	go ShowMsgWithTime("Hello", 10)
	go ShowMsgWithTime("World", 10)
	time.Sleep(time.Millisecond * 2000)
	fmt.Println("Hello World")
}
