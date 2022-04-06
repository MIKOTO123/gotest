package testobj

import "fmt"

type Coder struct {
	Person
	gan string
	Pet Pet
}

func (c Coder) work() {
	fmt.Println("cording...")
}

func (c Coder) GetPet() Pet {
	return c.Pet
}
