package main

import "fmt"

func location(city string) (string, string) {
	var region string
	var continent string

	switch city {
	case "Delhi", "Hyderabad", "Mumbai", "Kolkata", "Chennai", "Kochi":
		region, continent = "India", "Asia"
	case "Columbus", "Cleveland", "Dayton":
		region, continent = "Ohio, USA", "North America"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

func main() {
	region, continent := location("Columbus")
	fmt.Printf("Allen works in %s, %s\n", region, continent)
}
