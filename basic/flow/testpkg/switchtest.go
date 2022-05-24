package testpkg

import "fmt"

func SwitchTest() {
	grade := 'A'

	switch grade {
	case 'A':
		println("优秀")
	case 'B':
		println("良")
	default:
		println("一般")
	}
}

func SwitchTest2() {
	day := 2
	switch day {
	case 1, 2, 3, 4, 5:
		println("工作日")
	case 6, 7:
		println("休息日")
	default:
		println("非法输入")
	}
}

func SwithTeste3() {
	//score := 90
	var score int
	fmt.Scan(&score)

	switch {
	case score >= 90:
		println("干的不从,好好休息一下")
	case score > 80 && score < 90:
		println("还需要好好学习")
	default:
		println("加班晚自习")
	}

}

func SwitchTest4() {

	a := 100
	switch a {
	case 100:
		println("100")
		fallthrough
	case 200:
		println("200")
	case 300:
		println("300")
	default:
		println("othere")
	}
	//	打印 100 200

}
