package main

import "fmt"

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true

	// %T is the type of the variable
	// %v is the value in the default format

	fmt.Printf("%T: %v \n", a, a)
	fmt.Printf("%T: %v \n", b, b)
	fmt.Printf("%T: %v \n", c, c)
	fmt.Printf("%T: %v \n", d, d)
}
