package testpkg

import "fmt"

func DefineSlice() {
	var s1 []int
	var s2 []string
	var s3 = make([]int, 2)
	println(s1)
	fmt.Printf("%v\n", s1)
	println(s2)
	fmt.Printf("%v\n", s2)
	println(s3)
	println(s1 == nil)
	println(s2 == nil)
	println(s3 == nil)
}

func BasicOperation() {
	s1 := []int{1, 2, 3}
	println(len(s1))
	println(cap(s1))
	println(s1[2])
}

func BasicOperation2() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	//[)
	s2 := s1[:3] //1 2 3
	s3 := s1[2:] //3,4,5,6
	s4 := s1[:]
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2) //123
	fmt.Printf("%v\n", s3) //3456
	fmt.Printf("%v\n", s4) //123456

}

func ForSlice() {
	s1 := []int{1, 2, 3, 3, 4, 5}
	for i := 0; i < len(s1); i++ {
		println(s1[i])
	}
}

func RangeSlice() {
	s1 := []int{1, 2, 3, 3, 4, 5}
	for i, value := range s1 {
		println(i, value)
	}
}
