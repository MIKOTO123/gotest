package main

import "fmt"

func main() {
	client := NewExampleClient(SetName("hahha"))
	fmt.Printf("%v\n", client)
}

//定义结构体
type ExampleClient struct {
	Name string
	Job  string
}

//定义配置选项函数（关键）
type Option func(*ExampleClient)

func SetName(name string) Option { // 返回一个Option类型的函数（闭包）：接受ExampleClient类型指针参数并修改之
	return func(this *ExampleClient) {
		this.Name = name
	}
}

func SetJob(job string) Option { // 返回一个Option类型的函数（闭包）：接受ExampleClient类型指针参数并修改之
	return func(this *ExampleClient) {
		this.Job = job
	}
}

//应用函数选项配置
func NewExampleClient(opts ...Option) ExampleClient {
	// 初始化默认值
	defaultClient := ExampleClient{
		Name: "default",
		Job:  "default",
	}

	// 依次调用opts函数列表中的函数，为结构体成员赋值
	for _, o := range opts {
		o(&defaultClient)
	}

	return defaultClient
}
