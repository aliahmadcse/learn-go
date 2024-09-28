package main

import "fmt"

const PI float32 = 3.1415

func main() {
	var name string = "John Wick"
	fmt.Println(name)
	fmt.Printf("The type of the variable is %T\n", name)

	var defaultInt int
	fmt.Println(defaultInt)

	var anotherName = "Ali Ahmad"
	fmt.Println(anotherName)

	// only works inside the functions
	myString := "This is also a string"
	fmt.Printf("The type of myString is %T\n", myString)

	fmt.Println(PI)
}
