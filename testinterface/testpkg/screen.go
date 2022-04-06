package testpkg

type Screen struct {
	Name   string
	Width  int
	Height int
}

func (s Screen) Read() {
	println("read") //只实现了读这个方法,会报错
}

//func (s Screen) Write() {
//	println("Write")
//}
