package testpkg

import (
	"fmt"
	"io"
	"os"
)

func readFile() {
	b, _ := os.ReadFile("test2.txt")
	fmt.Printf("b: %v\n", string(b[:])) //bytes数组可以转string
}

func Read() {
	f, _ := os.Open("a.txt")
	for {
		buf := make([]byte, 3)
		n, err := f.Read(buf) //n表示实际读取到的几个字符,比如读取到了2个字符,那么n=2
		if err == io.EOF {
			break
		}
		fmt.Printf("n: %v\n", n)
		fmt.Printf("string(buf): %v\n", string(buf))
	}
	f.Close()
}

func ReadAt() {
	f, _ := os.Open("a.txt")
	buf := make([]byte, 4)
	n, _ := f.ReadAt(buf, 3)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf))
}

func ReadDir() {
	de, _ := os.ReadDir("morestrings")
	for _, v := range de {
		fmt.Printf("v.IsDir(): %v\n", v.IsDir())
		fmt.Printf("v.Name(): %v\n", v.Name())
	}
}

func Seek() {
	f, _ := os.Open("a.txt")
	f.Seek(3, 0) //0文件开头,1当前位置,2文件结尾
	buf := make([]byte, 3)
	n, _ := f.Read(buf)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf))
	f.Seek(-1, 2) //从结尾出-3的位置开始读
	buf = make([]byte, 2)
	n, _ = f.Read(buf)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf))
	f.Close()
}
