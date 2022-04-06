package testpkg

import (
	"fmt"
	"struct/testobj"
)

func showPerson(p testobj.Person) {
	p.Name = "showPerson"
}

func showPerson2(p *testobj.Person) {
	p.Name = "showPerson2"
}

func TestArgs() {
	p := testobj.Person{
		Name: "test",
		Age:  10,
	}
	showPerson(p) //值复制
	fmt.Println(p)
	showPerson2(&p) //地址传过去了
	fmt.Println(p)
}
