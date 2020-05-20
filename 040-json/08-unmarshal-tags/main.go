//unmarshal and encode stuff..package main
//use Go data structure values for printing

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cities []struct {
	Postal    string  `json:"Postal"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	Address   string  `json:"Address"`
	CityName  string  `json:"City"`
	State     string  `json:"State"`
	ZipCode   string  `json:"Zip"`
	Country   string  `json:"Country"`
}

func main() {
	var data cities
	recvd := `[{"Postal":"zip","Latitude":37.7668,"Longitude":-122.3959,
	"Address":"","City":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},
	{"Postal":"zip","Latitude":37.371991,"Longitude":-122.02602,
	"Address":"","City":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`

	err := json.Unmarshal([]byte(recvd), &data)
	if err != nil {
		log.Fatalln("error unmarshalling", err)
	}

	for i, v := range data {
		fmt.Println(i, v.CityName, v.ZipCode)
	}
	fmt.Println(data)
}
