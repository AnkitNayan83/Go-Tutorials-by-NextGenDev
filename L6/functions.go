package main

import "fmt"

// func <name of the func>() <type of return> {

// }

func greet(name string) string {
	grt := fmt.Sprintf("Hello %v, nice to meet you",name)
	return grt
}

func magic(num1 int,num2 int) (addition int,multiply int) {
	add := num1+num2
	mul := num2*num1

	return add,mul
}

func sum(nums ...int) int {
	var ans int
	for _,val := range nums {
		ans+= val // ans = ans+val
	}

	return ans
}

type DivFunc func(int,int) float64

func HOF(a,b int, func1 DivFunc) float64 {
	return func1(a,b)
}

func div(a,b int) float64 {
	return float64(a/b)
}

func main() {

	func () {
		fmt.Print("Hello World!\n")
	} ()

	fmt.Println(greet("Jake"))

	n1,_ := magic(5,4)

	fmt.Println(n1)


	ans := sum(1,2,3,4,5,6,7,8,9)

	fmt.Println(ans)

	hof := HOF(22,4,div)

	fmt.Println(hof)


}