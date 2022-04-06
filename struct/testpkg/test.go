package testpkg

import (
	"fmt"
	"struct/testobj"
)

func Test() {

	Coder := new(testobj.Coder)
	Coder.Name = "test"
	Coder.Pet = testobj.Dog{
		Name: "aba",
		Age:  23,
	}
	fmt.Printf("%v\n", *Coder)
	Coder.GetPet().Eat()
}
