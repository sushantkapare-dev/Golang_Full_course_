package main

import (
	"fmt"
)

func main() {
	var city string
	city = "Pune"

	var channel = "TrainwithSushant" // inferred to string

	// :=

	subscribers := 5000

	subscribers = subscribers + 1000

	likes, comments := 100, 30

	fmt.Println(city, channel, subscribers, likes, comments)
}
