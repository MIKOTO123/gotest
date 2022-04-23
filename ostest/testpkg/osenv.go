package testpkg

import (
	"fmt"
	"os"
)

func PrintEnv() {
	fmt.Printf("os.Environ(): %v\n", os.Environ())
	s := os.Getenv("GOPATH")
	fmt.Printf("s: %v\n", s)
}

func GetEnv() {
	s := os.Getenv("GOPATH")
	fmt.Printf("s: %v\n", s)
	fmt.Printf("os.Getenv(\"JAVA_HOME\"): %v\n", os.Getenv("JAVA_HOME"))
}

func LookUpEnv() {
	s, b := os.LookupEnv("GOPATH")
	if b {
		fmt.Printf("s: %v\n", s)
	}
}

func SetEnv() {
	os.Setenv("env1", "env1")
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")
	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))
}

func ClearEnv() {
	// os.Clearenv()
}
