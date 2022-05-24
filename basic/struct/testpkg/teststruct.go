package testpkg

import (
	"fmt"
	"struct/testobj"
)

func TestStruct() {
	p1 := testobj.Person{
		Id:    1,
		Name:  "pikazo",
		Age:   18,
		Email: "smgui@smgui.com",
	}
	name := p1.Name
	fmt.Printf("%-10v %-20T %-10v\n", "p1:", p1, p1)
	fmt.Printf("name:%v\n", name)
	p1.Cry()
}
