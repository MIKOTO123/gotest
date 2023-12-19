package testpkg

import (
	"fmt"
	"sync"
	"time"
)

var i int = 100

var wg sync.WaitGroup

var lock sync.Mutex

func add() {
	defer wg.Done()
	lock.Lock()
	i += 1
	fmt.Printf("i++: %v\n", i)
	time.Sleep(time.Millisecond * 10)
	lock.Unlock()
}

func sub() {
	lock.Lock()
	time.Sleep(time.Millisecond * 2)
	defer wg.Done()
	i -= 1
	fmt.Printf("i--: %v\n", i)
	lock.Unlock()
}
func Mutexttest() {
	fmt.Println("我到底什么时候开始执行")
	for h := 0; h < 100; h++ {
		wg.Add(1)
		go add()
		wg.Add(1)
		go sub()
	}
	fmt.Println("可能在加减前打印.")
	wg.Wait()

	//上锁的话,最后这个结果i一定是100,如果不上锁,会出现脏读的情况,最后就不一定是100
	fmt.Printf("end i: %v\n", i)
}
