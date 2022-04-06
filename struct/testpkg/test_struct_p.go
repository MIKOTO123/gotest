package testpkg

import (
	"fmt"
	"struct/testobj"
)

//结构体指针

func TestPoint() {
	person := new(testobj.Person)
	//正常访问应该是要:*person.Name = "张三",不过可以省略
	person.Name = "张三"
	person.Age = 20
	fmt.Println(*person)
}

func TestPoint2() {
	var p *testobj.Person
	tom := testobj.Person{
		Id:    123,
		Name:  "haha",
		Age:   123,
		Email: "",
	}
	p = &tom
	fmt.Println(p)
}
