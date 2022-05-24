package commonfunc

import (
	"fmt"
	"runtime"
)

// 获取正在运行的函数名
func RunFuncName(skip int) string {
	pc := make([]uintptr, 1)
	runtime.Callers(skip, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func PrintEndInFunc3() {
	fmt.Println("====================================", RunFuncName(3), "====================================")
}

func PrintStartInFunc3() {
	fmt.Println("====================================", RunFuncName(3), "====================================")
}
