package testpkg

import (
	"fmt"
	"os"
)

func getProgress() {
	// 获得当前正在运行的进程id
	fmt.Printf("os.Getpid(): %v\n", os.Getpid())
	// 父id
	fmt.Printf("os.Getppid(): %v\n", os.Getppid())
}

func StartProcess() {
	//设置新进程的属性
	attr := &os.ProcAttr{
		//files指定新进程继承的活动文件对象
		//前三个分别为，标准输入、标准输出、标准错误输出
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		//新进程的环境变量
		Env: os.Environ(),
	}
	p, err := os.StartProcess("C:\\Windows\\System32\\notepad.exe", []string{"C:\\Windows\\System32\\notepad.exe", "D:\\a.txt"}, attr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	fmt.Println("进程ID：", p.Pid)
}

func FindProcess() {
	// 查找指定进程
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	fmt.Println("进程ID：", p.Pid)
}

func KillProcess() {
	// 杀死指定进程
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		fmt.Println(err)
	}
	////向p进程发送退出信号
	//		p.Signal(os.Kill)
	err = p.Kill()
	if err != nil {
		fmt.Println(err)
	}
}
