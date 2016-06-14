package main

import "fmt"

func factorial(n int) chan int {
	// por algún motivo, esta declaración no funcion
	// var out chan int
	out := make(chan int)
	go func() {
		var fact int = 1
		for i := n; i > 0; i-- {
			fact = fact * i
		}
		out <- fact
		close(out)
	}()
	return out
}

func instructions() int {
	var t int
	fmt.Print("Please, write a positive integer: ")
	fmt.Scan(&t)
	return t
}

func main()  {
	t := instructions()
	// fmt.Println(t)
	var c chan int
	c = factorial(t)
	fmt.Print("The factorial of ", t, " is ")
	for n := range c {
		fmt.Println(n)
	}
}