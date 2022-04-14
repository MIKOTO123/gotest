package testpkg

func Definition() {
	chan1 := make(chan int)     //无缓冲
	chan2 := make(chan int, 10) //有缓冲
	println(chan1)
	println(chan2)
}
