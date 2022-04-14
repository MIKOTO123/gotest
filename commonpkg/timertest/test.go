package timertest

import (
	"fmt"
	"time"
)

func TimerTest() {
	timer := time.NewTimer(time.Second * 2)
	fmt.Printf("time.Now(): %v\n", time.Now())
	t1 := <-timer.C // 阻塞的，指定时间到了
	fmt.Printf("t1: %v\n", t1)
}

func TimerTest2() {
	fmt.Printf("time.Now(): %v\n", time.Now())
	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	fmt.Printf("time.Now(): %v\n", time.Now())
}

func TimerTest3() {
	<-time.After(time.Second * 2) //等待2秒
	fmt.Println("....")
}

func TimerTest4() {
	timer := time.NewTimer(time.Second)

	go func() {
		<-timer.C
		fmt.Println("func...")
	}()

	s := timer.Stop() //直接中断了func的内容,后面都不执行了
	if s {
		fmt.Println("stop...")
	}

	time.Sleep(time.Second * 3)
	fmt.Println("main end...")
}

func TimerTest5() {
	fmt.Println("before")
	timer4 := time.NewTimer(time.Second * 5) //原来设置5s
	timer4.Reset(time.Second * 1)            //重新设置时间,即修改NewTimer的时间
	<-timer4.C
	fmt.Println("after")
}
