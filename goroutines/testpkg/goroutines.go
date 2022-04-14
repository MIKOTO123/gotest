package testpkg

import (
	"fmt"
	"time"
)

func GoTest() {
	go ShowMsgWithTime("Hello", 10)
	go ShowMsgWithTime("World", 10)
	time.Sleep(time.Millisecond * 2000)
	fmt.Println("Hello World")
}

func GoTest2() {
	go responseSize("http://www.baidu.com")
	go responseSize("https://www.bilibili.com/")
	go responseSize("https://www.feijisu09.com")
	time.Sleep(10 * time.Second)
}

func GoTestQuit() {
	go ForEachPrintUntil("Hello", 10)
	go ForEachPrintUntil("smgui", 20)
	time.Sleep(time.Second * 10)
}
