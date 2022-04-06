package testpkg

type Computer struct {
	Name string
}

func (c Computer) Write() {
	println("Computer.Write()")
}

func (c Computer) Read() {
	println("Computer.Read()")
}
