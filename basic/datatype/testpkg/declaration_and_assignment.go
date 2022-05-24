package testpkg

import "fmt"

type func1 func()
type myint int     //最后打印类型的类型的时候是myint
type inttype = int //最后打印类型的时候还是int

var name string
var age int
var i myint
var b bool
var p *int
var f1 float32
var f2 float64
var arr [3]int
var slice []int
var map1 map[string]string
var func2 func1
var ch1 chan int

func Declare() {
	//默认值
	PrintValue(name, age, i, f1, f2, p, b, arr, slice, map1, func2, ch1)

}

func Assign() {
	name := "pikazo"
	age := 20
	i = 20
	f1 := 3.14
	f2 := 3.141592
	b := true
	p := &age
	arr = [3]int{1, 2, 3} //数组是有固定长度的,
	slice = []int{4, 5, 6, 7}
	//map1 = map[string]string{"name": "pikazpo", "age": "20"} //之前的某个版本是不能用var map1 map[string]string 定义出来的是nil,不能后续赋值,的用make的才可以,后面的版本默认初始值改成map[],才可以进行后续赋值
	func2 := Assign
	//ch1 <- 1 //写入通道,必须关闭否则会报错
	//close(ch1)
	PrintValue(name, age, i, f1, f2, p, b, arr, slice, map1, func2, ch1)
}

func PrintValue(args ...interface{}) {
	for _, v := range args {
		//fmt.Println(i, v)
		fmt.Printf("%-10v %-20T %-10v\n", "v:", v, v)
	}
}
