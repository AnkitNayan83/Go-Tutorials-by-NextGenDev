package main

import "fmt"

func input() {

	fmt.Println("Enter your name:")
	var name string;

	fmt.Scanln(&name)

	fmt.Printf("Your name has %v charecters",len(name))

}