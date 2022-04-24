package testpkg

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

/**
 * @description:继承error接口
 * @return:返回一个string
 */
func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func TestMyError() {
	m := new(MyError)
	m.When = time.Now()
	m.What = "这是我自定义的错误"
	//m.Error()
	fmt.Println(m.Error())
}
