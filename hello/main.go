package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("hello from the Go programming language!")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")

	name, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("Error reading name: %v\n", err)
		return
	}
	fmt.Println("Your name is: ", name)

	fmt.Print("Enter your age: ")
	ageStr, err := reader.ReadString('\n')

	age, err := strconv.ParseInt(strings.TrimSpace(ageStr), 10, 32)
	if err != nil {
		fmt.Printf("Error parsing age: %v\n", err)
		return
	}
	fmt.Println("Your age seems like: " + strconv.FormatInt(age, 10))
}
