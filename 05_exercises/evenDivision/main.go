package main

import "fmt"

func main() {
	var number int
	fmt.Print("Please, enter an integer: ")
	fmt.Scan(&number)
	e, half := even(number)
	fmt.Println(number, "divided by 2 is", e)
	fmt.Println("Is", number, "even?", half)

	// My initial solution
	// fmt.Println(even(number))
}

func even(n int) (f float32, b bool) {
	f = float32(n) / 2
	return f, n%2 == 0
}
