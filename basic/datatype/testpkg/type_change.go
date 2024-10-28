package testpkg

import (
	"fmt"
	"strconv"
)

//类型转换的几种演示

func StringAndInt() {
	// 字符串转整数
	strValue := "123"
	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(intValue) // 输出: 123

	// 整数转字符串
	intValue = 123
	strValue = strconv.Itoa(intValue)
	fmt.Println(strValue) // 输出: "123"
}

func Affirm() {
	// 接口类型
	var i interface{} = "hello"

	// 类型断言
	s := i.(string)
	fmt.Println(s) // 输出: hello

	// 检查类型断言是否成功
	f, ok := i.(float64)
	if ok {
		fmt.Println(f)
	} else {
		fmt.Println("Not a float64")
	}
}

func SwitchType() {
	//接口类型
	var i interface{} = 123

	// 类型转换
	switch v := i.(type) {
	case string:
		fmt.Println("String:", v)
	case int:
		fmt.Println("Integer:", v)
	default:
		fmt.Println("Unknown type")
	}
}
