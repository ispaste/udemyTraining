package main

import "fmt"

func main() {
	// Print with format: %d = decimal, %b = binary, %x = hexadecimal, %o = octal, %q = UTF
	// %X = hexadecimal in capital letters, %#x = complete 8-bit hexadecimal
	// \t = tab, \n = new line
	fmt.Printf("%d - %b - %x - %o \n", 42, 42, 42, 42)
	fmt.Printf("%d \t %b \t %X \n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %#x \n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
}
