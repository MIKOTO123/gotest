package testpkg

import (
	"fmt"
	"os"
)

func Remove() {
	err := os.Remove("a.txt") // 删除文件或者目录,不能递归删除
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
func RemoveAll() {
	err := os.RemoveAll("a.txt") // 删除文件或者目录,可以递归删除
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
