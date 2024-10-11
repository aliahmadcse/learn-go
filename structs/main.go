package main

import "fmt"

type city struct {
	name       string
	population int
}

func newCity(name string) city {
	return city{name: name}
}

// (c city) represents the receiver
func (c city) printPopulation() {
	fmt.Printf("The population of %s is %d\n", c.name, c.population)
}

func (c *city) updatePopulation(newPopulation int) {
	c.population = newPopulation
}

func main() {
	london := newCity("London")
	london.updatePopulation(97823000)

	london.printPopulation()
	fmt.Println(london.name)
}
