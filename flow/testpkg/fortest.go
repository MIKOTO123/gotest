package testpkg

import (
	"fmt"
	"time"
)

func Fortest() {
	//死循环
	for {
		fmt.Println("run...")
		time.Sleep(1)
		break //跳出这个死循环
	}
}

func ForTest2() {
	i := 1
	for i <= 10 {
		println(i)
		i++
	}
}

func Fortest3() {
	i := 1
	for ; i <= 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

func Fortest4() {
	for i := 0; i < 10; i++ {
		println(i)
	}
}

func ForRange() {
	var map1 = make(map[string]string)
	map1["name"] = "pikazo" //panic: assignment to entry in nil map
	map1["age"] = "20"
	map1["email"] = "piakzo@smgui.cn"
	for i, s := range map1 {
		fmt.Println(i, s)
	}
}

func ForRange2() {
	a1 := [4]int{1, 2, 3, 4}
	for _, value := range a1 {
		fmt.Println(value)
	}
}

func ForRange3() {
	s := "hello world"
	for _, v := range s {
		fmt.Printf("v: %c\n", v)
	}
}
