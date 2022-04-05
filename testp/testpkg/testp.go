package testpkg

import "fmt"

func DefineP() {
	var p *int
	var p2 **int
	fmt.Printf("p:%v\n", p)
	fmt.Printf("%-10v %-20T %-10v\n", "p:", p, p)

	i := 100
	p = &i
	fmt.Printf("%-10v %-20T %-10v\n", "p:", p, p)
	fmt.Printf("%-10v %-20T %-10v\n", "*p:", *p, *p)

	p2 = &p
	fmt.Printf("%-10v %-20T %-10v\n", "p2:", p2, p2)
	fmt.Printf("%-10v %-20T %-10v\n", "*p2:", *p2, *p2)
	fmt.Printf("%-10v %-20T %-10v\n", "**p2:", **p2, **p2)

}

func ArrP() {
	var arrp [3]*int //一个存放指针的数组
	arr := [3]int{1, 2, 4}
	for i, _ := range arr {
		arrp[i] = &arr[i]
	}
	//for i, v := range arr {
	//	arrp[i] = v //注意和上面额for做对,因为v是复制的值,所以后面用*arrp[1]=100改变了也不影响原数组
	//}
	fmt.Printf("%-10v %-20T %-10v\n", "arrp:", arrp, arrp)
	*arrp[1] = 100
	fmt.Printf("%-10v %-20T %-10v\n", "arr:", arr, arr)

}

//todo:something about the point
