package testpkg

import (
	"testinterface/testpkg/testcompose"
	"testinterface/testpkg/testocp"
)

func Test() {
	mobile := Mobile{Name: "xiaomi"}
	computer := Computer{Name: "huipu"}

	mobile.Write()
	mobile.Read()
	computer.Write()
	computer.Read()

}

func TestInterface() {
	//var usb USB
	//usb = &Screen{Name: "feilipu"}  //要使用接口的话,必须2个接口都实现,否则会报错
	//fmt.Printf("%v\n", usb)
}

func TestComposer() {

	var ff testcompose.FlyFish
	ff = testcompose.Fish{}
	ff.Fly()
	ff.Swim()
}

func TestOcp() {
	p1 := &testocp.People{
		Name: "pikazo",
	}
	dog1 := testocp.Dog{Name: "small white"}

	p1.Care(dog1)

}
