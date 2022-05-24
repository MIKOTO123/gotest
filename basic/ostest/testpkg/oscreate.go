package testpkg

import (
	"fmt"
	"os"
)

func Create() {
	// 等价于：OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
	create, err := os.Create("a.txt")
	if err != nil {
		panic(err)
	} else {
		println("成功的创建了文件" + create.Name())
		create.Close()
	}
}
func CreateTemp() {
	// 第一个参数 目录默认：Temp 第二个参数 文件名前缀
	f2, _ := os.CreateTemp("", "temp")
	fmt.Printf("f2.Name(): %v\n", f2.Name())
}
