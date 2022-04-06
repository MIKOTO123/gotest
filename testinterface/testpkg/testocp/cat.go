package testocp

type Cat struct {
	Name string
	Age  int
}

func (c Cat) Eat() {
	println("eat slowly")
}

func (c Cat) Sleep() {
	println("sleep quietly")
}

func (c Cat) GetName() string {
	return c.Name
}
