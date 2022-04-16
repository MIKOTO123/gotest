package testpkg

import (
	"fmt"
	"os"
)

func Wd() {
	dir, _ := os.Getwd() //获取当前工作目录
	fmt.Printf("dir: %v\n", dir)
	os.Chdir("d:/") //设置当前工作目录
	dir, _ = os.Getwd()
	fmt.Printf("dir: %v\n", dir)

	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}

func MakeDir() {
	err := os.Mkdir("a", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

}

func MakeDirRecursively() {
	err2 := os.MkdirAll("a/b/c", os.ModePerm) //递归创建目录
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
}
