package combatpkg

import (
	"fmt"
)

//这个包是我自己设想的实战应用

//自定义错误类型
type CusError struct {
	Code int
	Msg  string
}

func (e *CusError) Error() string {
	return e.Msg
}

type CusOtherError struct {
	Code int
	Msg  string
}

func (e CusOtherError) Error() string {
	return e.Msg
}

func Combat() {
	test()
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case *CusError:
				fmt.Printf("%p", err)
				fmt.Println(err.(*CusError).Error())
			case CusOtherError:
				fmt.Printf("%p", &err)
				fmt.Println(err.(CusOtherError).Error())
			default:
				fmt.Println("未知错误")
			}
		}
	}()

	//panic(&CusError{Code: 1, Msg: "自定义错误"})

	cusError := &CusError{Code: 2, Msg: "自定义错误"}
	//打印cusOtherError的地址
	fmt.Printf("%p", cusError)
	panic(cusError)

	cusOtherError := CusOtherError{Code: 2, Msg: "自定义错误2"}
	//打印cusOtherError的地址
	fmt.Printf("%p", &cusOtherError)
	panic(cusOtherError)

}
