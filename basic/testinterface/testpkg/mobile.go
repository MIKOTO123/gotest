package testpkg

type Mobile struct {

	// Name of the mobile
	Name string

	// Price of the mobile
	Price float64

	// Color of the mobile
	Color string
}

func (m Mobile) Read() {
	println("Mobile.Read()")
}

func (m Mobile) Write() {
	println("Mobile.Read()")
}
