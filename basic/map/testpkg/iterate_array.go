package testpkg

func ForMap() {
	map1 := make(map[string]string)
	map1["name"] = "pikazo"
	map1["age"] = "20"
	map1["email"] = "smgui@smgui.com"
	for k, value := range map1 {
		println(k, value)
	}
}
