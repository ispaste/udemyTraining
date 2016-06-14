package main

import "fmt"

func main()  {
	// Set up the pipeline and consume the output
	for n := range square(square(input(12, 1, 5, -3))) {
		fmt.Println(n)
	}
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