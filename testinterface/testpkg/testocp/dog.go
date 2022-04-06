package testocp

type Dog struct {
	Name string
	Age  int
}

func (d Dog) Eat() {
	println("eat creazy")
}

func (d Dog) Sleep() {
	println("sleep with noice")
}

func (d Dog) GetName() string {
	return d.Name
}
