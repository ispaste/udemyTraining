package main

import "fmt"

func instructions() int {
	var t int
	fmt.Print("Please, write a positive integer: ")
	fmt.Scan(&t)
	return t
}

/* func generation(n int) chan int{
	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
} */

func fanOut(r int) []chan int {
	// var out []chan int
	out := make([]chan int, 0, r)
	for i := 0; i < r; i++ {
		newChan := make(chan int)
		newChan <- i
		fmt.Println(<-newChan)
		out = append(out, newChan)
		// close(newChan)
		}
	return out
}

func operating(chans []chan int) chan int{
	out := make(chan int)
	go func() {
		for _, n := range chans {
			out <- factorial(<-n)
		}
		close(out)
	}()
	return out
}

/* func alltogethernow(numbers chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range numbers {
			out <- factorial(n)
		}
		close(out)
	}()
	return out
} */

func factorial(n int) int {
	var fact int = 1
	for i := n; i > 0; i-- {
		fact = fact * i
	}
	return fact
}



func main()  {
	t := instructions()
	fmt.Println(t)
	
	channels := fanOut(t)
	
	for _, c := range channels {
		fmt.Println(c)
	}
	
	factorials := operating(channels)
	
	for n := range factorials {
		fmt.Println(n)
	}
}