package testpkg

import (
	"encoding/json"
	"fmt"
	"json/testpkg/structs"
	"os"
)

func TestJsonToStruct() {
	p := structs.Person{
		Name:   "pikazo",
		Age:    0,
		Email:  "pikazo@smgui.com",
		Parent: []string{"p1", "p2"},
	}
	b, _ := json.Marshal(p)
	fmt.Printf("b: %v\n", string(b))
}

func TestStructToJson() {
	b := []byte(`{"Name":"tom","Age":20,"Email":"tom@gmail.com"}`)
	var p structs.Person
	json.Unmarshal(b, &p)
	fmt.Printf("p: %v\n", p)
}

func TestJsontoMap() {
	var result map[string]interface{}
	jsonStr := `{"Name":"tom","Age":20,"Email":""}`
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err == nil {
		fmt.Printf("result: %v\n", result)
	}
}

func TestJson1() {

	p := structs.Person{
		Name:   "tom",
		Age:    20,
		Email:  "tom@gmail.com",
		Parent: []string{"big tom", "big kite"},
	}

	f, _ := os.OpenFile("a.json", os.O_WRONLY|os.O_CREATE, 0777)
	defer f.Close()
	e := json.NewEncoder(f)
	err := e.Encode(p)
	if err != nil {
		fmt.Println(err)
	}

}

func TestJson2() {
	f, _ := os.Open("a.json")
	defer f.Close()
	d := json.NewDecoder(f)
	var v map[string]interface{}
	//var v structs.Person
	d.Decode(&v)

	fmt.Printf("v: %v\n", v)
	fmt.Printf("T: %T\n", v)
}
