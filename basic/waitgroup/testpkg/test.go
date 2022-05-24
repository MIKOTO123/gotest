package testpkg

import (
	"fmt"
	"sync"
)

func TestWg() {
	var wg sync.WaitGroup
	wg.Add(2) //counter  3 2 1

	go fun1(&wg)
	go fun2(&wg)

	fmt.Println("main 进入阻塞状态。。等待wg中的子goroutien结束。。")
	wg.Wait() // 表示main goroutine进入等待，意味着阻塞
	fmt.Println("main..解除阻塞。。")
}

func fun1(wg *sync.WaitGroup) {
	for i := 1; i < 10; i++ {
		fmt.Println("fun1()函数中打印。。A ", i)
	}

	wg.Done() //给wg等待组中的counter数值减1。同	wg.Add(-1)
}

func fun2(wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 1; j < 10; j++ {
		fmt.Println("\tfun2()函数打印。。", j)
	}
}
