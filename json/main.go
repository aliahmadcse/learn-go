// Write your answer here, and then test your code.
// Your job is to implement the getCartFromJson() method.

package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.

type cartItem struct {
	Name     string
	Price    float64
	Quantity int
}

// getCartFromJson() returns a slice containing cartItem objects.
func getCartFromJson(jsonString string) []cartItem {
	var cart []cartItem
	// Your code goes here.

	decoder := json.NewDecoder(strings.NewReader(jsonString))
	_, err := decoder.Token()
	if err != nil {
		panic("token not valid")
	}

	for decoder.More() {
		var item cartItem
		decoder.Decode(&item)
		cart = append(cart, item)
	}

	return cart
}

func main() {
	jsonString :=
		`[{"name":"apple","price":2.99,"quantity":3},
  {"name":"orange","price":1.50,"quantity":8},
  {"name":"banana","price":0.49,"quantity":12}]`

	result := getCartFromJson(jsonString)
	fmt.Println(result)
	time.Now()

}
