package main

import "fmt"

func main() {

	var num int = 5
	var num1 int = num

	var ptr *int = &num

	num1++
	*ptr++ // dereference

	// fmt.Println(ptr)
	// fmt.Println(*ptr)

	fmt.Println(num)

	// Call by value and call by reference


	var name string = "Andrew"

	ptr1 := &name

	// CBV(name)
	CBR(ptr1)

	fmt.Println(name)

}

func CBV(name string) {
	name = "Jake"
}

func CBR(name *string) {
	*name = "Jake"
}