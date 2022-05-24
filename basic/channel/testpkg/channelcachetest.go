package testpkg

import (
	"strconv"
	"time"
)

//注意不要造成死循环
func ChannelCacheTest() {
	ch1 := make(chan int, 3)
	go taks1(ch1)
	//在主函数这里等待接受
	time.Sleep(time.Second * 3)
	//通过如果通过循环读取,,,如果没有设定关闭通道,就会一直死循环
	for v := range ch1 { //会一直阻塞,直到通道关闭
		time.Sleep(time.Second * 1)
		println("读出" + strconv.Itoa(v))
	}
	println("go进程")
	time.Sleep(time.Second * 30)
}

func taks1(ch1 chan int) {
	for i := 0; i < 10; i++ {
		ch1 <- i                        //写入通道
		println("输入" + strconv.Itoa(i)) //3个就满了,然后就阻塞
	}
	close(ch1) //关闭通道,不然会造成死循环报错
}
