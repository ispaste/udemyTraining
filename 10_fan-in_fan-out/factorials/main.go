package main

import "fmt"

func instructions() int {
	// Ask the user for a number, and then show the factorial from 0 to that number
	
	var t int
	fmt.Print("Please, write a positive integer: ")
	fmt.Scan(&t)
	return t
}


func fanOut(r int) []chan int {
	// Create as many channels as the number entered by the user, and add a number between 0 and that number to the channel
	
	// The function returns a slice of channels of type int
	
	out := make([]chan int, 0, r)
	for i := 0; i < r; i++ {
		newChan := make(chan int)
		newChan <- i
		fmt.Println(<-newChan)
		out = append(out, newChan)
		close(newChan)
		}
	return out
}

func operating(chans []chan int) chan int{
	// The function receives a slice of channels of type int, and for each channel in the slice, it produces the factorial of whatever is inside the channel, and puts it in an "out" channel, that eventually gets returned
	
	out := make(chan int)
	go func() {
		for _, n := range chans {
			out <- factorial(<-n)
		}
		close(out)
	}()
	return out
}

func factorial(n int) int {
	// Simple factorial function
	
	var fact int = 1
	for i := n; i > 0; i-- {
		fact = fact * i
	}
	return fact
}

func main()  {
	// t is asigned to a positive integer entered by the user
	
	t := instructions()
	fmt.Println(t)
	
	// t is passed to fanOut and the slice of channels that is returned is assigned to "channels"
	
	channels := fanOut(t)
	
	// Verify that channels are populated with integers
	
	for _, c := range channels {
		fmt.Println(c)
	}
	
	// channels is passed to the operating function and the return is stored in channel of ints called "factorials"
	
	factorials := operating(channels)
	
	// Print the contents of "factorials"
	
	for n := range factorials {
		fmt.Println(n)
	}
}