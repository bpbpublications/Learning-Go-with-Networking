package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	// Create a new Person message
	person := &proto.Person{
		Name:    "John",
		Age:     25,
		Hobbies: []string{"Reading", "Gaming"},
	}

	// Serialize the Person message to bytes
	serializedData, err := proto.Marshal(person)
	if err != nil {
		log.Fatal("Error while marshaling:", err)
	}

	fmt.Println("Serialized Data:", serializedData)

	// Deserialize the bytes to a new Person message
	newPerson := &proto.Person{}
	err = proto.Unmarshal(serializedData, newPerson)
	if err != nil {
		log.Fatal("Error while unmarshaling:", err)
	}

	fmt.Println("Deserialized Person:", newPerson)
}