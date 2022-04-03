package testpkg

import "fmt"

func DefineArray() {
	var a1 [2]int
	var a2 [3]string
	fmt.Printf("%v", a1)
	fmt.Printf("%v", a2)
	fmt.Printf("%T", a1)
	fmt.Printf("%T", a2)
}

func InitArray() {
	var a1 = [3]int{1, 2}
	var a2 = [2]string{"hello", "world"}
	var a3 = [...]string{"hello", "world"} //初始化的时候是多少,之后就定长了
	fmt.Printf("%v", a1)
	fmt.Printf("%v", a2)
	fmt.Printf("%v", len(a3))

}

func Test() {
	var a1 = [...]int{0: 1, 3: 100, 5: 10}
	var a2 = [...]bool{2: true, 5: false}
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
}

func AccessArray() {
	var a1 [2]int
	a1[0] = 100
	a1[1] = 200
	fmt.Printf("a1[0]: %v\n", a1[0])
	fmt.Printf("a1[1]: %v\n", a1[1])
	fmt.Println("------------")
	a1[0] = 1000
	a1[1] = 2000
	fmt.Printf("a1[0]: %v\n", a1[0])
	fmt.Printf("a1[1]: %v\n", a1[1])
}

func RangeArray() {
	// for range
	var a1 = [3]int{1, 2, 3}
	/* for i, v := range a1 {
		fmt.Printf("a1[%v]: %v\n", i, v)
	} */

	for _, v := range a1 {
		fmt.Printf("v: %v\n", v)
	}
}

func ForArray() {
	// 数组的遍历 1. 根据长度和下标
	var a1 = [3]int{1, 2, 3}
	for i := 0; i < len(a1); i++ {
		fmt.Printf("a1[%v]: %v\n", i, a1[i])
	}
	/* for i := 0; i < len(a1); i++ {
		fmt.Printf("a1[%v]: %v\n", i, a1[i])
	} */
}
