package testpkg

import "fmt"

func BasicOperate() {
	var m1 = map[string]string{"name": "tom", "age": "20", "email": "tom@gmail.com"}
	var k1 = "name"
	var k2 = "age1"
	v, ok := m1[k1]
	fmt.Printf("v: %v\n", v)
	fmt.Printf("ok: %v\n", ok)
	fmt.Println("----------")
	v, ok = m1[k2]
	fmt.Printf("v: %v\n", v)
	fmt.Printf("ok: %v\n", ok)

}

func DefineMap() {
	var map1 = map[string]string{"name": "pikazo", "age": "20", "email": "smgui@smgui.com"}
	map2 := make(map[string]string)
	map2["name"] = "mikoto"
	map2["age"] = "20"
	map2["email"] = "mikoto@smgui.com"
	fmt.Printf("%v\n", map1)
	fmt.Printf("%v\n", map2)
}

func DeclareMap() {
	var map1 map[string]string //这种方式后续不能赋值
	map2 := make(map[string]string)
	fmt.Printf("%v\n", map1)
	fmt.Printf("%v\n", map2)
}
