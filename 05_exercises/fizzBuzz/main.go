package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		// another way of doing the first condition
		// if i % 15 == 0

		if i%3 == 0 && i%5 != 0 {
			fmt.Println("Fizz")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
