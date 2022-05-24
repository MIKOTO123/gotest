package testpkg

import "os"

func writeFile() {
	os.WriteFile("test2.txt", []byte("hello"), os.ModePerm)
}

func Write() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_TRUNC, 0755)
	f.Write([]byte(" hello golang"))
	f.Close()
}

func WriteString() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 0755)
	f.WriteString("hello java")
	f.Close()
}

func WriteAt() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR, 0755)
	f.WriteAt([]byte("aaa"), 3)
	f.Close()

}
