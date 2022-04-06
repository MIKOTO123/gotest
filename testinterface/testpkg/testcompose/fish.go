package testcompose

import "fmt"

type Fish struct {
}

func (fish Fish) Fly() {
	fmt.Println("fly")
}

func (fish Fish) Swim() {
	fmt.Println("swim...")
}
