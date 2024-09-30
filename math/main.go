package main

import (
	"fmt"
	"math"
)

func addNumbers(numbers ...int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}

	for i := 0; i < len(numbers); i++ {
	}

	return result
}

func main() {
	i1, i2, i3 := 12, 45, 68
	sum := addNumbers(i1, i2, i3)
	fmt.Println(sum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	sumFloat := f1 + f2 + f3
	fmt.Println("Sum of floats: ", sumFloat)

	sumFloat = math.Round(sumFloat*100) / 100
	fmt.Println("Rounded sum of floats: ", sumFloat)

	circleRadius := 15.56
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference of circle with radius %.1f is %.2f\n", circleRadius, circumference)
}
