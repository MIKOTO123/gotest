package testpkg

import (
	"errors"
	"fmt"
)

func TestError() {
	s, err := check("")
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
	} else {
		fmt.Printf("s: %v\n", s)
	}
}

/**
 * @Description:检测是否为空
 */
func check(s string) (string, error) {
	if s == "" {
		err := errors.New("字符串不能为空")
		return "", err
	} else {
		return s, nil
	}
}
