package main

import "fmt"

func main() {
	var number1 int
	var number2 int
	fmt.Print("Please, enter two numbers: ")
	fmt.Scan(&number1, &number2)

	if number1 > number2 {
		fmt.Println(number1, "divided by", number2, "is", number1/number2, "with a remainder of", number1%number2)
	} else {
		fmt.Println(number2, "divided by", number1, "is", number2/number1, "with a remainder of", number2%number1)
	}

}
