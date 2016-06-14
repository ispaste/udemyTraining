package main

import "fmt"

func greatest(n ...int) int {
	var largest int
	for _, i := range n {
		if i > largest {
			largest = i
		}
	}
	return largest
}

func main() {
	// var numbers []int
	// fmt.Print("Please, enter a list of numbers separated by spaces: ")
	// fmt.Scan(numbers)
	// fmt.Printf("%T", numbers)
	// fmt.Println(numbers)

	numbers := []int{-4, 7, 20, 8, -2, 15, 42, 12}

	biggest := greatest(numbers...)
	fmt.Println("The largest number of the list is", biggest)

}
