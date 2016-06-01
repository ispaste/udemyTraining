package main

import "fmt"

func factorial(n int) int {
	var fact int = 1
	for i := n; i > 0; i-- {
		fact = fact * i
	}
	return fact
}

func main()  {
	var t int
	fmt.Print("Please, write a positive integer: ")
	fmt.Scan(&t)
	fmt.Println("The factorial of", t, "is", factorial(t))
}