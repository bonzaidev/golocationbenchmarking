package iplocation


import "fmt"

type Location struct {
	city string
	state string
	country string
}

type LocalGeo struct {
	Latitude  float64
	Longitude float64
}

func PrintLocation(loc *Location) {
	fmt.Printf("City = %s State = %s Country = %s\n",loc.city, loc.state, loc.country)
}