package testpkg

import (
	"fmt"
	"sort"
)

func TestSort() {
	s := []int{2, 4, 1, 3}
	sort.Ints(s)
	fmt.Printf("s: %v\n", s)
}

func TestMySort() {
	n := NewInts{2, 4, 1, 3}
	sort.Sort(n)
	fmt.Println(n)
}

func TestStructSort() {
	ls := testSlice{
		{Name: "n1", Age: 12},
		{Name: "n2", Age: 11},
		{Name: "n3", Age: 10},
	}

	fmt.Println(ls) //[{n1 12} {n2 11} {n3 10}]
	sort.Sort(ls)
	fmt.Println(ls) //[{n3 10} {n2 11} {n1 12}]
}
