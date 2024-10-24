package main

import (
	"fmt"
	"learn-go/json-proj/data"
)

func cityMap(cities []data.City) map[string]data.City {
	cityMap := make(map[string]data.City)
	for _, city := range cities {
		cityMap[city.Id] = city
	}
	return cityMap
}

func main() {
	reader := data.NewReader()

	cities, err := reader.ReadData()

	if err != nil {
		fmt.Printf("Error reading data: %v\n", err)
		return
	}
	for _, city := range cities {
		fmt.Printf("%s: Population=%d, Has Beach=%t, Has Mountain=%t, Temperature=%.2f°C\n",
			city.Name, city.Population, city.HasBeach, city.HasMountain, city.TempC)
	}

	cityMap := cityMap(cities)
	fmt.Println("City Map:" + cityMap["ldn"].Name)
}
