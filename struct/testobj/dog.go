package testobj

import "fmt"

type Dog struct {
	Name string
	Age  int
}

func (d Dog) Eat() {
	println("Eat")
}

func (d Dog) Sleep() {
	println("Sleep")
}

func (d Dog) bark() {
	fmt.Println("wang!wang!")
}
