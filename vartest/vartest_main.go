package main

import (
	"fmt"
	"github.com/spf13/viper"
	"vartest/commonfunc"
	"vartest/commonutils"
) //这里应用的包名,试用go mod之后, 是用go.mod文件下面module vartest 这个包做为基础的报名,

func main() {

	//Test1() //这里如果要Test1可用的话,必须是 "go run vartest_main.go vartest1.go"   ,两个文件一起包含,然而下面commonutils.Toupper却不需要

	test()
	test2()
	test3()

	testConfig()

}

func test3() {
	commonfunc.PrintStartInFunc3()
	var name string = "wo shinidie"
	var sex bool = true
	println(name, sex)
	commonfunc.PrintEndInFunc3()

}

func testConfig() {
	viper.SetConfigFile("./conf/test.json")
	//viper.SetConfigType("json")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	name := viper.Get("name")
	fmt.Println(name)
}

func test() {
	var name string
	name = "你好我是name"
	println(name)
}

func test2() {
	var (
		name  string
		age   int
		email string
	)

	age = 18
	name = "pikazo"
	email = "pikazo@smgui.com"

	name = commonutils.Toupper(name) //

	println(age)
	println(name)
	println(email)
}
