package testpkg

import (
	"fmt"
	"os"
)

func Open() {
	f, err := os.Open("a.txt") //只读方式打开文件,内部调用的仍然是OpenFile
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
		f.Close()
	}
}

func OpenClose() {
	f, err := os.OpenFile("a1.txt", os.O_RDWR|os.O_CREATE, 755)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
		f.Close()
	}

}
