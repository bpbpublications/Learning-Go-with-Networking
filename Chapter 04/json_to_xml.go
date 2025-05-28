package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
}

func main() {
	person := Person{
		Name: "John",
		Age:  25,
	}

	xmlData, err := xml.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(xmlData))
}