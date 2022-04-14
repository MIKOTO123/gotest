package testpkg

import (
	"math/rand"
	"time"
)

func ChannelTest() {

	ch1 := make(chan int)
	defer close(ch1) //关闭channel
	go fun2(ch1)
	go fun3(ch1)
	println("go进程")
	time.Sleep(time.Second * 10)
}

func fun3(ch1 chan int) {
	println("我是不是在下一步就停止了行动")
	data := <-ch1 //从通道读取,这里会发生阻塞
	println(data)
}

func fun2(ch1 chan int) {
	rand.Seed(time.Now().UnixNano()) //设置种子,保证每次生成的随机数不一样
	v := rand.Intn(100)
	time.Sleep(time.Second * 5)
	ch1 <- v //写入通道
	println("我已经发送了数据")
	time.Sleep(time.Second * 2)
	println("func2结束了")
}
