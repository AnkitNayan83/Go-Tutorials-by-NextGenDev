package main

import "fmt"

func main() {

	// intializer (only once) -> (condition -> code -> increment)

	// for //intializer; //contidion; //increment {
	// code
	// }

	// while

	// while(condition) {
	// 	code
	// }

	// for //condtion {
	// code
	// }

	// i = i+1 || i+=1 || i++

	// for i := 0; i < 50; i += 1 {
	// 	fmt.Println("Hello World")
	// }

	// j := 0

	// for j<5 {
	// 	fmt.Println("small")
	// 	j++
	// }

	arr := [5]int{1,2,3,4,5}

	// size of array = len(arr)

	// for i:= 0; i< len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// range
	// for index, value := range array {

	// }

	// if you wnat to ignore then use '_'

	for _,value := range arr {
		fmt.Printf("%v\n",value)
	}
}