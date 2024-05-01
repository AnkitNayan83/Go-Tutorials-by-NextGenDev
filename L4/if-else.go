package main

import "fmt"

func main() {

	traffic_light := "green"

	if traffic_light == "red" {
		fmt.Print("Stop")
	} else if traffic_light == "green" {
		fmt.Print("Go")
	} else {
		fmt.Print("Wait")
	}
}