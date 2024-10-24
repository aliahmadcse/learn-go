package data

import (
	"encoding/json"
	"os"
)

type City struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Population  int     `json:"population"`
	HasBeach    bool    `json:"hasBeach"`
	HasMountain bool    `json:"hasMountain"`
	TempC       float64 `json:"tempC"`
}

type DataReader interface {
	ReadData() ([]City, error)
}

type reader struct {
	path string
}

func NewReader() DataReader {
	return reader{path: "./data/cities.json"}
}

func (r reader) ReadData() ([]City, error) {
	file, err := os.ReadFile(r.path)
	if err != nil {
		return nil, err
	}

	var cities []City
	err = json.Unmarshal(file, &cities)
	if err != nil {
		return nil, err
	}

	return cities, nil
}
