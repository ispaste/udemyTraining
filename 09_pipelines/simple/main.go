package main

import "fmt"

func main()  {
	// Set up the pipeline
	var numbers chan int
	numbers = input (10, 77)
	
	var result chan int
	result = square (numbers)
	
	// Consume the output
	fmt.Println(<-result)
	fmt.Println(<-result)
	
	
}

func input(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(list chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range list {
			out <- n * n
		}
		close(out)
	}()
	return out
}