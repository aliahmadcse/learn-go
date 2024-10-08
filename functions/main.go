package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func area(a int, b int) (*int, error) {
	if a <= 0 || b <= 0 {
		return nil, fmt.Errorf("Both sides must be positive integers")
	}

	area := a * b
	return &area, nil
}

func main() {
	fmt.Println("Result:", add(5, 7))
	fmt.Println("Starting")

	sub := func(x int, y int) int {
		return x - y
	}

	fmt.Println("Result:", sub(10, 3))
	fmt.Println("Ending")

	area, err := area(-5, 7)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Area:", *area)
	}
  
}
