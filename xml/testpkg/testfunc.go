package testpkg

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func marshal() {
	person := Person{
		Name:  "tom",
		Age:   20,
		Email: "tom@gmail.com",
	}

	// b, _ := xml.Marshal(person)
	b, _ := xml.MarshalIndent(person, "", "  ")

	fmt.Printf("%v\n", string(b))
}

func Unmarshal() {
	s := `<person>
	<name>tom</name>
	<age>20</age>
	<email>tom@gmail.com</email>
	</person>
	`
	b := []byte(s)

	var per Person
	xml.Unmarshal(b, &per)
	fmt.Printf("per: %v\n", per)
}

func Read() {
	b, _ := ioutil.ReadFile("a.xml")
	var per Person
	xml.Unmarshal(b, &per)
	fmt.Printf("per: %v\n", per)
}

func Write() {

	p := Person{
		Name:  "tom",
		Age:   20,
		Email: "tom@gmail.com",
	}

	f, _ := os.OpenFile("a.xml", os.O_WRONLY, 0777) // Writer
	defer f.Close()
	e := xml.NewEncoder(f)
	e.Encode(p)
}
