package testpkg

import "fmt"

type Person struct {
	id    int
	name  string
	age   int
	email string
}

func Test() {
	p1 := Person{
		id:    1,
		name:  "pikazo",
		age:   18,
		email: "smgui@smgui.com",
	}
	name := p1.name
	fmt.Printf("%-10v %-20T %-10v\n", "p1:", p1, p1)
	fmt.Printf("name:%v\n", name)
}
