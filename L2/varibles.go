package main

import "fmt"

func main() {
	// varibles

	// Integers
	var age int

	num := 12

	fmt.Print(age,num)
	// floats
	var marks float64 = 45.5

	fmt.Printf("Your marks is %.2f \n",marks)

	// bool

	// var isAdmin bool = true

	// Arrays

	arr := [5]int{1,2,3,4,5}

	fmt.Print(arr[1])

	// strings

	var name string = "Nextgen"

	// char (rune)

	var char rune = 'a'

	fmt.Print(name,char)
}