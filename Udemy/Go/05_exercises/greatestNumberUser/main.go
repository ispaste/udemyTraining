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

func big(n int) bool {
	if n {
		return true
	}
}

func main() {
	numbers := make([]int, 50)
	fmt.Print("Please, enter a list of numbers separated by spaces: ")
    numbers[] = fmt.Scan(n...)
    // fmt.Printf("%T", numbers)
	// fmt.Println(numbers)
    


	biggest := greatest(numbers...)
	fmt.Println("The largest number of the list is", biggest)

}
