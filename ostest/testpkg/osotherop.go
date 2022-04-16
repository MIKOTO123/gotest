package testpkg

import (
	"fmt"
	"os"
)

func rename() {
	err := os.Rename("test.txt", "test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
