package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Welcome to our game.");

	lottery := 250
	var user_input int

	fmt.Print("Enter your value: ");
	fmt.Scanln(&user_input)

	diff := math.Abs(float64(lottery - user_input))



	if diff <= 50 {
		fmt.Println("You won $100");
	} else if diff>50 && diff<=100 {
		fmt.Println("You won $50")
	} else {
		fmt.Print("You lost")
	}
}