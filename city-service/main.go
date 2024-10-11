package main

import (
	"fmt"
	"learn-go/city-service/models"
)



func main() {
	london := models.NewCity("London", 7.5)

	fmt.Printf("The temperature in London is %.2f°C or %.2f°F\n", london.TempC(), london.TempF())
}