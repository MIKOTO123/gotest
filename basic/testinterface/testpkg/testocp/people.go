package testocp

type People struct {
	Name string
	Age  int
}

func (p People) Care(pet Pet) {
	println(p.Name, "care", pet.GetName())
}
