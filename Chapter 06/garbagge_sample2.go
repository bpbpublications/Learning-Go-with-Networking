package main
import (
	"fmt"
	"runtime/debug"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Creating a new Person object and assigning it to p1
	p1 := createPerson("Alice", 30)
	fmt.Println("p1:", p1)

	// Creating a new Person object and assigning it to p2
	p2 := createPerson("Bob", 25)
	fmt.Println("p2:", p2)

	// Setting p1 to nil, making it no longer reachable
	p1 = nil

	// Triggering garbage collection explicitly
	// Note: Garbage collection is automatic and doesn't require manual triggering in most cases
	// Explicit triggering is used here for demonstration purposes
	performGarbageCollection()

	// Printing p1 after garbage collection
	fmt.Println("p1 after garbage collection:", p1)

	// Accessing p2 after garbage collection
	fmt.Println("p2 after garbage collection:", p2)
}

func createPerson(name string, age int) *Person {
	p := &Person{
		Name: name,
		Age:  age,
	}
	return p
}

func performGarbageCollection() {
	// Run garbage collection explicitly
	// Note: This is not necessary in most cases as garbage collection is automatic
	// Explicit triggering is used here for demonstration purposes
	fmt.Println("Running garbage collection...")
	debug.FreeOSMemory()
	fmt.Println("Garbage collection complete!")
}