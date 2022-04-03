package testpkg

func Sum4(a int, b int) int {
	return a + b
}

//普通的参数是 形参 ,函数内部发生改变不会影响外部
//传入slice,map,chan都是引用类型,在函数会影响到外面
func F1(a int, s []int, map1 map[string]string) {
	a = 1000
	// 这里外部的会发生改变的
	s[0] = 100
	map1["domain"] = "smgui.com"
}

func Func3(args ...int) {
	for k, va := range args {
		println(k, va)
	}
}

func Func4(name string, age int, contacts ...string) {
	println(name, age)
	for _, contact := range contacts {
		println(contact)
	}
}
