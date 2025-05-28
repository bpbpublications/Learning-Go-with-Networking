package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age int
}

// Custom serialization method
func (p *Person) MarshalJSON() ([]byte, error) {
	// Define a custom JSON structure
	customJSON := struct {
		FullName string `json:"fullName"`
		YearsOld int    `json:"yearsOld"`
	}{
		FullName: p.Name,
		YearsOld: p.Age,
	}

	return json.Marshal(customJSON)
}

// Custom deserialization method
func (p *Person) UnmarshalJSON(data []byte) error {
	// Define a custom JSON structure
	customJSON := struct {
		FullName string `json:"fullName"`
		YearsOld int    `json:"yearsOld"`
	}{}

	err := json.Unmarshal(data, &customJSON)
	if err != nil {
		return err
	}

	p.Name = customJSON.FullName
	p.Age = customJSON.YearsOld

	return nil
}

func main() {
	// Create a new Person object
	person := &Person{
		Name: "John Doe",
		Age: 25,
	}

	// Serialize the Person object to JSON using the custom MarshalJSON method
	jsonData, err := person.MarshalJSON()
	if err != nil {
		log.Fatal("Error while serializing to JSON:", err)
	}

	fmt.Println("Serialized JSON data:", string(jsonData))

	// Create a new empty Person object
	newPerson := &Person{}

	// Deserialize the JSON data into the new Person object using the custom UnmarshalJSON method
	err = newPerson.UnmarshalJSON(jsonData)
	if err != nil {
		log.Fatal("Error while deserializing from JSON:", err)
	}

	fmt.Println("Deserialized Person:", newPerson)
}