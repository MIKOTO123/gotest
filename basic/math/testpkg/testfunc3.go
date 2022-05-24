package testpkg

import (
	"fmt"
	"math/rand"
	"time"
)

func TestMath3() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Println(a)
	}
	fmt.Println("----------------")

	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		fmt.Println(a)
	}

	fmt.Println("-------------")

	for i := 0; i < 10; i++ {
		a := rand.Float32()
		fmt.Println(a)
	}
}
