package main

import "fmt"

func instructions() int {
	var t int
	fmt.Print("Please, write a positive integer: ")
	fmt.Scan(&t)
	return t
}

func generation(n int) chan int{
	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func alltogethernow(numbers chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range numbers {
			out <- factorial(n)
		}
		close(out)
	}()
	return out
}

func factorial(n int) int {
	var fact int = 1
	for i := n; i > 0; i-- {
		fact = fact * i
	}
	return fact
}



func main()  {
	t := instructions()
	// fmt.Println(t)
	
	input := generation(t)
	
	output := alltogethernow(input)
	
	for n := range output {
		fmt.Println(n)
	}
}