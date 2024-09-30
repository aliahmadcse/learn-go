package main

import (
	"fmt"
	"sort"
)

func main() {
	// declare a map
	fruitsCount := make(map[string]int)

	// add key-value pairs
	fruitsCount["apple"] = 5
	fruitsCount["banana"] = 3
	fruitsCount["orange"] = 7

	// fmt.Println(fruitsCount)

	// pointers
	age := 24
	var p *int
	p = &age
	fmt.Println(*p)

	studentCount := 23
	ptrToStudentCount := &studentCount
	fmt.Println("Before incrementing:", *ptrToStudentCount)

	//arrays
	var colors [3]string
	colors[0] = "red"
	colors[1] = "blue"
	colors[2] = "green"
	// fmt.Println(colors)

	numbers := [3]int{12, 22, 32}
	fmt.Println(numbers)
	fmt.Println(len(numbers))

	// slices
	colorSlice := []string{"Red", "Green", "Yellow"}
	fmt.Println(colorSlice)
	fmt.Println(len(colorSlice))
	colorSlice = append(colorSlice, "Purple")
	fmt.Println(colorSlice)
	fmt.Println(len(colorSlice))

	colorSlice = append(colorSlice[1:len(colorSlice)])
	fmt.Println(colorSlice)

	nums := make([]int, 5)
	nums[0] = 10
	nums[1] = 4
	fmt.Println(nums)
	sort.Ints(nums)
	fmt.Println(nums)

	// maps
	fruitPrices := map[string]float64{
		"apple":  0.5,
		"banana": 0.7,
		"orange": 0.8,
	}

	keys := make([]string, 0, len(fruitPrices))
	for fruit := range fruitPrices {
    keys = append(keys, fruit)
  }
	fmt.Println(keys)

	fmt.Println(fruitPrices)
	delete(fruitPrices, "banana")
	fmt.Println(fruitPrices)

	// range
	for fruit, price := range fruitPrices {
		fmt.Printf("%s: %.2f\n", fruit, price)
	}

	states := make(map[string]string)
	states["WA"] = "Washington"

	for abbrev, state := range states {
    fmt.Printf("%s: %s\n", abbrev, state)
  }

	// structs
	type Person struct {
    Name string
    Age  int
  }
	person := Person{Name: "John Doe", Age: 25}
	fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
}
