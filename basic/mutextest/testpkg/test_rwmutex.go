package testpkg

import (
	"fmt"
	"sync"
	"time"
)

var rwMutex *sync.RWMutex
var wg1 *sync.WaitGroup

//多个读锁可以同时进行读取,但是读锁在读数据的时候,写锁会阻塞.写锁会阻塞读锁.写锁之间也会相互阻塞,
func TestRWmutex() {
	rwMutex = new(sync.RWMutex)
	wg1 = new(sync.WaitGroup)

	//下面例子中2456,可以同时进行读取.
	wg1.Add(6)
	go writeData(1)
	go readData(2)
	go writeData(3)
	go readData(4)
	go readData(5)
	go readData(6)

	wg1.Wait()
	fmt.Println("main..over...")
}

func writeData(i int) {
	defer wg1.Done()
	fmt.Println(i, "开始写：write start。。")
	rwMutex.Lock() //写操作上锁
	fmt.Println(i, "正在写：writing。。。。")
	time.Sleep(10 * time.Second)
	rwMutex.Unlock()
	fmt.Println(i, "写结束：write over。。")
}

func readData(i int) {
	defer wg1.Done()

	fmt.Println(i, "开始读：read start。。")

	rwMutex.RLock() //读操作上锁
	fmt.Println(i, "正在读取数据：reading。。。")
	time.Sleep(10 * time.Second)
	rwMutex.RUnlock() //读操作解锁
	fmt.Println(i, "读结束：read over。。。")
}
