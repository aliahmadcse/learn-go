package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()

	// Print the current date and time
	fmt.Println("Current date and time: ", n)

	// Print the current date in the format "YYYY-MM-DD"
	fmt.Println("Current date (YYYY-MM-DD): ", n.Format("2006-01-02"))

	// Print the current time in the format "HH:MM:SS"
	fmt.Println("Current time (HH:MM:SS): ", n.Format("15:04:05"))

	// Print the current time in the format "YYYY-MM-DD HH:MM:SS"
	fmt.Println("Current date and time (YYYY-MM-DD HH:MM:SS): ", n.Format("2006-01-02 15:04:05"))

	// go was launched at
	t := time.Date(2009, time.November, 1, 1, 0, 0, 0, time.UTC)
	fmt.Println("Go was launched at: ", t)
}
