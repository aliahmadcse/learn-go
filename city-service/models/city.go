package models

type CityTemp interface {
	Name() string
	TempC() float64
	TempF() float64
}

type city struct {
	name  string
	tempC float64
}

func NewCity(name string, tempC float64) CityTemp {
	return city{name: name, tempC: tempC}
}

func (c city) Name() string {
	return c.name
}

func (c city) TempC() float64 {
	return c.tempC
}

func (c city) TempF() float64 {
	return (c.tempC * 9 / 5) + 32
}
