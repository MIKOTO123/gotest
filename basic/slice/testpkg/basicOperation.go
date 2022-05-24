package testpkg

import "fmt"

func Add() {
	s1 := []int{100, 200, 300}
	println(cap(s1))
	s1 = append(s1, 400)
	println(cap(s1)) //扩容之后,成倍增长
	s1 = append(s1, 500)
	println(cap(s1))
	fmt.Printf("%v\n", s1)
}

func Del() {
	s1 := []int{100, 200, 300, 4, 5}
	s1 = append(s1[:2], s1[3:]...)
	fmt.Printf("%v\n", s1)
}

// copy
func Copytest() {
	var s1 = []int{1, 2, 3, 4}
	// var s2 = s1  //这种可以看成同一个切片,因为地址是一样的
	var s2 = make([]int, 3)
	copy(s2, s1)
	s2[0] = 100
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2) //2的容量不会改变 结果:s2: [100 2 3]
}

// query
func Query() {
	var s1 = []int{1, 2, 3, 4}
	var key = 2
	for i, v := range s1 {
		if v == key {
			fmt.Printf("i: %v\n", i)
		}
	}
}

// update
func Update() {
	var s1 = []int{1, 2, 3, 4}
	s1[1] = 100
	fmt.Printf("s1: %v\n", s1)
}
