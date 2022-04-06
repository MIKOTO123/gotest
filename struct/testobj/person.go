package testobj

import "fmt"

type Person struct {
	Id    int
	Name  string
	Age   int
	Email string
}

func (p Person) Cry() {
	fmt.Println("wuwuwuwuwuwu")
}

func (p Person) work() {
	fmt.Println("working...")
}
