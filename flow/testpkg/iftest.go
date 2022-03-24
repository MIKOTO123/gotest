package testpkg

import "fmt"

func Iftest() {
	a, b := 100, 2
	if a > b {
		println(a)
	} else {
		println(b)
	}
}

func Iftest2() {
	if age := 16; age > 18 {
		fmt.Println("成年")
		fmt.Printf("age: %v\n", age)
	} else {
		fmt.Println("未成年")
		fmt.Printf("age: %v\n", age)
	}
	//println(age)//这边是没有age变量的
}

func Iftest3() {
	flag1 := true
	if flag1 {
		fmt.Printf("flag: %v\n", flag1)
	}

	fmt.Println(flag1)

	//下面这段会编译不成功,请注意对比
	/*a := 100
	if a {
		fmt.Println("真")
	}*/
}

func Iftest4() {
	var score int
	fmt.Println("请输入一个整形得分,用空格分割")
	fmt.Scan(&score)
	if score > 85 {
		fmt.Println("优秀")
	} else if score >= 70 && score < 85 {
		fmt.Println("良")
	} else if score >= 60 && score < 70 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
}
