package testpkg

import "encoding/xml"

type Person struct {
	// 反引号
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email"`
}
