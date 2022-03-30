package testpkg

import "fmt"

func TestGoTo() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i >= 2 && j >= 2 {
				goto END //直接跳出双层循环
			}
			fmt.Printf("%v,%v\n", i, j)
		}
	}

END:
	fmt.Printf("END...")
}

func TestGoTo2() {
	i := 1
	if i >= 2 {
		fmt.Println("2")
	} else {
		goto END
	}

END:
	fmt.Println("END....")
}
